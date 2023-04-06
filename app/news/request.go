package news

type NewsRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      string `json:"author" binding:"required"`
	PhoneNumber int    `json:"phone_number" binding:"required,number"`
}

type NewsUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	PhoneNumber int    `json:"phone_number binding:number"`
}
