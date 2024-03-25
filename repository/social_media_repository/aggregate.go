package social_media_repository

import (
	"myGram/dto"
	"myGram/entity"
)

type SocialMediaUserPhotoMapped struct {
	SocialMedia entity.SocialMedia
	User        entity.User
	Photo       entity.Photo
}

type SocialMediaUserPhoto struct {
	SocialMedia entity.SocialMedia
	User        entity.User
	Photo       entity.Photo
}

func (s *SocialMediaUserPhotoMapped) HandleMappingSocialMediaWithUserAndPhoto(socialMediaUserPhoto []SocialMediaUserPhoto) []*dto.GetSocialMedia {
	socialMediasWithUserAndPhoto := []*dto.GetSocialMedia{}

	for _, eachSocialMediaWithUserAndPhoto := range socialMediaUserPhoto {
		socialMediaWithUserAndPhoto := &dto.GetSocialMedia{
			Id:             eachSocialMediaWithUserAndPhoto.SocialMedia.Id,
			Name:           eachSocialMediaWithUserAndPhoto.SocialMedia.Name,
			SocialMediaUrl: eachSocialMediaWithUserAndPhoto.SocialMedia.SocialMediaUrl,
			UserId:         eachSocialMediaWithUserAndPhoto.SocialMedia.UserId,
			CreatedAt:      eachSocialMediaWithUserAndPhoto.SocialMedia.CreatedAt,
			UpdatedAt:      eachSocialMediaWithUserAndPhoto.SocialMedia.UpdatedAt,
			User: dto.SocialMediaUser{
				Id:              eachSocialMediaWithUserAndPhoto.User.Id,
				Username:        eachSocialMediaWithUserAndPhoto.User.Username,
				ProfileImageUrl: eachSocialMediaWithUserAndPhoto.Photo.PhotoUrl,
			},
		}

		socialMediasWithUserAndPhoto = append(socialMediasWithUserAndPhoto, socialMediaWithUserAndPhoto)
	}

	return socialMediasWithUserAndPhoto
}

func (s *SocialMediaUserPhotoMapped) HandleMappingSocialMediaWithUserAndPhotoById(socialMediaUserPhoto SocialMediaUserPhoto) *dto.GetSocialMedia {

		socialMediaWithUserAndPhoto := &dto.GetSocialMedia{
			Id:             socialMediaUserPhoto.SocialMedia.Id,
			Name:           socialMediaUserPhoto.SocialMedia.Name,
			SocialMediaUrl: socialMediaUserPhoto.SocialMedia.SocialMediaUrl,
			UserId:         socialMediaUserPhoto.SocialMedia.UserId,
			CreatedAt:      socialMediaUserPhoto.SocialMedia.CreatedAt,
			UpdatedAt:      socialMediaUserPhoto.SocialMedia.UpdatedAt,
			User: dto.SocialMediaUser{
				Id:              socialMediaUserPhoto.User.Id,
				Username:        socialMediaUserPhoto.User.Username,
				ProfileImageUrl: socialMediaUserPhoto.Photo.PhotoUrl,
			},
		}

	return socialMediaWithUserAndPhoto
}
