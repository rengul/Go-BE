package delivery

import (
	"errors"
	"net/http"
	"re-home/auth/pkg/auth"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "id"
)

type AuthMiddleware struct {
	usecase auth.UseCase
}

func NewAuthMiddleware(usecase auth.UseCase) gin.HandlerFunc {
	return (&AuthMiddleware{
		usecase: usecase,
	}).Handle
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if headerParts[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := m.usecase.ParseToken(c.Request.Context(), headerParts[1])
	if err != nil {
		status := http.StatusInternalServerError
		if err == auth.ErrInvalidAccessToken {
			status = http.StatusUnauthorized
		}

		// if err == jwt.ErrTokenExpired {
		// 	status = http.StatusUnauthorized
		// }

		c.AbortWithStatus(status)
		return
	}

	c.Set(auth.CtxUserKey, user)

}

func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return "", errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return "", errors.New("user id is not an integer")
	}

	return strconv.Itoa(idInt), nil
}
