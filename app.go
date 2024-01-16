package main

import (
	"context"
	"fmt"

	"runtime"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	setInitialWindowPosition(ctx)

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
func (a *App) GetOSDetails() string {
	return fmt.Sprintf("Go Compiler %s, It's show time!", runtime.GOOS)
}

func setInitialWindowPosition(ctx context.Context) {
	screenDiemension, err := wailsRuntime.ScreenGetAll(ctx)
	if err != nil {
		fmt.Println(err)
	}
	wWidth, _ := wailsRuntime.WindowGetSize(ctx)
	firstScreenDiemension := screenDiemension[0]
	wailsRuntime.WindowSetPosition(ctx, firstScreenDiemension.Size.Width-wWidth, 10)

}
