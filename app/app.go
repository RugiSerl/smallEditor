package app

import (
	"github.com/RugiSerl/smallEditor/app/math"
	"github.com/RugiSerl/smallEditor/app/settings"
	"github.com/RugiSerl/smallEditor/app/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	APP_NAME = "smallEditor"
)

var (
	myInterface *ui.Interface
)

func Run() {
	start()
	for !rl.WindowShouldClose() {
		update()
	}
	quit()
}

func start() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 450, APP_NAME)
	settings.LoadSettings()
	myInterface = ui.NewInterface()

	rl.SetTargetFPS(-1)

}

func update() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	myInterface.Update(math.NewRect(math.NewVec2(0, 0), math.NewVec2(float64(rl.GetScreenWidth()), float64(rl.GetScreenHeight()))))

	rl.DrawFPS(0, 0)
	rl.EndDrawing()
}

func quit() {
	rl.CloseWindow()
}
