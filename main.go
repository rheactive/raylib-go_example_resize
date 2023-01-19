package main

import "github.com/gen2brain/raylib-go/raylib"

func main() {
	var windowWidth float32 = 1600
	var windowHeight float32 = 900
	var scaleX float32 = 1
	var scaleY float32 = 1
	var marginX float32 = 0.05 * windowWidth
	var marginY float32 = 0.1 * windowHeight

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(int32(windowWidth), int32(windowHeight), "Raylib example - window resize")
	
	rl.SetTargetFPS(60)

	var start rl.Vector2 = rl.NewVector2(marginX, 1.5 * marginY);
    var end rl.Vector2 = rl.NewVector2(windowWidth - marginX, windowHeight - 0.5 * marginY);

	for !rl.WindowShouldClose() {
		// Update
		if rl.IsWindowResized() || rl.IsWindowMaximized() {
			scaleX = float32(rl.GetScreenWidth())/windowWidth
			scaleY = float32(rl.GetScreenHeight())/windowHeight
			windowWidth = float32(rl.GetScreenWidth())
			windowHeight = float32(rl.GetScreenHeight())
			marginX *= scaleX
			marginY *= scaleY

			start.X *= scaleX
			start.Y *= scaleY
			end.X *= scaleX
			end.Y *= scaleY
		}
        //----------------------------------------------------------------------------------
        if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			start = rl.GetMousePosition()
		}
        if rl.IsMouseButtonDown(rl.MouseRightButton) {
			end = rl.GetMousePosition()
		}
		
        //----------------------------------------------------------------------------------
        // Draw
        //----------------------------------------------------------------------------------
        rl.BeginDrawing();

		rl.ClearBackground(rl.RayWhite);

		text := "USE MOUSE LEFT-RIGHT CLICK to DEFINE LINE START and END POINTS"

		fsz := int32(0.02 * windowWidth)
		text_x := (int32(windowWidth) - rl.MeasureText(text, fsz))/2
		text_y := int32(0.05 * windowHeight)

		lwd := 0.004 * windowHeight
		innerMarginX := 0.016 * windowWidth
		innerMarginY := 0.016 * windowHeight
		off_x := 2 * marginX - 2 * innerMarginX
		off_y := 2 * marginY - 2 * innerMarginY

		rec := rl.NewRectangle(start.X - innerMarginX, start.Y - innerMarginY, windowWidth - off_x, windowHeight - off_y)
		rl.DrawRectangleLinesEx(rec, 0.75 * lwd, rl.Gray)

		rl.DrawText(text, text_x, text_y, fsz, rl.Black);

        rl.DrawLineBezier(start, end, lwd, rl.Red);

        rl.EndDrawing();

	}

	rl.CloseWindow()
}