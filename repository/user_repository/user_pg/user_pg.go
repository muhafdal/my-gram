package user_pg

import (
	"database/sql"
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/user_repository"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user_repository.UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (userRepo *userRepositoryImpl) Create(userPayload *entity.User) (*dto.UserResponse, errs.Error) {

	var user dto.UserResponse
	err := userRepo.db.QueryRow(
		createUserQuery,
		userPayload.Username,
		userPayload.Email,
		userPayload.Age,
		userPayload.Password,
	).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Age,
	)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (userRepo *userRepositoryImpl) FetchByEmail(email string) (*entity.User, errs.Error) {

	user := entity.User{}
	err := userRepo.db.QueryRow(fetchUserByEmail, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Age,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (userRepo *userRepositoryImpl) FetchByUsername(username string) (*entity.User, errs.Error) {

	user := entity.User{}
	err := userRepo.db.QueryRow(fetchUserByUsername, username).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Age,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (userRepo *userRepositoryImpl) Update(userPayload *entity.User) (*dto.UserUpdateResponse, errs.Error) {

	row := userRepo.db.QueryRow(updateUserQuery, userPayload.Id, userPayload.Username, userPayload.Email)

	var user dto.UserUpdateResponse
	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.Age,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (userRepo *userRepositoryImpl) FetchById(userId int) (*entity.User, errs.Error) {

	user := entity.User{}
	err := userRepo.db.QueryRow(fetchUserById, userId).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Age,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &user, nil
}

func (userRepo *userRepositoryImpl) Delete(userId int) errs.Error {

	_, err := userRepo.db.Exec(deleteUserQuery, userId)

	if err != nil {
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}
