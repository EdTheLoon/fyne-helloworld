package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Initialise the application
	a := app.New()

	// Initialise the window and set the title
	w := a.NewWindow("Hello")

	// Create a label widget
	hello := widget.NewLabel("Hello, Fyne!")

	// Add the label to the window
	w.SetContent(container.NewVBox( // Create a vertical box container
		hello, // First element is the label
		widget.NewButton("Hi!", // Second element is a button. Set button text here
			func() { // Button's event handler. Can also call a function written elsewhere
				hello.SetText("Welcome :)")
			}),
	)) // End of container.NewVBox and w.SetContent

	// Show and run the event handling for the window
	// TODO: Research whether this can run concurrently?
	w.ShowAndRun()
}
