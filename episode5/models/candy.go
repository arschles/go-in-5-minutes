package models

import "encoding/json"

type Candy struct {
	Name string `bson:"name"`
}

func (c Candy) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{"name": c.Name})
}
