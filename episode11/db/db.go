package db

import (
	"github.com/arschles/go-in-5-minutes/episode11/models"
)

// DB is the generic key/value interface for databases that work with this server
type DB interface {
	// Save saves the given model under the given key. Example usage:
	//
	//	err := mydb.Save(models.NewAppKey("app123"), models.App{Name:"myapp"})
	//	// check error and continue
	Save(models.Key, models.Model) error
	// Delete deletes the model under the given key. After delete exits successfully, both the key and the model under it should be removed entirely. Example usage:
	//
	//	err := mydb.Delete(models.NewAppKey("app123"))
	//	// check error and continue
	Delete(models.Key) error
	// Get gets the model under the given key and writes it into the given model. Callers should pass a pointer to a model implementation so that this func can unmarshal the model from the database and write it into the given model. Example usage:
	//
	//	app := &models.App{}
	//	if err := mydb.Get(models.NewAppKey("app123"), app); err != nil {
	//		fmt.Println("error getting model!")
	//		return
	//	}
	//	// use app
	Get(models.Key, models.Model) error
}
