package comment_pg

import (
	"database/sql"
	"fmt"
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/comment_repository"
)

type commentRepositoryImpl struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) comment_repository.CommentRepository {
	return &commentRepositoryImpl{
		db: db,
	}
}

func (commentRepo *commentRepositoryImpl) AddComment(commentPayload *entity.Comment) (*dto.NewCommentResponse, errs.Error) {
	tx, err := commentRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	var comment dto.NewCommentResponse
	err = tx.QueryRow(
		addCommentQuery,
		commentPayload.UserId,
		commentPayload.PhotoId,
		commentPayload.Message,
	).Scan(
		&comment.Id,
		&comment.Message,
		&comment.PhotoId,
		&comment.UserId,
		&comment.CreatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &comment, nil
}

func (commentRepo *commentRepositoryImpl) GetComments() ([]comment_repository.CommentUserPhotoMapped, errs.Error) {

	var commentsUserPhoto []comment_repository.CommentUserPhoto
	rows, err := commentRepo.db.Query(getCommentQuery)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	for rows.Next() {
		commentUserPhoto := comment_repository.CommentUserPhoto{}
		err = rows.Scan(
			&commentUserPhoto.Comment.Id,
			&commentUserPhoto.Comment.UserId,
			&commentUserPhoto.Comment.PhotoId,
			&commentUserPhoto.Comment.Message,
			&commentUserPhoto.Comment.CreatedAt,
			&commentUserPhoto.Comment.UpdatedAt,
			&commentUserPhoto.User.Id,
			&commentUserPhoto.User.Username,
			&commentUserPhoto.User.Email,
			&commentUserPhoto.Photo.Id,
			&commentUserPhoto.Photo.Title,
			&commentUserPhoto.Photo.Caption,
			&commentUserPhoto.Photo.PhotoUrl,
			&commentUserPhoto.Photo.UserId,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		commentsUserPhoto = append(commentsUserPhoto, commentUserPhoto)
	}

	result := comment_repository.CommentUserPhotoMapped{}
	return result.HandleMappingCommentsUserPhoto(commentsUserPhoto), nil
}

func (commentRepo *commentRepositoryImpl) GetCommentById(commentId int) (*comment_repository.CommentUserPhotoMapped, errs.Error) {

	var commentUserPhoto comment_repository.CommentUserPhoto
	rows := commentRepo.db.QueryRow(getCommentByIdQuery, commentId)

	err := rows.Scan(
		&commentUserPhoto.Comment.Id,
		&commentUserPhoto.Comment.UserId,
		&commentUserPhoto.Comment.PhotoId,
		&commentUserPhoto.Comment.Message,
		&commentUserPhoto.Comment.CreatedAt,
		&commentUserPhoto.Comment.UpdatedAt,
		&commentUserPhoto.User.Id,
		&commentUserPhoto.User.Username,
		&commentUserPhoto.User.Email,
		&commentUserPhoto.Photo.Id,
		&commentUserPhoto.Photo.Title,
		&commentUserPhoto.Photo.Caption,
		&commentUserPhoto.Photo.PhotoUrl,
		&commentUserPhoto.Photo.UserId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("comment not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	result := comment_repository.CommentUserPhotoMapped{}
	return result.HandleMappingCommentUserPhoto(commentUserPhoto), nil
}

// DeleteComment implements comment_repository.CommentRepository.
func (commentRepo *commentRepositoryImpl) DeleteComment(commentId int) errs.Error {
	tx, err := commentRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(deleteCommentQuery, commentId)

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (commentRepo *commentRepositoryImpl) UpdateComment(commentId int, commentPayload *entity.Comment) (*entity.Comment, errs.Error) {

	row := commentRepo.db.QueryRow(updateCommentQuery, commentId, commentPayload.Message)

	var comment entity.Comment

	err := row.Scan(
		&comment.Id,
		&comment.Message,
		&comment.PhotoId,
		&comment.UserId,
		&comment.UpdatedAt,
	)

	if err != nil {
		fmt.Printf("[UpdateComment-repo] err: %s\n", err.Error())
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &comment, nil
}
