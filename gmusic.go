package main

import "fmt"
import gc "github.com/rthornton128/goncurses"
import "log"

func main() {
	fmt.Println("vim-go")

	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	// Turn off character echo, hide the cursor and disable input buffering
	gc.Echo(false)
	gc.CBreak(true)
	gc.Cursor(0)

	// Determine the center of the screen and offset those coordinates by
	// half of the window size we are about to create. These coordinates will
	// be used to move our window around the screen
	rows, cols := stdscr.MaxYX()
	height, width := 5, 10
	y, x := (rows-height)/2, (cols-width)/2

	// Create a new window centered on the screen and enable the use of the
	// keypad on it so the arrow keys are available
	var win *gc.Window
	win, err = gc.NewWindow(height, width, y, x)
	if err != nil {
		log.Fatal(err)
	}
	win.Keypad(true)

main:
	for {
		// Clear the section of screen where the box is currently located so
		// that it is blanked by calling Erase on the window and refreshing it
		// so that the chances are sent to the virtual screen but not actually
		// output to the terminal
		win.Erase()
		win.NoutRefresh()
		// Move the window to it's new location (if any) and redraw it
		win.MoveWindow(y, x)
		win.Box(gc.ACS_VLINE, gc.ACS_HLINE)
		win.NoutRefresh()
		// Update will flush only the characters which have changed between the
		// physical screen and the virtual screen, minimizing the number of
		// characters which must be sent
		gc.Update()

		// In order for the window to display correctly, we must call GetChar()
		// on it rather than stdscr
		switch win.GetChar() {
		case 'q':
			break main
		case gc.KEY_LEFT:
			if x > 0 {
				x--
			}
		case gc.KEY_RIGHT:
			if x < cols-width {
				x++
			}
		case gc.KEY_UP:
			if y > 1 {
				y--
			}
		case gc.KEY_DOWN:
			if y < rows-height {
				y++
			}
		}
	}
	win.Delete()
}
