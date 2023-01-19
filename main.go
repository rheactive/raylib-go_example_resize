package main

import "github.com/gen2brain/raylib-go/raylib"

func main() {
	var window_width float32 = 1600
	var window_height float32 = 900
	var scale_x float32 = 1
	var scale_y float32 = 1
	var margin_x float32 = 0.05 * window_width
	var margin_y float32 = 0.1 * window_height

	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(int32(window_width), int32(window_height), "Raylib example - window resize")
	
	rl.SetTargetFPS(60)

	var start rl.Vector2 = rl.NewVector2(margin_x, 1.5 * margin_y);
    var end rl.Vector2 = rl.NewVector2(window_width - margin_x, window_height - 0.5 * margin_y);

	for !rl.WindowShouldClose() {
		// Update
		if rl.IsWindowResized() || rl.IsWindowMaximized() {
			scale_x = float32(rl.GetScreenWidth())/window_width
			scale_y = float32(rl.GetScreenHeight())/window_height
			window_width = float32(rl.GetScreenWidth())
			window_height = float32(rl.GetScreenHeight())
			margin_x *= scale_x
			margin_y *= scale_y

			start.X *= scale_x
			start.Y *= scale_y
			end.X *= scale_x
			end.Y *= scale_y
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

		fsz := int32(0.02 * window_width)
		text_x := (int32(window_width) - rl.MeasureText(text, fsz))/2
		text_y := int32(0.05 * window_height)

		lwd := 0.004 * window_height
		inner_margin_x := 0.016 * window_width
		inner_margin_y := 0.016 * window_height
		off_x := 2 * margin_x - 2 * inner_margin_x
		off_y := 2 * margin_y - 2 * inner_margin_y

		rec := rl.NewRectangle(margin_x - inner_margin_x, 1.5 * margin_y - inner_margin_y, window_width - off_x, window_height - off_y)
		rl.DrawRectangleLinesEx(rec, 0.75 * lwd, rl.Gray)

		rl.DrawText(text, text_x, text_y, fsz, rl.Black);

        rl.DrawLineBezier(start, end, lwd, rl.Red);

        rl.EndDrawing();

	}

	rl.CloseWindow()
}