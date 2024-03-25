package comment_repository

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
)

type CommentRepository interface {
	AddComment(commentPayload *entity.Comment) (*dto.NewCommentResponse, errs.Error)
	GetComments() ([]CommentUserPhotoMapped, errs.Error)
	GetCommentById(commentId int) (*CommentUserPhotoMapped, errs.Error)
	DeleteComment(commentId int) errs.Error
	UpdateComment(commentId int, commentPayload *entity.Comment) (*entity.Comment, errs.Error)
}
