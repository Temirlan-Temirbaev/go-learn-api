package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	AuthorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (handler *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(AuthorizationHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Not authorized")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid token")
		return
	}
	userId, err := handler.services.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId)
}
