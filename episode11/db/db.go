package db

import (
	"github.com/arschles/go-in-5-minutes/episode11/models"
)

type DB interface {
	Save(models.Key, models.Model) error
	Delete(models.Key) error
	Get(models.Key, models.Model) error
}
