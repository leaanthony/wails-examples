package main

import (
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/menu"

	"github.com/wailsapp/wails/v2"
)

// App application struct
type App struct {
	runtime *wails.Runtime
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (b *App) startup(runtime *wails.Runtime) {
	// Perform your setup here
	b.runtime = runtime
}

// shutdown is called at application termination
func (b *App) shutdown() {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (b *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

func (b *App) ApplicationMenu() *menu.Menu {

	handleRadio := func(cbdata *menu.CallbackData) {
		println("Radio Item selected:", cbdata.MenuItem.Label)
	}

	checkbox := &menu.MenuItem{
		Label:   "Checked",
		Type:    menu.CheckboxType,
		Checked: true,
		Click: func(data *menu.CallbackData) {
			println("checked =", data.MenuItem.Checked)
		},
	}

	radioMenu1 := menu.Radio("Radio 1", true, nil, handleRadio)
	radioMenu2 := menu.Radio("Radio 2", false, nil, handleRadio)
	radioMenu3 := menu.Radio("Radio 3", false, nil, handleRadio)

	secretMenu := menu.SubMenu("Secret Menu",
		menu.NewMenuFromItems(
			menu.Text("test", nil, nil),
			radioMenu1,
			radioMenu2,
			radioMenu3,
			checkbox,
		),
	)
	secretMenu.Hidden = true

	toggleSecretMenu := func(cbdata *menu.CallbackData) {
		secretMenu.Hidden = !secretMenu.Hidden
		if secretMenu.Hidden {
			cbdata.MenuItem.Label = "Show secret menu...😉"
		} else {
			cbdata.MenuItem.Label = "Hide secret menu! 😲"
		}
		b.runtime.Menu.UpdateApplicationMenu()
	}

	return menu.NewMenuFromItems(
		menu.SubMenu("File", menu.NewMenuFromItems(
			menu.Text("Open", nil, func(_ *menu.CallbackData) {
				println("Open CALLback called!")
			}),
			menu.Separator(),
			menu.SubMenu("Submenu",
				menu.NewMenuFromItems(
					&menu.MenuItem{
						Label:    "Disabled",
						Type:     menu.TextType,
						Disabled: true,
					},
					&menu.MenuItem{
						Label:  "Hidden",
						Hidden: true,
						Type:   menu.TextType,
					},
					checkbox,
					checkbox,
					checkbox,
					radioMenu1,
					radioMenu2,
					radioMenu3,
					menu.Text("Text Menu Item", nil, nil),
				),
			),
			menu.Text("Quit", nil, func(_ *menu.CallbackData) {
				println("QUIT CALLBACK CALLED!")
			}),
		)),
		menu.SubMenu("Tests", menu.NewMenuFromItems(
			menu.Text("Show secret menu...😉", nil, toggleSecretMenu),
			menu.Separator(),
			secretMenu,
		),
		),
	)

}
