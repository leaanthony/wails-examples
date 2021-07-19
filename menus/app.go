package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime/dialog"
	"github.com/wailsapp/wails/v2/pkg/runtime/events"
	"github.com/wailsapp/wails/v2/pkg/runtime/log"
)

// App application struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	println("startup called with ctx!")
	a.ctx = ctx
}

// shutdown is called at application termination
func (a *App) shutdown() {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

// ApplicationMenu creates a menu that we will use for the
// application
func (a *App) ApplicationMenu() *menu.Menu {

	// We can use the same callback for multiple menu items
	handleRadio := func(cbdata *menu.CallbackData) {
		// The menu item that was selected can be accessed through the
		// callback data
		events.Emit(a.ctx, "updatecolour", cbdata.MenuItem.Label)
	}

	// We can define menu items and reuse them.
	// Their state will stay in sync.
	checkbox := &menu.MenuItem{
		Label:   "Checked",
		Type:    menu.CheckboxType,
		Checked: true,
		Click: func(cbdata *menu.CallbackData) {
			log.Debug(a.ctx, fmt.Sprintf("checked = %v", cbdata.MenuItem.Checked))
			events.Emit(a.ctx, "checkboxstatus", cbdata.MenuItem.Label)
		},
	}

	// Radio groups are automatically made by adjacent radio menu items
	radioMenu1 := menu.Radio("White", true, nil, handleRadio)
	radioMenu2 := menu.Radio("Green", false, nil, handleRadio)
	radioMenu3 := menu.Radio("Red", false, nil, handleRadio)

	secretMenu := menu.SubMenu("Secret Menu",
		menu.NewMenuFromItems(
			menu.Text("test", nil, nil),
			radioMenu1,
			radioMenu2,
			radioMenu3,
			checkbox,
		),
	)

	// Menu items can be hidden
	secretMenu.Hidden = true

	// Menus can be updated by callbacks. Updates will get reflected after
	// calling `runtime.Menu.UpdateApplicationMenu()`
	toggleSecretMenu := func(cbdata *menu.CallbackData) {
		secretMenu.Hidden = !secretMenu.Hidden
		if secretMenu.Hidden {
			cbdata.MenuItem.Label = "Show secret menu...😉"
		} else {
			cbdata.MenuItem.Label = "Hide secret menu! 😲"
		}
		//a.runtime.Menu.UpdateApplicationMenu()
	}

	return menu.NewMenuFromItems(
		menu.SubMenu("File", menu.NewMenuFromItems(
			menu.Text("&Open", keys.CmdOrCtrl("o"), func(cbdata *menu.CallbackData) {
				filename, err := dialog.OpenFile(a.ctx, dialog.OpenDialogOptions{
					DefaultDirectory:           "",
					DefaultFilename:            "",
					Title:                      "Selct yer fil",
					Filters:                    nil,
					AllowFiles:                 false,
					AllowDirectories:           false,
					ShowHiddenFiles:            false,
					CanCreateDirectories:       false,
					ResolvesAliases:            false,
					TreatPackagesAsDirectories: false,
				})
				_, err = dialog.Message(a.ctx, dialog.MessageDialogOptions{
					Type:          "info",
					Title:         "Hello",
					Message:       filename,
					Buttons:       nil,
					DefaultButton: "",
					CancelButton:  "",
					Icon:          "",
				})
				if err != nil {
					return
				}
				if err != nil {
					return
				}
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
			menu.Text("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
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
