package database

type initializer interface {
	InitializeDatabase() error
}
