package main

// A simple program that counts down from 5 and then exits.

import (
	"fmt"
	"log"
	"time"

	"github.com/WhiCu/mangazeya/internal/tui/animator"
	tea "github.com/charmbracelet/bubbletea"
)

var frame = []string{
	"=",
	"==",
	"===",
	"====",
	"=====",
	"======",
	"=======",
	"========",
	"=========",
	"==========",
	"=========",
	"=======",
	"=====",
	"====",
	"===",
	"==",
}

func main() {
	// Log to a file. Useful in debugging since you can't really log to stdout.
	// Not required.
	if _, err := tea.LogToFile("tmp/log.txt", "debug"); err != nil {
		log.Fatal(err)
	}

	// Initialize our program
	p := tea.NewProgram(
		animator.New(
			animator.StringFrames(frame),
			time.Second/10,
		),
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

// A model can be more or less any type of data. It holds all the data for a
// program, so often it's a struct. For this simple example, however, all
// we'll need is a simple integer.
type model struct {
	count int
}

// Init optionally returns an initial command we should run. In this case we
// want to start the timer.
func (m model) Init() tea.Cmd {
	return tick
}

// Update is called when messages are received. The idea is that you inspect the
// message and send back an updated model accordingly. You can also return
// a command, which is a function that performs I/O and returns a message.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tickMsg:
		m.count--
		if m.count <= 0 {
			return m, tea.Quit
		}
		return m, tick
	}
	return m, nil
}

// View returns a string based on data in the model. That string which will be
// rendered to the terminal.
func (m model) View() string {
	return fmt.Sprintf("Hi. This program will exit in %d seconds.\n\nTo quit sooner press ctrl-c, or press ctrl-z to suspend...\n", m.count)
}

// Messages are events that we respond to in our Update function. This
// particular one indicates that the timer has ticked.
type tickMsg time.Time

func tick() tea.Msg {
	fmt.Println("tick")
	time.Sleep(time.Second)
	return tickMsg{}
}
