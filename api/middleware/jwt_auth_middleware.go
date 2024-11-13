package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if len(authHeader) == 0 || !strings.HasPrefix(authHeader, "Bearer") {
			handleInvalidAuthentication(c, fmt.Sprintf("Invalid Authorization Header: {'%s'}", authHeader))
			return
		}

		tokenStr := strings.TrimSpace(strings.SplitN(authHeader, " ", 2)[1])
		authorized, err := tokenutil.IsAuthorized(tokenStr, secret)
		if !authorized {
			handleInvalidAuthentication(c, err.Error())
			return
		}

		userID, err := tokenutil.ExtractIDFromToken(tokenStr, secret)
		if err != nil {
			handleInvalidAuthentication(c, err.Error())
			return
		}

		c.Set("x-user-id", userID)
		c.Next()
	}
}

func handleInvalidAuthentication(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: msg})
	c.Abort()
}
