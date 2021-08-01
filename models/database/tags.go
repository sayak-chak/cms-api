package models

//TODO: add more tags as needed

type Action struct {
	ContentId int      `pg:",pk"`         //PK & FK id
	Content   *Content `pg:"rel:has-one"` //FK reference
}

type Adventure struct {
	ContentId int      `pg:",pk"`
	Content   *Content `pg:"rel:has-one"`
}

type Horror struct {
	ContentId int      `pg:",pk"`
	Content   *Content `pg:"rel:has-one"`
}

type Drama struct {
	ContentId int      `pg:",pk"`
	Content   *Content `pg:"rel:has-one"`
}
