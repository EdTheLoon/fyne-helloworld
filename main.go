package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Initialise the application with a unique app ID
	a := app.NewWithID("scot.reids.hellofyne")

	// Initialise the window and set the title
	w := a.NewWindow("Hello!")

	// Header label
	header := widget.NewRichTextFromMarkdown("# Hello, World!\n\nvia Fyne!")

	// Separator
	sep := widget.NewSeparator()

	// Create a label widget
	hello := widget.NewLabel("This was created using Fyne! Try saying Hi!")

	// Create a button with text "Hi!" and a ConfirmIcon then define a function
	// that runs when the button is clicked!
	// We declare a variable which is a pointer so that we can self reference later.
	// This is unnecessary if we do not need the button's text to change when clicked.
	// See below commented code for the previous example without self referencing.
	count := 0
	var button *widget.Button
	button = widget.NewButtonWithIcon("Hi!", theme.ConfirmIcon(), func() {
		count++
		fmt.Printf("Button clicked %v time(s)!\n", count)
		var helloTxt string
		if count == 1 {
			helloTxt = "Hello :) (clicked once)"
		} else {
			helloTxt = fmt.Sprintf("Hello :) (clicked %v times)", count)
		}
		hello.SetText(helloTxt)
		buttonTxt := fmt.Sprintf("Hi! (%v)", count)
		button.SetText(buttonTxt)
	})
	button.IconPlacement = widget.ButtonIconTrailingText

	// Old button code without self referencing
	/*
		button := widget.NewButtonWithIcon("Hi!", theme.ConfirmIcon(), func() {
			count++
			fmt.Printf("Button clicked %v time(s)!\n", count)
			var helloTxt string
			if count == 1 {
				helloTxt = "Hello :) (clicked once)"
			} else {
				helloTxt = fmt.Sprintf("Hello :) (clicked %v times)", count)
			}
			hello.SetText(helloTxt)
		})
		button.IconPlacement = widget.ButtonIconTrailingText
	*/

	// Use a VBox to lay out the elements
	vbox := container.NewVBox(header, sep, hello, button)

	// Add the VBox and included elements to the window.
	w.SetContent(vbox)

	// Show and run the event handling for the window
	// TODO: Research whether this can run concurrently?
	w.ShowAndRun()
}
