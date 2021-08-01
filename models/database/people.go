package models

type Author struct {
	AuthorCredsId int          `pg:",pk"`
	Name          string       `pg:"type:varchar(50)"`
	AuthorCreds   *AuthorCreds `pg:"rel:has-one"`
}

type AuthorCreds struct {
	Id             int `pg:",pk"`
	Mobile         int
	HashedPassword string
}

type AuthorTempCreds struct {
	Mobile       int `pg:",pk"`
	TempPassword string
	CreationTime int64
}

type Subscriber struct {
	Email string `pg:"type:varchar(320),pk"`
}
