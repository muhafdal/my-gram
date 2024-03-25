package social_media_repository

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
)

type SocialMediaRepository interface {
	AddSocialMedia(socialMediaPayload *entity.SocialMedia) (*dto.NewSocialMediaResponse, errs.Error)
	UpdateSocialMedia(socialMediaId int, socialMediaPayload *entity.SocialMedia) (*dto.SocialMediaUpdateResponse, errs.Error)
	GetSocialMedias() ([]*dto.GetSocialMedia, errs.Error)
	GetSocialMediaById(socialMediaId int) (*dto.GetSocialMedia, errs.Error)
	DeleteSocialMedia(socialMediaId int) errs.Error
}
