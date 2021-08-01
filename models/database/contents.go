package models

type Content struct {
	Id            int `pg:",pk"`
	AuthorCredsId int
	Content       string
	ImageSrc      string
	Title         string
	Summary       string
	Votes         int          // indexed here
	AuthorCreds   *AuthorCreds `pg:"rel:has-one"`
}
