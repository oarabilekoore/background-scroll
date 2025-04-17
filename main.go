package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	var screenWidth int32 = 800
	var screenHeight int32 = 450

	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - background scrolling")
	defer rl.CloseWindow()

	// Declare Variables & Load Textures
	var background rl.Texture2D = rl.LoadTexture("resources/cyberpunk_street_background.png")
	var midground rl.Texture2D = rl.LoadTexture("resources/cyberpunk_street_midground.png")
	var foreground rl.Texture2D = rl.LoadTexture("resources/cyberpunk_street_foreground.png")

	var scrollingBack float32 = 0.0
	var scrollingMid float32 = 0.0
	var scrollingFore float32 = 0.0

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		scrollingBack = 0.1
		scrollingMid = 0.5
		scrollingFore = 1.0

		//Texture is scaled twice when scrolling
		if scrollingBack <= float32(-background.Width*2) {
			scrollingBack = 0
		}
		if scrollingMid <= float32(-midground.Width*2) {
			scrollingMid = 0
		}
		if scrollingFore <= float32(-foreground.Width*2) {
			scrollingFore = 0
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.GetColor(0x052c46ff))

		rl.DrawTextureEx(background, rl.Vector2{scrollingBack, 20}, 0.0, 2.0, rl.White)
		rl.DrawTextureEx(background, rl.Vector2{background.Width*2 + scrollingBack, 20}, 0.0, 2.0, rl.White)
		rl.EndDrawing()
	}
}
