package handler

import (
	"myGram/dto"
	"myGram/entity"
	"myGram/pkg/errs"
	"myGram/pkg/helper"
	"myGram/service/comment_service"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	cs comment_service.CommentService
}

func NewCommentHandler(commentService comment_service.CommentService) *commentHandler {
	return &commentHandler{
		cs: commentService,
	}
}

// AddComment implements CommentHandler.
// AddComment godoc
// @Summary Add new comment
// @Description Add new comment
// @Tags Comments
// @Accept json
// @Produce json
// @Param dto.NewCommentRequest body dto.NewCommentRequest true "body request for add new comment"
// @Success 201 {object} dto.GetCommentResponse
// @Router /comments [post]
// @Security BearerAuth
func (c *commentHandler) AddComment(ctx *gin.Context) {
	user, ok := ctx.MustGet("userData").(entity.User)

	if !ok {
		internalServerErr := errs.NewInternalServerError("something went wrong")
		ctx.AbortWithStatusJSON(internalServerErr.Status(), internalServerErr)
		return
	}
	commentPayload := &dto.NewCommentRequest{}

	if err := ctx.ShouldBindJSON(commentPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := c.cs.AddComment(user.Id, commentPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// DeleteComment implements CommentHandler.
// DeleteComment godoc
// @Summary Delete comment
// @Description Delete comment
// @Tags Comments
// @Accept json
// @Produce json
// @Param commentId path int true "commentId"
// @Success 200 {object} dto.GetCommentResponse
// @Router /comments/{commentId} [delete]
// @Security BearerAuth
func (c *commentHandler) DeleteComment(ctx *gin.Context) {
	commentId, errParam := helper.GetParamId(ctx, "commentId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	response, err := c.cs.DeleteComment(commentId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// GetComments implements CommentHandler.
// GetComments godoc
// @Summary Get comments
// @Description Get comments
// @Tags Comments
// @Accept json
// @Produce json
// @Success 200 {object} dto.GetCommentResponse
// @Router /comments [get]
// @Security BearerAuth
func (c *commentHandler) GetComments(ctx *gin.Context) {
	response, err := c.cs.GetComments()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

// UpdateComment implements CommentHandler.
// UpdateComment godoc
// @Summary Update comment
// @Description Update comment
// @Tags Comments
// @Accept json
// @Produce json
// @Param commentId path int true "commentId"
// @Param dto.UpdateCommentRequest body dto.UpdateCommentRequest true "body request for update comment"
// @Success 200 {object} dto.GetCommentResponse
// @Router /comments/{commentId} [put]
// @Security BearerAuth
func (c *commentHandler) UpdateComment(ctx *gin.Context) {
	commentId, errParam := helper.GetParamId(ctx, "commentId")

	if errParam != nil {
		ctx.AbortWithStatusJSON(errParam.Status(), errParam)
		return
	}

	commentPayload := &dto.UpdateCommentRequest{}

	if err := ctx.ShouldBindJSON(commentPayload); err != nil {
		errBindJson := errs.NewUnprocessableEntityError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := c.cs.UpdateComment(commentId, commentPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
