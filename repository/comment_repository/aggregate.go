package comment_repository

import (
	"myGram/entity"
	"time"
)

type CommentUserPhoto struct {
	Comment entity.Comment
	User    entity.User
	Photo   entity.Photo
}

type user struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type photo struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}

type CommentUserPhotoMapped struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	PhotoId   int       `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      user      `json:"user"`
	Photo     photo     `json:"photo"`
}

func (cup *CommentUserPhotoMapped) HandleMappingCommentsUserPhoto(commentUserPhoto []CommentUserPhoto) []CommentUserPhotoMapped {
	commentsUserPhotoMapped := []CommentUserPhotoMapped{}

	for _, eachCommentUserPhoto := range commentUserPhoto {
		commentUserPhotoMapped := CommentUserPhotoMapped{
			Id:        eachCommentUserPhoto.Comment.Id,
			UserId:    eachCommentUserPhoto.Comment.UserId,
			PhotoId:   eachCommentUserPhoto.Comment.PhotoId,
			Message:   eachCommentUserPhoto.Comment.Message,
			CreatedAt: eachCommentUserPhoto.Comment.CreatedAt,
			UpdatedAt: eachCommentUserPhoto.Comment.UpdatedAt,
			User: user{
				Id:       eachCommentUserPhoto.User.Id,
				Username: eachCommentUserPhoto.User.Username,
				Email:    eachCommentUserPhoto.User.Email,
			},
			Photo: photo{
				Id:       eachCommentUserPhoto.Photo.Id,
				Title:    eachCommentUserPhoto.Photo.Title,
				Caption:  eachCommentUserPhoto.Photo.Caption,
				PhotoUrl: eachCommentUserPhoto.Photo.PhotoUrl,
				UserId:   eachCommentUserPhoto.Photo.UserId,
			},
		}

		commentsUserPhotoMapped = append(commentsUserPhotoMapped, commentUserPhotoMapped)
	}

	return commentsUserPhotoMapped
}

func (cup *CommentUserPhotoMapped) HandleMappingCommentUserPhoto(commentUserPhoto CommentUserPhoto) *CommentUserPhotoMapped {

	commentUserPhotoMapped := CommentUserPhotoMapped{
		Id:        commentUserPhoto.Comment.Id,
		UserId:    commentUserPhoto.Comment.UserId,
		PhotoId:   commentUserPhoto.Comment.PhotoId,
		Message:   commentUserPhoto.Comment.Message,
		CreatedAt: commentUserPhoto.Comment.CreatedAt,
		UpdatedAt: commentUserPhoto.Comment.UpdatedAt,
		User: user{
			Id:       commentUserPhoto.User.Id,
			Username: commentUserPhoto.User.Username,
			Email:    commentUserPhoto.User.Email,
		},
		Photo: photo{
			Id:       commentUserPhoto.Photo.Id,
			Title:    commentUserPhoto.Photo.Title,
			Caption:  commentUserPhoto.Photo.Caption,
			PhotoUrl: commentUserPhoto.Photo.PhotoUrl,
			UserId:   commentUserPhoto.Photo.UserId,
		},
	}

	return &commentUserPhotoMapped
}
