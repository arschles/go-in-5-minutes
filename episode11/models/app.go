package models

type App struct {
	key PrimaryKey
}

func (a *App) Key() PrimaryKey {
	return a.key
}
