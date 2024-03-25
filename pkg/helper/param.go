package helper

import (
	"fmt"
	"strconv"

	"myGram/pkg/errs"

	"github.com/gin-gonic/gin"
)

func GetParamId(c *gin.Context, key string) (int, errs.Error) {
	value := c.Param(key)

	id, err := strconv.Atoi(value)

	if err != nil {
		return 0, errs.NewBadRequestError(fmt.Sprintf("parameter '%s' has to be a number", key))
	}

	return id, nil
}
