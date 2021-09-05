package database

type Database interface {
	repository
	initializer
	mocks
}
