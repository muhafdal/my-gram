package user_service

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/pkg/helper"
	"myGram/repository/user_repository"
	"net/http"
)

type UserService interface {
	Add(userPayload *dto.NewUserRequest) (*dto.GetUserResponse, errs.Error)
	Login(userPayload *dto.UserLoginRequest) (*dto.GetUserResponse, errs.Error)
	Edit(userId int, userPayload *dto.UserUpdateRequest) (*dto.GetUserResponse, errs.Error)
	Remove(userId int) (*dto.GetUserResponse, errs.Error)
}

type userServiceImpl struct {
	ur user_repository.UserRepository
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userServiceImpl{
		ur: userRepo,
	}
}

func (u *userServiceImpl) Add(userPayload *dto.NewUserRequest) (*dto.GetUserResponse, errs.Error) {

	err := helper.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username: userPayload.Username,
		Email:    userPayload.Email,
		Age:      uint(userPayload.Age),
		Password: userPayload.Password,
	}

	err = user.HashPassword()

	if err != nil {
		return nil, err
	}

	err = u.validateDuplicateEmailAndUsername(
		userPayload.Email,
		userPayload.Username,
	)

	if err != nil {
		return nil, err
	}

	response, err := u.ur.Create(user)

	if err != nil {
		return nil, err
	}

	return &dto.GetUserResponse{
		StatusCode: http.StatusCreated,
		Message:    "create new user successfully",
		Data:       response,
	}, nil
}

func (us *userServiceImpl) Login(userPayload *dto.UserLoginRequest) (*dto.GetUserResponse, errs.Error) {

	err := helper.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	user, err := us.ur.FetchByEmail(userPayload.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewBadRequestError("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(userPayload.Password)

	if isValidPassword == false {
		return nil, errs.NewBadRequestError("invalid email/password")
	}

	token := user.GenerateToken()

	return &dto.GetUserResponse{
		StatusCode: http.StatusOK,
		Message:    "successfully loged in",
		Data: dto.TokenResponse{
			Token: token,
		},
	}, nil
}

func (u *userServiceImpl) Edit(userId int, userPayload *dto.UserUpdateRequest) (*dto.GetUserResponse, errs.Error) {

	err := helper.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	_, err = u.ur.FetchById(userId)

	if err != nil {
		return nil, err
	}

	err = u.validateDuplicateEmailAndUsername(
		userPayload.Email,
		userPayload.Username,
	)

	if err != nil {
		return nil, err
	}

	dbPayload := &entity.User{
		Id:       userId,
		Email:    userPayload.Email,
		Username: userPayload.Username,
	}

	response, err := u.ur.Update(dbPayload)

	if err != nil {
		return nil, err
	}

	return &dto.GetUserResponse{
		StatusCode: http.StatusOK,
		Message:    "user has been successfully updated",
		Data:       response,
	}, nil
}

func (u *userServiceImpl) Remove(userId int) (*dto.GetUserResponse, errs.Error) {
	err := u.ur.Delete(userId)

	if err != nil {
		return nil, err
	}

	return &dto.GetUserResponse{
		StatusCode: http.StatusOK,
		Message:    "Your account has been successfully deleted",
		Data:       nil,
	}, nil
}
