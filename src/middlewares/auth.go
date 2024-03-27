package middlewares

import (
	"net/http"
	"strings"

	auth "github.com/frani/go-gin-api/src/services/auth"
	utils "github.com/frani/go-gin-api/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorize(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.Request.Header.Get("Authorization")
		tokenSplitted := strings.Split(authorization, " ")
		if len(tokenSplitted) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "",
				"message": "unauthorized",
				"success": false,
			})
			return
		}

		token, err := auth.Authenticate(tokenSplitted[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   err.Error(),
				"message": "unauthorized",
				"success": false,
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   claims.Valid().Error(),
				"message": "unauthorized",
				"success": false,
			})
			return
		}

		claimsRoles := utils.ConvertInterfacesToSlice((claims["roles"].([]interface{})))

		if utils.SomeElementInSlice(claimsRoles, roles) {
			ctx.Next()
			return
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error":   "",
			"message": "forbidden",
			"success": false,
		})

	}
}
