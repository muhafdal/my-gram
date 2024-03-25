package handler

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/pkg/helper"
	"myGram/service/photo_service"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	ps photo_service.PhotoService
}

func NewPhotoHandler(photoService photo_service.PhotoService) *photoHandler {
	return &photoHandler{
		ps: photoService,
	}
}

// AddPhoto implements PhotoHandler.
// AddPhoto godoc
// @Summary Add new photo
// @Description Add new photo
// @Tags Photos
// @Accept json
// @Produce json
// @Param dto.NewPhotoRequest body dto.NewPhotoRequest true "body request for add new photo"
// @Success 201 {object} dto.GetPhotoResponse
// @Router /photos [post]
// @Security BearerAuth
func (p *photoHandler) AddPhoto(ctx *gin.Context) {
	user, ok := ctx.MustGet("userData").(entity.User)

	if !ok {
		internalServerErr := errs.NewInternalServerError("something went wrong")
		ctx.AbortWithStatusJSON(internalServerErr.Status(), internalServerErr)
		return
	}

	photoPayload := &dto.NewPhotoRequest{}

	if err := ctx.ShouldBindJSON(photoPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := p.ps.AddPhoto(user.Id, photoPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// DeletePhoto godoc
// @Summary Delete photo
// @Description Delete photo
// @Tags Photos
// @Accept json
// @Produce json
// @Param photoId path int true "photoId"
// @Success 200 {object} dto.GetPhotoResponse
// @Router /photos/{photoId} [delete]
// @Security BearerAuth
func (p *photoHandler) DeletePhoto(ctx *gin.Context) {
	photoId, errParam := helper.GetParamId(ctx, "photoId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	response, err := p.ps.DeletePhoto(photoId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// GetPhotos godoc
// @Summary Get photos
// @Description Get photos
// @Tags Photos
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetPhotoResponse
// @Router /photos [get]
// @Security BearerAuth
func (p *photoHandler) GetPhotos(ctx *gin.Context) {
	response, err := p.ps.GetPhotos()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// UpdatePhoto implements PhotoHandler.
// UpdatePhoto godoc
// @Summary Update photo
// @Description Update photo
// @Tags Photos
// @Accept json
// @Produce json
// @Param photoId path int true "photoId"
// @Param dto.PhotoUpdateRequest body dto.PhotoUpdateRequest true "body request for update photo"
// @Success 200 {object} dto.GetPhotoResponse
// @Router /photos/{photoId} [put]
// @Security BearerAuth
func (p *photoHandler) UpdatePhoto(ctx *gin.Context) {
	photoId, errParam := helper.GetParamId(ctx, "photoId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	photoPayload := &dto.PhotoUpdateRequest{}

	if err := ctx.ShouldBindJSON(photoPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := p.ps.UpdatePhoto(photoId, photoPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
