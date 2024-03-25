package photo_service

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/pkg/helper"
	"myGram/repository/photo_repository"
	"net/http"
)

type PhotoService interface {
	AddPhoto(userId int, photoPayload *dto.NewPhotoRequest) (*dto.AddNewPhotoResponse, errs.Error)
	GetPhotos() (*dto.GetPhotoResponse, errs.Error)
	UpdatePhoto(photoId int, photoPayload *dto.PhotoUpdateRequest) (*dto.UpdatePhotoResponse, errs.Error)
	DeletePhoto(photoId int) (*dto.GetPhotoResponse, errs.Error)
}

type photoServiceImpl struct {
	pr photo_repository.PhotoRepository
}

func NewPhotoService(photoRepository photo_repository.PhotoRepository) PhotoService {
	return &photoServiceImpl{
		pr: photoRepository,
	}
}

func (p *photoServiceImpl) AddPhoto(userId int, photoPayload *dto.NewPhotoRequest) (*dto.AddNewPhotoResponse, errs.Error) {

	err := helper.ValidateStruct(photoPayload)

	if err != nil {
		return nil, err
	}

	photo := &entity.Photo{
		Title:    photoPayload.Title,
		Caption:  photoPayload.Caption,
		PhotoUrl: photoPayload.PhotoUrl,
		UserId:   userId,
	}

	response, err := p.pr.AddPhoto(photo)

	if err != nil {
		return nil, err
	}

	result := &dto.AddNewPhotoResponse{
		StatusCode: http.StatusCreated,
		Message:    "new photo successfully added",
		Data: dto.AddNewPhotoResponseData{
			Id:        response.Id,
			Title:     response.Title,
			Caption:   response.Caption,
			PhotoUrl:  response.PhotoUrl,
			UserId:    response.UserId,
			CreatedAt: response.CreatedAt,
		},
	}

	return result, nil

}

func (p *photoServiceImpl) GetPhotos() (*dto.GetPhotoResponse, errs.Error) {

	result, err := p.pr.GetPhotos()

	if err != nil {
		return nil, err
	}

	photoResponseData := []dto.GetPhotoResponseData{}

	for _, v := range result {
		photo := dto.GetPhotoResponseData{
			Id:        v.Id,
			Title:     v.Title,
			Caption:   v.Caption,
			PhotoUrl:  v.PhotoUrl,
			UserId:    v.UserId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			User: struct {
				Email    string "json:\"email\""
				Username string "json:\"username\""
			}{
				Email:    v.User.Email,
				Username: v.User.Username,
			},
		}

		photoResponseData = append(photoResponseData, photo)
	}

	return &dto.GetPhotoResponse{
		StatusCode: http.StatusOK,
		Message:    "photos successfully fetched",
		Data:       photoResponseData,
	}, nil
}

func (p *photoServiceImpl) UpdatePhoto(photoId int, photoPayload *dto.PhotoUpdateRequest) (*dto.UpdatePhotoResponse, errs.Error) {

	err := helper.ValidateStruct(photoPayload)

	if err != nil {
		return nil, err
	}

	photo := &entity.Photo{
		Title:    photoPayload.Title,
		Caption:  photoPayload.Caption,
		PhotoUrl: photoPayload.PhotoUrl,
	}

	response, err := p.pr.UpdatePhoto(photoId, photo)

	if err != nil {
		return nil, err
	}

	result := &dto.UpdatePhotoResponse{
		StatusCode: http.StatusOK,
		Message:    "photo has been successfully updated",
		Data: dto.UpdatePhotoResponseData{
			Id:        response.Id,
			Title:     response.Title,
			Caption:   response.Caption,
			PhotoUrl:  response.PhotoUrl,
			UserId:    response.UserId,
			UpdatedAt: response.UpdatedAt,
		},
	}

	return result, nil
}

func (p *photoServiceImpl) DeletePhoto(photoId int) (*dto.GetPhotoResponse, errs.Error) {

	err := p.pr.DeletePhoto(photoId)

	if err != nil {
		return nil, err
	}

	return &dto.GetPhotoResponse{
		StatusCode: http.StatusOK,
		Message:    "Your photo has been successfully deleted",
		Data:       nil,
	}, nil
}
