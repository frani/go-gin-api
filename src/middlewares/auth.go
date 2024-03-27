package middlewares

import (
	"fmt"
	"net/http"

	auth "github.com/frani/go-gin-api/src/services/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorize(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		token, err := auth.Authenticate(tokenString)

		fmt.Println("tokenString", tokenString)
		fmt.Println("err", err)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   err.Error(),
				"message": "unauthorized",
				"success": false,
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   claims.Valid().Error(),
				"message": "unauthorized",
				"success": false,
			})
			return
		}

		// TODO: Consultar la base de datos para obtener los roles del usuario

		for _, role := range roles {
			if role == "admin" {
				ctx.Next()
				return
			}
		}

		ctx.JSON(http.StatusForbidden, gin.H{
			"error":   "",
			"message": "forbidden",
			"success": false,
		})

	}
}
