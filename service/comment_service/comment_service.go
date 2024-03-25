package comment_service

import (
	"fmt"
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/pkg/helper"
	"myGram/repository/comment_repository"
	"myGram/repository/photo_repository"
	"net/http"
)

type CommentService interface {
	AddComment(userId int, commentPayload *dto.NewCommentRequest) (*dto.GetCommentResponse, errs.Error)
	GetComments() (*dto.GetCommentResponse, errs.Error)
	DeleteComment(commentId int) (*dto.GetCommentResponse, errs.Error)
	UpdateComment(commentId int, commentPayload *dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, errs.Error)
}

type commentServiceImpl struct {
	pr photo_repository.PhotoRepository
	cr comment_repository.CommentRepository
}

func NewCommentService(commentRepo comment_repository.CommentRepository, photoRepo photo_repository.PhotoRepository) CommentService {
	return &commentServiceImpl{
		pr: photoRepo,
		cr: commentRepo,
	}
}

// AddComment implements CommentService.
func (c *commentServiceImpl) AddComment(userId int, commentPayload *dto.NewCommentRequest) (*dto.GetCommentResponse, errs.Error) {

	err := helper.ValidateStruct(commentPayload)

	if err != nil {
		return nil, err
	}

	_, err = c.pr.GetPhotoId(commentPayload.PhotoId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	comment := &entity.Comment{
		UserId:  userId,
		PhotoId: commentPayload.PhotoId,
		Message: commentPayload.Message,
	}

	response, err := c.cr.AddComment(comment)

	if err != nil {
		return nil, err
	}

	return &dto.GetCommentResponse{
		StatusCode: http.StatusCreated,
		Message:    "new comment successfully added",
		Data:       response,
	}, nil
}

// GetComments implements CommentService.
func (c *commentServiceImpl) GetComments() (*dto.GetCommentResponse, errs.Error) {

	data, err := c.cr.GetComments()

	if err != nil {
		return nil, err
	}

	return &dto.GetCommentResponse{
		StatusCode: http.StatusOK,
		Message:    "comments successfully fetched",
		Data:       data,
	}, nil
}

// DeleteComment implements CommentService.
func (c *commentServiceImpl) DeleteComment(commentId int) (*dto.GetCommentResponse, errs.Error) {

	err := c.cr.DeleteComment(commentId)

	if err != nil {
		return nil, err
	}

	return &dto.GetCommentResponse{
		StatusCode: http.StatusOK,
		Message:    "Your comment has been successfully deleted",
		Data:       nil,
	}, nil
}

func (c *commentServiceImpl) UpdateComment(commentId int, commentPayload *dto.UpdateCommentRequest) (*dto.UpdateCommentResponse, errs.Error) {

	err := helper.ValidateStruct(commentPayload)

	if err != nil {
		return nil, err
	}

	comment := &entity.Comment{
		Message: commentPayload.Message,
	}

	data, err := c.cr.UpdateComment(commentId, comment)

	if err != nil {
		fmt.Printf("[UpdateComment-service] err: %s\n", err.Message())
		return nil, err
	}

	result := &dto.UpdateCommentResponse{
		StatusCode: http.StatusOK,
		Message:    "comment has been successfully updated",
		Data: dto.UpdateCommentResponseData{
			Id:        data.Id,
			UserId:    data.UserId,
			PhotoId:   data.PhotoId,
			Message:   data.Message,
			UpdatedAt: data.UpdatedAt,
		},
	}

	return result, nil
}
