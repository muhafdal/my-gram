package auth_service

import (
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/pkg/helper"
	"myGram/repository/comment_repository"
	"myGram/repository/photo_repository"
	"myGram/repository/social_media_repository"
	"myGram/repository/user_repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	AuthorizationPhoto() gin.HandlerFunc
	AuthorizationComment() gin.HandlerFunc
	AuthorizationSocialMedia() gin.HandlerFunc
}

type authServiceImpl struct {
	ur user_repository.UserRepository
	pr photo_repository.PhotoRepository
	cr comment_repository.CommentRepository
	sr social_media_repository.SocialMediaRepository
}

func NewAuthService(userRepo user_repository.UserRepository, photoRepo photo_repository.PhotoRepository, commentRepo comment_repository.CommentRepository, socialMediaRepo social_media_repository.SocialMediaRepository) AuthService {
	return &authServiceImpl{
		ur: userRepo,
		pr: photoRepo,
		cr: commentRepo,
		sr: socialMediaRepo,
	}
}

// Authentication implements a.
func (a *authServiceImpl) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		invalidToken := errs.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		user := entity.User{}

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		_, err = a.ur.FetchById(user.Id)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidToken.Status(), invalidToken)
			return
		}

		ctx.Set("userData", user)
		ctx.Next()
	}
}

func (a *authServiceImpl) AuthorizationPhoto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, ok := ctx.MustGet("userData").(entity.User)

		if !ok {
			internalServerErr := errs.NewInternalServerError("something went wrong")
			ctx.AbortWithStatusJSON(internalServerErr.Status(), internalServerErr)
			return
		}

		photoId, _ := strconv.Atoi(ctx.Param("photoId"))

		photo, err := a.pr.GetPhotoId(photoId)

		if err != nil {

			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if photo.UserId != user.Id {
			errUnathorized := errs.NewUnathorizedError("you are not authorized to modify the photo")
			ctx.AbortWithStatusJSON(errUnathorized.Status(), errUnathorized)
		}

		ctx.Next()
	}
}

// AuthorizationComment implements a.
func (a *authServiceImpl) AuthorizationComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet("userData").(entity.User)

		commentId, errParam := helper.GetParamId(ctx, "commentId")

		if errParam != nil {
			ctx.AbortWithStatusJSON(errParam.Status(), errParam)
			return
		}
		comment, err := a.cr.GetCommentById(commentId)

		if err != nil {

			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if comment.UserId != user.Id {
			errUnathorized := errs.NewUnathorizedError("you are not authorized to modify the comment")
			ctx.AbortWithStatusJSON(errUnathorized.Status(), errUnathorized)
		}

		ctx.Next()
	}
}

// AuthorizationSocialMedia implements AuthService.
func (a *authServiceImpl) AuthorizationSocialMedia() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user, ok := ctx.MustGet("userData").(entity.User)

		if !ok {
			internalServerErr := errs.NewInternalServerError("something went wrong")
			ctx.AbortWithStatusJSON(internalServerErr.Status(), internalServerErr)
			return
		}

		socialMediaId, errParam := helper.GetParamId(ctx, "socialMediaId")

		if errParam != nil {
			ctx.AbortWithStatusJSON(errParam.Status(), errParam)
			return
		}

		socialMedia, err := a.sr.GetSocialMediaById(socialMediaId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if socialMedia.UserId != user.Id {
			errUnathorized := errs.NewUnathorizedError("you are not authorized to modify this comment")
			ctx.AbortWithStatusJSON(errUnathorized.Status(), errUnathorized)
		}

		ctx.Next()
	}
}
