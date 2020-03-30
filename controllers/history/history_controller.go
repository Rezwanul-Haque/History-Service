package history

import (
	"github.com/gin-gonic/gin"
	"github.com/rezwanul-haque/History-Service/domain/history"
	"github.com/rezwanul-haque/History-Service/services"
	"github.com/rezwanul-haque/History-Service/utils/errors"
	"net/http"
)

func Get(c *gin.Context) {
	domain := c.GetHeader("RLS-Referrer")
	if domain == "" {
		restErr := errors.NewBadRequestError("RLS-Referrer header is not present")
		c.JSON(restErr.Status, restErr)
		return
	}
	var queryParams history.QueryParamRequest
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		if queryParamsErr := queryParams.Validate(); queryParamsErr != nil {
			c.JSON(queryParamsErr.Status, queryParamsErr)
			return
		}
	}

	response, getErr := services.HistoryService.GetHistory(domain, queryParams.UserId, *queryParams.StartDate, *queryParams.EndDate)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, response)
}
