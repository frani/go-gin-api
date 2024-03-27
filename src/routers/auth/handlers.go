package auth

import (
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/frani/go-gin-api/src/services/auth"
	userService "github.com/frani/go-gin-api/src/services/users"
	"github.com/frani/go-gin-api/src/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostLogIn(ctx *gin.Context) {

	var body postLogIn
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println("err", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "bad request",
			"success": false,
		})
		return
	}

	// Create new User struct
	user, err := userService.FindOne(bson.M{
		"email": body.Email,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "",
			"success": false,
			"message": "wrong email or password",
		})
		return
	}

	HasedPassword := user["password"].(string)

	decodedHashedPassword, err := hex.DecodeString(HasedPassword)
	if err != nil {
		fmt.Println("err", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "",
			"success": false,
			"message": "internal server error",
		})
		return
	}

	varified := utils.VerifyPassword(body.Password, decodedHashedPassword)
	if !varified {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "",
			"success": false,
			"message": "wrong email or password",
		})
		return
	}

	var roles []string
	for _, value := range user["roles"].(primitive.A) {
		roles = append(roles, value.(string))
	}

	claims := auth.SignTokenClaims{
		Email:    user["email"].(string),
		Lastname: user["lastname"].(string),
		Name:     user["name"].(string),
		Roles:    roles,
	}

	jwt, err := auth.SignToken(claims)

	if err != nil {
		fmt.Println(`ERR `, ctx.Request.Method, ctx.Request.URL.Path, err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "",
			"success": false,
			"message": "internal server error",
		})
		return
	}

	// Return status 200 OK.
	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"success": true,
		"message": "ok",
		"data":    jwt,
	})
}

func PostLogUp(ctx *gin.Context) {

	var body postLogUp
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "bad request",
			"success": false,
		})
		return
	}

	// Create new User struct
	roles := []string{"user"}
	user, err := userService.CreateOne(body.Name, body.Lastname, body.Password, body.Email, body.Username, roles)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
			"message": "bad request",
		})
		return
	}

	// Return status 200 OK.
	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"success": true,
		"message": "user created",
		"data":    user,
	})
}
