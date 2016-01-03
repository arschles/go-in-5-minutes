package db

import (
	"github.com/arschles/go-in-5-minutes/episode11/models"
)

type DB interface {
	Save(models.Model) error
	Delete(models.Model) error
	Get(models.PrimaryKey, models.Model) error
}
