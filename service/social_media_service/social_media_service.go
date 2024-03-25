package social_media_service

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/pkg/helper"
	"myGram/repository/social_media_repository"
	"net/http"
)

type SocialMediaService interface {
	AddSocialMedia(userId int, socialMediaPayload *dto.NewSocialMediaRequest) (*dto.GetSocialMediaResponse, errs.Error)
	GetSocialMedias() (*dto.GetSocialMediaHttpResponse, errs.Error)
	UpdateSocialMedia(socialMediaId int, socialMediaPayload *dto.UpdateSocialMediaRequest) (*dto.GetSocialMediaResponse, errs.Error)
	DeleteSocialMedia(socialMediaId int) (*dto.GetSocialMediaResponse, errs.Error)
}

type socialMediaServiceImpl struct {
	sr social_media_repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepo social_media_repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaServiceImpl{
		sr: socialMediaRepo,
	}
}

// AddSocialMedia implements SocialMediaService.
func (s *socialMediaServiceImpl) AddSocialMedia(userId int, socialMediaPayload *dto.NewSocialMediaRequest) (*dto.GetSocialMediaResponse, errs.Error) {

	err := helper.ValidateStruct(socialMediaPayload)

	if err != nil {
		return nil, err
	}

	socialMedia := &entity.SocialMedia{
		Name:           socialMediaPayload.Name,
		SocialMediaUrl: socialMediaPayload.SocialMediaUrl,
		UserId:         userId,
	}

	data, err := s.sr.AddSocialMedia(socialMedia)

	if err != nil {
		return nil, err
	}

	return &dto.GetSocialMediaResponse{
		StatusCode: http.StatusCreated,
		Message:    "new social media successfully added",
		Data:       data,
	}, nil
}

// DeleteSocialMedia implements SocialMediaService.
func (s *socialMediaServiceImpl) DeleteSocialMedia(socialMediaId int) (*dto.GetSocialMediaResponse, errs.Error) {

	err := s.sr.DeleteSocialMedia(socialMediaId)

	if err != nil {
		return nil, err
	}

	return &dto.GetSocialMediaResponse{
		StatusCode: http.StatusOK,
		Message:    "Your social media has been successfully deleted",
		Data:       nil,
	}, nil
}

// GetSocialMedias implements SocialMediaService.
func (s *socialMediaServiceImpl) GetSocialMedias() (*dto.GetSocialMediaHttpResponse, errs.Error) {

	socialMedia, err := s.sr.GetSocialMedias()

	if err != nil {
		return nil, err
	}

	return &dto.GetSocialMediaHttpResponse{
		StatusCode:  http.StatusOK,
		Message:     "social medias successfully fetched",
		SocialMedia: socialMedia,
	}, nil
}

// UpdateSocialMedia implements SocialMediaService.
func (s *socialMediaServiceImpl) UpdateSocialMedia(socialMediaId int, socialMediaPayload *dto.UpdateSocialMediaRequest) (*dto.GetSocialMediaResponse, errs.Error) {

	err := helper.ValidateStruct(socialMediaPayload)

	if err != nil {
		return nil, err
	}

	socialMedia := &entity.SocialMedia{
		Name:           socialMediaPayload.Name,
		SocialMediaUrl: socialMediaPayload.SocialMediaUrl,
	}

	data, err := s.sr.UpdateSocialMedia(socialMediaId, socialMedia)

	if err != nil {
		return nil, err
	}

	return &dto.GetSocialMediaResponse{
		StatusCode: http.StatusOK,
		Message:    "social media successfully updated",
		Data:       data,
	}, nil
}
