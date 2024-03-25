package handler

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/pkg/helper"
	"myGram/service/social_media_service"

	"github.com/gin-gonic/gin"
)

type socialMediaHandler struct {
	ss social_media_service.SocialMediaService
}

func NewSocialMediasHandler(socialMediaService social_media_service.SocialMediaService) *socialMediaHandler {
	return &socialMediaHandler{
		ss: socialMediaService,
	}
}

// AddSocialMedia godoc
// @Summary Add new social media
// @Description Add new social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param dto.NewSocialMediaRequest body dto.NewSocialMediaRequest true "body request for add new social media"
// @Success 201 {object} dto.GetSocialMediaResponse
// @Router /socialmedias [post]
// @Security BearerAuth
func (s *socialMediaHandler) AddSocialMedia(ctx *gin.Context) {

	socialMediaPayload := &dto.NewSocialMediaRequest{}
	user, ok := ctx.MustGet("userData").(entity.User)

	if !ok {
		internalServerErr := errs.NewInternalServerError("something went wrong")
		ctx.AbortWithStatusJSON(internalServerErr.Status(), internalServerErr)
		return
	}

	if err := ctx.ShouldBindJSON(socialMediaPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := s.ss.AddSocialMedia(user.Id, socialMediaPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// DeleteSocialMedia godoc
// @Summary Delete social media
// @Description Delete social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param socialMediaId path int true "socialMediaId"
// @Success 200 {object} dto.GetSocialMediaResponse
// @Router /socialmedias/{socialMediaId} [delete]
// @Security BearerAuth
func (s *socialMediaHandler) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId, errParam := helper.GetParamId(ctx, "socialMediaId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	response, err := s.ss.DeleteSocialMedia(socialMediaId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// GetSocialMedias godoc
// @Summary Get social medias
// @Description Get social medias
// @Tags Social Media
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetSocialMediaHttpResponse
// @Router /socialmedias [get]
// @Security BearerAuth
func (s *socialMediaHandler) GetSocialMedias(ctx *gin.Context) {
	response, err := s.ss.GetSocialMedias()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// UpdateSocialMedia godoc
// @Summary Update social media
// @Description Update social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param socialMediaId path int true "socialMediaId"
// @Param dto.UpdateSocialMediaRequest body dto.UpdateSocialMediaRequest true "body request for update social media"
// @Success 200 {object} dto.GetSocialMediaResponse
// @Router /socialmedias/{socialMediaId} [put]
// @Security BearerAuth
func (s *socialMediaHandler) UpdateSocialMedia(ctx *gin.Context) {

	socialMediaPayload := &dto.UpdateSocialMediaRequest{}
	socialMediaId, errParam := helper.GetParamId(ctx, "socialMediaId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	if err := ctx.ShouldBindJSON(socialMediaPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := s.ss.UpdateSocialMedia(socialMediaId, socialMediaPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
