package user_pg

const (
	createUserQuery = `
		INSERT INTO 
			users (username, email, age, password)
		VALUES
			($1, $2, $3, $4)
		RETURNING
			id, username, email, age
	`

	fetchUserByEmail = `
		SELECT
			id, 
			username, 
			email, 
			password, 
			age, 
			created_at, 
			updated_at
		FROM
			users
		WHERE
			email = $1
	`

	fetchUserByUsername = `
	SELECT
		id, 
		username, 
		email, 
		password, 
		age, 
		created_at, 
		updated_at
	FROM
		users
	WHERE
		username = $1
`

	fetchUserById = `
		SELECT
			id, 
			username, 
			email, 
			password, 
			age, 
			created_at, 
			updated_at
		FROM
			users
		WHERE
			id = $1
	`

	updateUserQuery = `
		UPDATE 
			users
		SET
			username= $2,
			email= $3,
			updated_at = now()
		WHERE
			id = $1
		RETURNING
			id, email, username, age, updated_at
	`

	deleteUserQuery = `
		DELETE
		FROM
			users
		WHERE
			id = $1
	`
)
