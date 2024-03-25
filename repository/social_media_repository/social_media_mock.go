package social_media_repository

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
)

type socialMediaMock struct {
}

var (
	AddSocialMedia    func(socialMediaPayload *entity.SocialMedia) (*dto.NewSocialMediaResponse, errs.Error)
	DeleteSocialMedia func(socialMediaId int) errs.Error
	UpdateSocialMedia func(socialMediaId int, socialMediaPayload *entity.SocialMedia) (*dto.SocialMediaUpdateResponse, errs.Error)
	GetSocialMediaById func(socialMediaId int) (*dto.GetSocialMedia, errs.Error)
	GetSocialMedias func() ([]*dto.GetSocialMedia, errs.Error)
)

func NewSocialMediaMock() SocialMediaRepository {
	return &socialMediaMock{}
}

// AddSocialMedia implements SocialMediaRepository.
func (s *socialMediaMock) AddSocialMedia(socialMediaPayload *entity.SocialMedia) (*dto.NewSocialMediaResponse, errs.Error) {
	return AddSocialMedia(socialMediaPayload)
}

// DeleteSocialMedia implements SocialMediaRepository.
func (s *socialMediaMock) DeleteSocialMedia(socialMediaId int) errs.Error {
	return DeleteSocialMedia(socialMediaId)
}

// UpdateSocialMedia implements SocialMediaRepository.
func (s *socialMediaMock) UpdateSocialMedia(socialMediaId int, socialMediaPayload *entity.SocialMedia) (*dto.SocialMediaUpdateResponse, errs.Error) {
	return UpdateSocialMedia(socialMediaId, socialMediaPayload)
}

// GetSocialMediaById implements SocialMediaRepository.
func (s *socialMediaMock) GetSocialMediaById(socialMediaId int) (*dto.GetSocialMedia, errs.Error) {
	return GetSocialMediaById(socialMediaId)
}

// GetSocialMedias implements SocialMediaRepository.
func (s *socialMediaMock) GetSocialMedias() ([]*dto.GetSocialMedia, errs.Error) {
	return GetSocialMedias()
}
