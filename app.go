package main

import (
	"SCSE/backend"
	"context"
	"fmt"
	"os"
	"path/filepath"
)

// App struct
type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetGames() []backend.GameSchema {
	return backend.GetGames()
}

func (a *App) LoadSave(gameID string) (map[string]string, error) {
	path, err := a.getSavePath(gameID)
	if err != nil {
		return nil, err
	}
	return backend.ReadINI(path)
}

func (a *App) SaveData(gameID string, data map[string]string) error {
	path, err := a.getSavePath(gameID)
	if err != nil {
		return err
	}
	return backend.WriteINI(path, data)
}

func (a *App) getSavePath(gameID string) (string, error) {
	games := backend.GetGames()
	var fileName string
	for _, g := range games {
		if g.ID == gameID {
			fileName = g.FileName
			break
		}
	}
	if fileName == "" {
		return "", fmt.Errorf("unknown game id: %s", gameID)
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "MMFApplications", fileName), nil
}
