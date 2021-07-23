package models

type TopContentsResponse struct {
	Name          string `json:"author"`
	Id            int    `json:"contentId"`
	AuthorCredsId int    `json:"authorId"`
	Content       string `json:"body"`
	ImageSrc      string `json:"imageSrc"`
	Title         string `json:"title"`
	Summary       string `json:"summary"`
	Votes         int    `json:"votes"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	AuthorId int    `json:"id"`
}

type ReadContentResponse struct {
	Id      int    `json:"contentId"`
	Content string `json:"body"`
	ImageSrc      string `json:"imageSrc"`
	Title         string `json:"title"`
	Summary       string `json:"summary"`
	Votes         int    `json:"votes"`
	Name          string `json:"author"`
	AuthorCredsId int    `json:"authorId"`
}
