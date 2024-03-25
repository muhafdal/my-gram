package photo_repository

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
)

var (
	AddPhoto    func(photoPayload *entity.Photo) (*dto.PhotoResponse, errs.Error)
	GetPhotos   func() ([]PhotoUserMapped, errs.Error)
	GetPhotoId  func(photoId int) (*PhotoUserMapped, errs.Error)
	UpdatePhoto func(photoId int, photoPayload *entity.Photo) (*dto.PhotoUpdateResponse, errs.Error)
	DeletePhoto func(photoId int) errs.Error
)

type photoRepositoryMock struct {
}

func NewPhotoRepositoryMock() PhotoRepository {
	return &photoRepositoryMock{}
}

// AddPhoto implements PhotoRepository.
func (prm *photoRepositoryMock) AddPhoto(photoPayload *entity.Photo) (*dto.PhotoResponse, errs.Error) {
	return AddPhoto(photoPayload)
}

// GetPhotos implements PhotoRepository.
func (prm *photoRepositoryMock) GetPhotos() ([]PhotoUserMapped, errs.Error) {
	return GetPhotos()
}

// GetPhotoId implements PhotoRepository.
func (prm *photoRepositoryMock) GetPhotoId(photoId int) (*PhotoUserMapped, errs.Error) {
	return GetPhotoId(photoId)
}

// UpdatePhoto implements PhotoRepository.
func (prm *photoRepositoryMock) UpdatePhoto(photoId int, photoPayload *entity.Photo) (*dto.PhotoUpdateResponse, errs.Error) {
	return UpdatePhoto(photoId, photoPayload)
}

// DeletePhoto implements PhotoRepository.
func (prm *photoRepositoryMock) DeletePhoto(photoId int) errs.Error {
	return DeletePhoto(photoId)
}
