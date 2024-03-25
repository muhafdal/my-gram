package database

import (
	"database/sql"
	"fmt"
	"log"
	"myGram/infra/config"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func handleDatabaseConnection() {

	appConfig := config.AppConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DbHost,
		appConfig.DbPort,
		appConfig.DbUser,
		appConfig.DbPassword,
		appConfig.DbName,
	)

	db, err = sql.Open(appConfig.DbDialect, dsn)

	if err != nil {
		log.Panicln("error occured while trying to validate database arguments: ", err.Error())
		return
	}

	if err := db.Ping(); err != nil {
		log.Panicln("error occured while trying to connect to database :", err.Error())
		return
	}

}

func handleRequiredTables() {
	const (
		createTableUsersQuery = `
			CREATE TABLE IF NOT EXISTS
				users
					(
						id SERIAL PRIMARY KEY,
						username VARCHAR(25) NOT NULL,
						email VARCHAR(50) NOT NULL,
						password TEXT NOT NULL,
						age INT NOT NULL,
						created_at TIMESTAMPTZ DEFAULT now(),
						updated_at TIMESTAMPTZ DEFAULT now(),
						CONSTRAINT
							unique_email
								UNIQUE(email),
						CONSTRAINT
							unique_username
								UNIQUE(username)
					)
		`
		createTablePhotosQuery = `
			CREATE TABLE IF NOT EXISTS
				photos
					(
						id SERIAL PRIMARY KEY,
						title TEXT NOT NULL,
						caption TEXT NOT NULL,
						photo_url TEXT NOT NULL,
						user_id INT NOT NULL,
						created_at TIMESTAMPTZ DEFAULT now(),
						updated_at TIMESTAMPTZ DEFAULT now(),
						CONSTRAINT
							fk_photos_user_id
								FOREIGN KEY(user_id)
									REFERENCES
										users(id)
											ON DELETE CASCADE
					)
		`
		createTableSocialMediaQuery = `
			CREATE TABLE IF NOT EXISTS
				social_media
					(
						id SERIAL PRIMARY KEY,
						name VARCHAR(50) NOT NULL,
						social_media_url TEXT NOT NULL,
						user_id INT NOT NULL,
						created_at TIMESTAMPTZ DEFAULT now(),
						updated_at TIMESTAMPTZ DEFAULT now(),
						CONSTRAINT
							fk_social_media_user_id
								FOREIGN KEY(user_id)
									REFERENCES
										users(id)
											ON DELETE CASCADE
					)
		`
		createTableCommentsQuery = `
			CREATE TABLE IF NOT EXISTS
				comments
					(
						id SERIAL PRIMARY KEY,
						user_id INT NOT NULL,
						photo_id INT NOT NULL,
						message TEXT NOT NULL,
						created_at TIMESTAMPTZ DEFAULT now(),
						updated_at TIMESTAMPTZ DEFAULT now(),
						CONSTRAINT
							fk_comments_user_id
								FOREIGN KEY(user_id)
									REFERENCES
										users(id)
											ON DELETE CASCADE,
						CONSTRAINT 
							fk_comments_photo_id
								FOREIGN KEY(photo_id)
									REFERENCES
										photos(id)
											ON DELETE CASCADE
					)
		`
	)

	_, err = db.Exec(createTableUsersQuery)

	if err != nil {
		log.Panic("error while creating users table: ", err.Error())
		return
	}

	_, err = db.Exec(createTablePhotosQuery)

	if err != nil {
		log.Panic("error while creating photos table: ", err.Error())
		return
	}

	_, err = db.Exec(createTableSocialMediaQuery)

	if err != nil {
		log.Panic("error while creating social_media table: ", err.Error())
		return
	}

	_, err = db.Exec(createTableCommentsQuery)

	if err != nil {
		log.Panic("error while creating comments table: ", err.Error())
		return
	}

}

func InitializeDatabase() {
	handleDatabaseConnection()
	handleRequiredTables()
}

func GetInstanceDatabaseConnection() *sql.DB {
	return db
}
