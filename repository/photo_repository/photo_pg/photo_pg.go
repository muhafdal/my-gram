package photo_pg

import (
	"database/sql"
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/repository/photo_repository"
)

type photoRepositoryImpl struct {
	db *sql.DB
}

func NewPhotoRepository(db *sql.DB) photo_repository.PhotoRepository {
	return &photoRepositoryImpl{
		db: db,
	}
}

func (photoRepo *photoRepositoryImpl) AddPhoto(photoPayload *entity.Photo) (*dto.PhotoResponse, errs.Error) {

	tx, err := photoRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	row := tx.QueryRow(addNewPhotoQuery, photoPayload.UserId, photoPayload.Title, photoPayload.Caption, photoPayload.PhotoUrl)
	var photo dto.PhotoResponse

	err = row.Scan(
		&photo.Id,
		&photo.Title,
		&photo.Caption,
		&photo.PhotoUrl,
		&photo.UserId,
		&photo.CreatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &photo, nil
}

func (photoRepo *photoRepositoryImpl) GetPhotos() ([]photo_repository.PhotoUserMapped, errs.Error) {

	photosUser := []photo_repository.PhotoUser{}
	rows, err := photoRepo.db.Query(getUserAndPhotos)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	for rows.Next() {
		photoUser := photo_repository.PhotoUser{}

		err = rows.Scan(
			&photoUser.Photo.Id,
			&photoUser.Photo.Title,
			&photoUser.Photo.Caption,
			&photoUser.Photo.PhotoUrl,
			&photoUser.Photo.UserId,
			&photoUser.Photo.CreatedAt,
			&photoUser.Photo.UpdatedAt,
			&photoUser.User.Email,
			&photoUser.User.Username,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewNotFoundError("photos not found")
			}
			return nil, errs.NewInternalServerError("something went wrong")
		}

		photosUser = append(photosUser, photoUser)
	}

	result := photo_repository.PhotoUserMapped{}
	return result.HandleMappingPhotoWithUser(photosUser), nil
}

// GetPhotoId implements photo_repository.PhotoRepository.
func (photoRepo *photoRepositoryImpl) GetPhotoId(photoId int) (*photo_repository.PhotoUserMapped, errs.Error) {

	photoUser := photo_repository.PhotoUser{}

	row := photoRepo.db.QueryRow(getUserAndPhotosById, photoId)
	err := row.Scan(
		&photoUser.Photo.Id,
		&photoUser.Photo.Title,
		&photoUser.Photo.Caption,
		&photoUser.Photo.PhotoUrl,
		&photoUser.Photo.UserId,
		&photoUser.Photo.CreatedAt,
		&photoUser.Photo.UpdatedAt,
		&photoUser.User.Email,
		&photoUser.User.Username,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("photo not found")
		}
		return nil, errs.NewInternalServerError("something went wrong")
	}

	result := photo_repository.PhotoUserMapped{}
	return result.HandleMappingPhotoWithUserByPhotoId(photoUser), nil
}

// UpdatePhoto implements photo_repository.PhotoRepository.
func (photoRepo *photoRepositoryImpl) UpdatePhoto(photoId int, photoPayload *entity.Photo) (*dto.PhotoUpdateResponse, errs.Error) {
	tx, err := photoRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	row := tx.QueryRow(UpdatePhotoQuery, photoId, photoPayload.Title, photoPayload.Caption, photoPayload.PhotoUrl)

	var photo dto.PhotoUpdateResponse

	err = row.Scan(
		&photo.Id,
		&photo.Title,
		&photo.Caption,
		&photo.PhotoUrl,
		&photo.UserId,
		&photo.UpdatedAt,
	)

	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &photo, nil
}

// DeletePhoto implements photo_repository.PhotoRepository.
func (photoRepo *photoRepositoryImpl) DeletePhoto(photoId int) errs.Error {
	tx, err := photoRepo.db.Begin()

	if err != nil {
		tx.Rollback()
		return errs.NewInternalServerError("something went wrong")
	}

	_, err = tx.Exec(deletePhotoById, photoId)

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
