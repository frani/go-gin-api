package auth

// POST one
type postLogIn struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,max=80"`
}

type postLogUp struct {
	Username string `json:"username" binding:"required,min=3,max=60"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=10,max=60"`
	Name     string `json:"name" binding:"required,min=3,max=60"`
	Lastname string `json:"lastname" binding:"required,min=3,max=60"`
}
