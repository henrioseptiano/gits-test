package repository

import (
	"gits-test/app"
)

type AppRepository struct {
	DB app.Database
}

func NewAppRepository(db app.Database) *AppRepository {
	return &AppRepository{DB: db}
}
