package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/devices"
	"github.com/ysmood/kit"
)

func main() {
	browser := rod.New().Connect()
	defer browser.Close()

	page := browser.Page("")

	// emulate iPhone 7 landscape
	err := page.EmulateE(devices.IPhone6or7or8, true)
	kit.E(err)

	page.Navigate("https://www.whatsmyua.info/")
	page.Screenshot("screenshot1.png")

	// reset
	page.Emulate("")

	page.Viewport(1920, 2000, 1, false)
	page.Navigate("https://www.whatsmyua.info/?a")
	page.Screenshot("screenshot2.png")
}
