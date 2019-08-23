package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/greymd/ojichat/generator"
	"github.com/pkg/errors"
)

const (
	appName = "OjiDesktop"
)

var (
	debug = flag.Bool("d", false, "enables the debug mode")
	w     *astilectron.Window
)

func main() {
	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap
	astilog.Debugf("Running app")
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            appName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
		},
		Debug: *debug,
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astilectron.PtrStr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{Role: astilectron.MenuItemRoleClose},
			},
		}},
		OnWait: func(_ *astilectron.Astilectron, ws []*astilectron.Window, _ *astilectron.Menu, _ *astilectron.Tray, _ *astilectron.Menu) error {
			w = ws[0]
			go func() {
				time.Sleep(5 * time.Second)
				if err := bootstrap.SendMessage(w, "check.out.menu", "Don't forget to check out the menu!"); err != nil {
					astilog.Error(errors.Wrap(err, "sending check.out.menu event failed"))
				}
			}()
			return nil
		},
		RestoreAssets: RestoreAssets,
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astilectron.PtrStr("#333"),
				Center:          astilectron.PtrBool(true),
				Height:          astilectron.PtrInt(470),
				Width:           astilectron.PtrInt(650),
			},
		}},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "ojimessage":
		if payload, err = ojiMessage(); err != nil {
			payload = err.Error()
		}
		return
	}
	return
}

func ojiMessage() (string, error) {
	en := rand.Intn(5)
	if en < 1 {
		en = 1
	}
	pn := rand.Intn(3)

	message, err := generator.Start(generator.Config{
		TargetName:        "astilectron",
		EmojiNum:          en,
		PunctiuationLebel: pn,
	})
	if err != nil {
		return "", err
	}

	return message, nil
}
