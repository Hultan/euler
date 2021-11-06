package euler_gui

import (
	"os"

	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/gtk-startup/pkg/tools"
)

const applicationTitle = "euler-gui"
const applicationVersion = "v0.0.2"
const applicationCopyRight = "©SoftTeam AB, 2021"

type MainForm struct {
	Window      *gtk.ApplicationWindow
	Helper      *GtkHelper
	AboutDialog *gtk.AboutDialog
}

// NewMainForm : Creates a new MainForm object
func NewMainForm() *MainForm {
	mainForm := new(MainForm)
	return mainForm
}

// OpenMainForm : Opens the MainForm window
func (m *MainForm) OpenMainForm(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new gtk helper
	builder, err := gtk.BuilderNewFromFile(GetResourcePath("../assets", "main.glade"))
	tools.ErrorCheckWithPanic(err, "Failed to create builder")
	helper := GtkHelperNew(builder)
	m.Helper = helper

	// Get the main window from the glade file
	window, err := helper.GetApplicationWindow("main_window")
	tools.ErrorCheckWithPanic(err, "Failed to find main_window")

	m.Window = window

	// Set up main window
	window.SetApplication(app)
	window.SetTitle("gtk-startup main window")

	// Hook up the destroy event
	_, err = window.Connect("destroy", window.Close)
	tools.ErrorCheckWithPanic(err, "Failed to connect the mainForm.destroy event")

	// Quit button
	button, err := helper.GetToolButton("toolbar_quit_button")
	tools.ErrorCheckWithPanic(err, "Failed to find toolbar_quit_button")
	_, err = button.Connect("clicked", window.Close)
	tools.ErrorCheckWithPanic(err, "Failed to connect the toolbar_quit_button.clicked event")

	listbox, err := helper.GetListBox("problem_listbox")
	tools.ErrorCheckWithPanic(err, "Failed to find problem_listbox")
	row := AddProblem("test")
	listbox.Add(row)

	// Status bar
	//statusBar, err := helper.GetStatusBar("main_window_status_bar")
	//tools.ErrorCheckWithPanic(err, "Failed to find main_window_status_bar")
	//statusBar.Push(statusBar.GetContextId("gtk-startup"), "gtk-startup : version 0.1.0")

	// Menu
	//m.setupMenu(window)

	// Show the main window
	window.ShowAll()
}

func AddProblem(name string) *gtk.ListBoxRow {
	row, err := gtk.ListBoxRowNew()
	if err != nil {
		panic(err)
	}
	hbox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	if err != nil {
		panic(err)
	}
	row.Add(hbox)
	label, err := gtk.LabelNew(name)
	if err != nil {
		panic(err)
	}
	hbox.PackStart(label, true, true, 0)
	return row
}


func (m *MainForm) setupMenu(window *gtk.ApplicationWindow) {
	menuQuit, err := m.Helper.GetMenuItem("menu_file_quit")
	tools.ErrorCheckWithPanic(err, "failed to find menu item menu_file_quit")
	_, err = menuQuit.Connect("activate", window.Close)
	tools.ErrorCheckWithoutPanic(err,"failed to connect menu_file_quit.activate signal")
}