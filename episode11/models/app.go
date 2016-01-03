package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// AppKey is a models.Key implementation specifically for App models
type AppKey struct {
	str string
}

// NewAppKey creates and returns a new AppKey with the given app name
func NewAppKey(name string) *AppKey {
	return &AppKey{str: name}
}

// String is the fmt.Stringer interface implementation
func (a *AppKey) String() string {
	return a.str
}

// App is a Model implementation for an application in the PaaS
type App struct {
	Name         string    `json:"name"`
	MaxInstances int       `json:"max_instances"`
	LastDeploy   time.Time `json:"last_deploy"`
}

// MarshalBinary is the encoding.BinaryMarshaler interface implementation
func (a *App) MarshalBinary() ([]byte, error) {
	return json.Marshal(a)
}

// UnmarshalBinary is the encoding.BinaryUnmarshaler interface implementation
func (a *App) UnmarshalBinary(b []byte) error {
	return json.Unmarshal(b, a)
}

// Set is the Model interface implementation
func (a *App) Set(m Model) error {
	app, ok := m.(*App)
	if !ok {
		return fmt.Errorf("given model %+v was not an *App", m)
	}
	*a = *app
	return nil
}
