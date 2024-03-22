package users

type listUserQuery struct {
	SortOrder string `form:"sortOrder" binding:"omitempty,oneof=asc des"`
	SortIndex string `form:"sortIndex" binding:"omitempty,alpha"`
	Page      uint   `form:"page" binding:"omitempty,numeric,min=1"`
	Limit     uint   `form:"limit" binding:"omitempty,numeric,min=10,max=200"`
}

type postUserBody struct {
	Username string `json:"username" binding:"required,min=3,max=60"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=10,max=60"`
	Name     string `json:"name" binding:"required,min=3,max=60"`
	Lastname string `json:"lastname" binding:"required,min=3,max=60"`
}
type patchUserBody struct {
	Username string `json:"username" bson:"username,omitempty" binding:"omitempty,min=3,max=60"`
	Email    string `json:"email" bson:"email,omitempty" binding:"omitempty,email"`
	Password string `json:"password" bson:"password,omitempty" binding:"omitempty,min=10,max=60"`
	Name     string `json:"name" bson:"name,omitempty" binding:"omitempty,min=3,max=60"`
	Lastname string `json:"lastname" bson:"lastname,omitempty" binding:"omitempty,min=3,max=60"`
}
type getUserParam struct {
	Id string `uri:"id" binding:"required,mongodb"`
}
type deleteUserParam struct {
	Id string `uri:"id" binding:"required,mongodb"`
}
