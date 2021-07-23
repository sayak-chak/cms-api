package models

type AddContentRequest struct {
	Tags     []string `json:"tags"`
	Title    string   `json:"title"`
	Summary  string   `json:"summary"`
	Body     string   `json:"body"`
	ImageSrc string   `json:"imageSrc"`
	AuthorId int      `json:"authorId"`
}

type SubscriberRequest struct { //TODO: tag based indexing if needed
	Email string `json:"email"`
}

type NewAuthorAccCreationRequest struct {
	Otp      string `json:"otp"`
	Password string `json:"password"`
	Mobile   int    `json:"mobile"`
	Name     string `json:"name"`
}

type LoginRequest struct {
	Password string `json:"password"`
	Mobile   int    `json:"mobile"`
}

type UpvoteRequest struct {
	ContentId int `json:"contentId"`
}
