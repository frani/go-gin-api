package users

import (
	"net/http"

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

	var query listUserQuery
	err := ctx.ShouldBindQuery(&query)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "bad request",
			"success": false,
		})
		return
	}

	page := int64(query.Page)
	limit := int64(query.Limit)

	listed, err := userService.List(bson.M{}, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "internal server error",
			"error":   err.Error(),
		})
		return
	}

	// Return status 200 OK.
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "docs found",
		"data": gin.H{
			"docs":          listed.Docs,
			"totalDocs":     listed.TotalDocs,
			"limit":         listed.Limit,
			"page":          listed.Page,
			"totalPages":    listed.TotalPages,
			"hasNextPage":   listed.HasNextPage,
			"nextPage":      listed.NextPage,
			"hasPrevPage":   listed.HasPrevPage,
			"prevPage":      listed.PrevPage,
			"pagingCounter": listed.PagingCounter,
		},
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

	var param getUserParam
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "bad request",
			"success": false,
		})
		return
	}
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
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
	// Define filter and update
	filter := bson.M{"_id": id}
	found, err := userService.FindOne(filter)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "document not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	// Return status 200 OK.
	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "document found",
		"success": true,
		"data":    found,
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

	var body postUserBody
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
	user, err := userService.CreateOne(body.Name, body.Lastname, body.Password, body.Email, body.Email, body.Roles)
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

	var param patchUserParam
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "bad request",
			"success": false,
		})
		return
	}

	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "bad request",
			"success": false,
			"data":    nil,
		})
		return
	}

	var body patchUserBody
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "bad request",
			"success": false,
		})
		return
	}

	// Define filter and update
	filter := bson.M{"_id": id}
	toSet := patchUserBody{
		Email:    body.Email,
		Name:     body.Name,
		Lastname: body.Lastname,
		Username: body.Username,
		Password: body.Password,
	}
	update := bson.M{"$set": toSet}
	updated, err := userService.UpdateOne(filter, update)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "document not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   nil,
		"message": "updated",
		"success": true,
		"data":    updated,
	})

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

	var param deleteUserParam
	err := ctx.ShouldBindUri(&param)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "bad request",
			"success": false,
		})
		return
	}

	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "bad request",
			"success": false,
			"data":    nil,
		})
		return
	}

	// Define filter and update
	filter := bson.M{"_id": id}
	result, err := userService.DeleteOne(filter)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   err.Error(),
			"message": "document not found",
			"success": false,
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   nil,
		"message": "deleted",
		"success": true,
		"data":    result,
	})

}
