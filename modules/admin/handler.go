package admin

import "cms-api/database"

type Handler interface {
}

func NewHandler(database database.Database) Handler {
	return &handler{
		service: &service{
			database: database,
		},
	}
}

type handler struct {
	service *service
}
