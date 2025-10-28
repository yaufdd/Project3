package request

type NewDescription struct {
	ID          int    `json:"id"`
	Description string `json:"desc"`
}

type LikePost struct {
	UserID int `json:"user_id"`
	PostID int `json:"post_id"`
}

type GetUserID struct {
	Username string `json:"username"`
}

type GetCompanyID struct {
	CompanyName string `json:"company_name"`
}

type GetProjectID struct {
	Title string `json:"title"`
}

type AuthCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=admin investor founder"`
}
