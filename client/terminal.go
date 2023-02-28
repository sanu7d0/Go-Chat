package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type loginInput struct {
	id string
}

// tview
var Pages = tview.NewPages()
var app = tview.NewApplication()
var menu = tview.NewFlex()
var chats = tview.NewFlex()

// user
var _loginInput = loginInput{}

func init() {
	// addMenu()
	// addChats()

	// if err := app.SetRoot(Pages, true).Run(); err != nil {
	// 	panic(err)
	// }
}

func addMenu() {
	feedback := tview.NewTextView()

	buttons := tview.NewForm().
		AddInputField("ID", "", 15, nil, func(text string) {
			_loginInput.id = text
		}).
		AddButton("Login", func() {
			_ = TryLogin(_loginInput.id)
			// if err != nil {
			// 	feedback.Clear()
			// 	feedback.SetTextColor(tcell.ColorDarkRed).SetText("Error login")
			// } else {
			// 	feedback.Clear()
			// }
		}).
		AddButton("Exit", func() {
			app.Stop()
		}).
		AddFormItem(feedback)
	buttons.SetBorder(true).SetTitle("Go CHAT")

	menu.AddItem(nil, 0, 1, false)
	menu.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(buttons, 0, 3, true).
		AddItem(nil, 0, 1, false), 0, 1, true)
	menu.AddItem(nil, 0, 1, false)

	Pages.AddPage("Menu", menu, true, true)
}

func addChats() {
	chats.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event
	})

	clients := tview.NewList()
	messages := tview.NewList()
	flex := tview.NewFlex().
		AddItem(clients, 0, 1, false).
		AddItem(messages, 0, 4, false)

	input := tview.NewInputField()

	chats.SetDirection(tview.FlexRow).
		AddItem(flex, 0, 5, false).
		AddItem(input, 0, 1, true)

	Pages.AddPage("Chats", chats, true, false)
}
