package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/frani/go-gin-api/src/configs"
	userService "github.com/frani/go-gin-api/src/services/users"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Getusers func gets all exists users.
// @Description Get all exists users.
// @Summary get all exists users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /v1/users [get]
func ListUsers(ctx *gin.Context) {

	// Get all users.
	cursor, err := configs.DB.Collection("users").Find(context.Background(), bson.D{{}})
	if err != nil {
		// Return, if users not found.
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "error trying to find users",
			"errors":  nil,
		})
	}

	var users []bson.M
	err = cursor.All(context.Background(), &users)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "interla server error",
			"error":   err.Error(),
		})
		return
	}

	// Return status 200 OK.
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "users were not found",
		"count":   len(users),
		"data":    users,
		"error":   false,
	})
}

// GetUser func gets user by given ID or 404 error.
// @Description Get user by given ID.
// @Summary get user by given ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /v1/user/{id} [get]
func GetUser(ctx *gin.Context) {
	// Catch user ID from URL.
	idStr := ctx.Param("id")
	fmt.Println(idStr)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "bad request",
			"success": false,
			"data":    nil,
		})
		return
	}

	// Get user by ID.
	var user bson.M
	err = configs.DB.Collection("users").FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).Decode(&user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "user with the given ID is not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	// Return status 200 OK.
	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": nil,
		"success": true,
		"data":    user,
	})
}

// CreateUser func for creates a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param user_attrs body models.UserAttrs true "User attributes"
// @Success 200 {object} models.User
// @Security ApiKeyAuth
// @Router /v1/user [post]
func PostUser(ctx *gin.Context) {

	var body postUserJSON
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "bad request",
			"success": false,
		})
	}

	// Create new User struct
	user, err := userService.CreateUser(body.Name, body.Lastname, body.Password, body.Email, body.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"success": false,
			"message": "bad request",
		})
	}

	// Return status 200 OK.
	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"success": true,
		"message": "user created",
		"data":    user,
	})
}

// UpdateUser func for updates user by given ID.
// @Description Update user.
// @Summary update user
// @Tags User
// @Accept json
// @Produce json
// @Param id body string true "User ID"
// @Param title body string true "Title"
// @Param author body string true "Author"
// @Param user_status body integer true "User status"
// @Param user_attrs body models.UserAttrs true "User attributes"
// @Success 201 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/user [put]
// func PatchUser(ctx *gin.Context) {
// 	// Get now time.
// 	now := time.Now().Unix()

// 	// Get claims from JWT.
// 	claims, err := utils.ExtractTokenMetadata(c)
// 	if err != nil {
// 		// Return status 500 and JWT parse error.
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": true,
// 			"message":   err.Error(),
// 		})
// 	}

// 	// Set expiration time from JWT data of current user.
// 	expires := claims.Expires

// 	// Checking, if now time greather than expiration from JWT.
// 	if now > expires {
// 		// Return status 401 and unauthorized error message.
// 		ctx.JSON(http.StatusUnauthorized, gin.H{
// 			"error": true,
// 			"message":   "unauthorized, check expiration time of your token",
// 		})
// 	}

// 	// Create new User struct
// 	user := &models.User{}

// 	// Check, if received JSON data is valid.
// 	if err := ctx.BindJSON(user); err != nil {
// 		// Return status 400 and error message.
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": true,
// 			"message":   err.Error(),
// 		})
// 	}

// 	// Create database connection.
// 	db, err := OpenDBConnection()
// 	if err != nil {
// 		// Return status 500 and database connection error.
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": true,
// 			"message":   err.Error(),
// 		})
// 	}

// 	// Checking, if user with given ID is exists.
// 	foundedUser, err := db.GetUser(user.ID)
// 	if err != nil {
// 		// Return status 404 and user not found error.
// 		ctx.JSON(http.StatusNotFound, gin.H{
// 			"error": true,
// 			"message":   "user with this ID not found",
// 		})
// 	}

// 	// Set initialized default data for user:
// 	user.UpdatedAt = time.Now()

// 	// Create a new validator for a User model.
// 	validate := utils.NewValidator()

// 	// Validate user fields.
// 	if err := validate.Struct(user); err != nil {
// 		// Return, if some fields are not valid.
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": true,
// 			"message":   utils.ValidatorErrors(err),
// 		})
// 	}

// 	// Update user by given ID.
// 	if err := db.UpdateUser(foundedUser.ID, user); err != nil {
// 		// Return status 500 and error message.
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": true,
// 			"message":   err.Error(),
// 		})
// 	}

// 	// Return status 201.
// 	return ctx.SendStatus(http.StatusCreated)
// }

// DeleteUser func for deletes user by given ID.
// @Description Delete user by given ID.
// @Summary delete user by given ID
// @Tags User
// @Accept json
// @Produce json
// @Param id body string true "User ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/user [delete]
// func DeleteUser(ctx *gin.Context) {
// 	// Get now time.
// 	now := time.Now().Unix()

// 	// Get claims from JWT.
// 	claims, err := utils.ExtractTokenMetadata(c)
// 	if err != nil {
// 		// Return status 500 and JWT parse error.
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": true,
// 			"message":   err.Error(),
// 		})
// 	}

// 	// Set expiration time from JWT data of current user.
// 	expires := claims.Expires

// 	// Checking, if now time greather than expiration from JWT.
// 	if now > expires {
// 		// Return status 401 and unauthorized error message.
// 		ctx.JSON(http.StatusUnauthorized, gin.H{
// 			"error": true,
// 			"message":   "unauthorized, check expiration time of your token",
// 		})
// 	}

// 	// Create new User struct
// 	user := &models.User{}

// 	// Check, if received JSON data is valid.
// 	if err := ctx.BodyParser(user); err != nil {
// 		// Return status 400 and error message.
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": true,
// 			"message":   err.Error(),
// 		})
// 	}

// 	// Create a new validator for a User model.
// 	validate := utils.NewValidator()

// 	// Validate only one user field ID.
// 	if err := validate.StructPartial(user, "id"); err != nil {
// 		// Return, if some fields are not valid.
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": true,
// 			"message":   utils.ValidatorErrors(err),
// 		})
// 	}

// 	// Create database connection.
// 	db, err := OpenDBConnection()
// 	if err != nil {
// 		// Return status 500 and database connection error.
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": true,
// 			"message":   err.Error(),
// 		})
// 	}

// 	// Checking, if user with given ID is exists.
// 	foundedUser, err := db.GetUser(user.ID)
// 	if err != nil {
// 		// Return status 404 and user not found error.
// 		ctx.JSON(http.StatusNotFound, gin.H{
// 			"error": true,
// 			"message":   "user with this ID not found",
// 		})
// 	}

// 	// Delete user by given ID.
// 	if err := db.DeleteUser(foundedUser.ID); err != nil {
// 		// Return status 500 and error message.
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": true,
// 			"message":   err.Error(),
// 		})
// 	}

// 	// Return status 204 no content.
// 	return ctx.SendStatus(http.StatusNoContent)
// }
