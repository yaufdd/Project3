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
