package social_media_pg

import (
	"database/sql"
	"myGram/entity"
	"myGram/repository/social_media_repository"
	"time"
)

type socialMediaWithUserAndPhoto struct {
	SocialMediaId             int
	SocialMediaName           string
	SocialMediaSocialMediaUrl string
	SocialMediaUserId         int
	SocialMediaCreatedAt      time.Time
	SocialMediaUpdatedAt      time.Time
	UserId                    int
	UserUsername              string
	UserEmail                 string
	UserPassword              string
	UserAge                   uint
	UserCreatedAt             time.Time
	UserUpdatedAt             time.Time
	PhotoId                   sql.NullInt64
	PhotoTitle                sql.NullString
	PhotoCaption              sql.NullString
	PhotoPhotoUrl             sql.NullString
	PhotoUserId               sql.NullInt64
	PhotoCreatedAt            sql.NullTime
	PhotoUpdatedAt            sql.NullTime
}

func (s *socialMediaWithUserAndPhoto) socialMediaWithUserAndPhotoToAggregate() social_media_repository.SocialMediaUserPhotoMapped {
	return social_media_repository.SocialMediaUserPhotoMapped{
		SocialMedia: entity.SocialMedia{
			Id:             s.SocialMediaId,
			Name:           s.SocialMediaName,
			SocialMediaUrl: s.SocialMediaSocialMediaUrl,
			UserId:         s.SocialMediaUserId,
			CreatedAt:      s.SocialMediaCreatedAt,
			UpdatedAt:      s.SocialMediaUpdatedAt,
		},
		User: entity.User{
			Id:        s.UserId,
			Username:  s.UserUsername,
			Email:     s.UserEmail,
			Password:  s.UserPassword,
			Age:       s.UserAge,
			CreatedAt: s.UserCreatedAt,
			UpdatedAt: s.UserUpdatedAt,
		},
		Photo: entity.Photo{
			Id:        int(s.PhotoId.Int64),
			Title:     s.PhotoTitle.String,
			Caption:   s.PhotoCaption.String,
			PhotoUrl:  s.PhotoPhotoUrl.String,
			UserId:    int(s.PhotoUserId.Int64),
			CreatedAt: s.PhotoCreatedAt.Time,
			UpdatedAt: s.PhotoUpdatedAt.Time,
		},
	}
}
