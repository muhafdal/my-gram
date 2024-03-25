package user_repository

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
)

type UserRepository interface {
	Create(userPayload *entity.User) (*dto.UserResponse, errs.Error)
	FetchByEmail(email string) (*entity.User, errs.Error)
	FetchByUsername(username string) (*entity.User, errs.Error)
	FetchById(userId int) (*entity.User, errs.Error)
	Update(userPayload *entity.User) (*dto.UserUpdateResponse, errs.Error)
	Delete(userId int) errs.Error
}
