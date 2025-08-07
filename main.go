// package main

// import "github.com/WhiCu/mangazeya/cmd"

//	func main() {
//		cmd.Execute()
//	}
package main

import (
	"log"

	"github.com/WhiCu/mangazeya/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// Log to a file. Useful in debugging since you can't really log to stdout.
	// Not required.
	if _, err := tea.LogToFile("tmp/log.txt", "debug"); err != nil {
		log.Println(err)
	}

	// Initialize our program
	p := tui.NewProgram()
	// p := tea.NewProgram(
	// 	animator.New(
	// 		animator.StringFrames(frame),
	// 		time.Second/10,
	// 	),
	// 	tea.WithAltScreen(),
	// )
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
