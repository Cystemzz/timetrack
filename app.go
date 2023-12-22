package main

import (
	"context"
	"github.com/Cystemzz/timetrack/tracker"
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

func (a *App) GetWorkspaces() []tracker.Workspace {
	return tracker.GetWorkspaces()
}
