package models

import (
	"encoding/json"
	"time"
)

type AppKey struct {
	str string
}

func NewAppKey(name string) *AppKey {
	return &AppKey{str: name}
}
func (a *AppKey) String() string {
	return a.str
}

type App struct {
	Name         string    `json:"name"`
	MaxInstances int       `json:"max_instances"`
	LastDeploy   time.Time `json:"last_deploy"`
}

func (a *App) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

func (a *App) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, a)
}
