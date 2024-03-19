package users

import (
	"net/http"
	"time"

	database "github.com/frani/go-gin-api/src/configs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	// Create database connection.
	db := database.DB

	// Get all users.
	users, err := db.Collection("users").Find(database.Ctx)
	if err != nil {
		// Return, if users not found.
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"data":    nil,
			"message": "error trying to find users",
			"errors":  nil,
		})
	}
	// Return status 200 OK.
	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "users were not found",
		"count": len(users),
		"users": users,
		"error": false,
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
	id := uuid.Parse(ctx.Param("id"))

	DB := database.DB

	// Get user by ID.
	user, err := DB.Collection("users").FindOne()
	if err != nil {
		// Return, if user not found.
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"msg":   "user with the given ID is not found",
			"user":  nil,
		})
	}

	// Return status 200 OK.
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   nil,
		"user":  user,
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
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current user.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new User struct
	user := &models.User{}

	// Check, if received JSON data is valid.
	if err := ctx.BodyParser(user); err != nil {
		// Return status 400 and error message.
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Set initialized default data for user:
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.userstatus = 1 // 0 == draft, 1 == active

	// Validate user fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Delete user by given ID.
	if err := db.CreateUser(user); err != nil {
		// Return status 500 and error message.
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   nil,
		"user":  user,
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
func PatchUser(ctx *gin.Context) {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current user.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new User struct
	user := &models.User{}

	// Check, if received JSON data is valid.
	if err := ctx.BindJSON(user); err != nil {
		// Return status 400 and error message.
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if user with given ID is exists.
	foundedUser, err := db.GetUser(user.ID)
	if err != nil {
		// Return status 404 and user not found error.
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"msg":   "user with this ID not found",
		})
	}

	// Set initialized default data for user:
	user.UpdatedAt = time.Now()

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Validate user fields.
	if err := validate.Struct(user); err != nil {
		// Return, if some fields are not valid.
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Update user by given ID.
	if err := db.UpdateUser(foundedUser.ID, user); err != nil {
		// Return status 500 and error message.
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 201.
	return ctx.SendStatus(http.StatusCreated)
}

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
func DeleteUser(ctx *gin.Context) {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current user.
	expires := claims.Expires

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new User struct
	user := &models.User{}

	// Check, if received JSON data is valid.
	if err := ctx.BodyParser(user); err != nil {
		// Return status 400 and error message.
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()

	// Validate only one user field ID.
	if err := validate.StructPartial(user, "id"); err != nil {
		// Return, if some fields are not valid.
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if user with given ID is exists.
	foundedUser, err := db.GetUser(user.ID)
	if err != nil {
		// Return status 404 and user not found error.
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": true,
			"msg":   "user with this ID not found",
		})
	}

	// Delete user by given ID.
	if err := db.DeleteUser(foundedUser.ID); err != nil {
		// Return status 500 and error message.
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 204 no content.
	return ctx.SendStatus(http.StatusNoContent)
}
