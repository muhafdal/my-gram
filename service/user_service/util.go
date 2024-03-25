package user_service

import (
	"myGram/pkg/errs"
	"net/http"
)

func (u *userServiceImpl) validateDuplicateEmailAndUsername(email string, username string) errs.Error {
	usr, err := u.ur.FetchByEmail(email)

	if err != nil && err.Status() != http.StatusNotFound {
		return err
	}

	if usr != nil {
		return errs.NewConflictError("email has been used")
	}

	usr, err = u.ur.FetchByUsername(username)

	if err != nil && err.Status() != http.StatusNotFound {
		return err
	}

	if usr != nil {
		return errs.NewConflictError("username has been used")
	}
	return nil
}
