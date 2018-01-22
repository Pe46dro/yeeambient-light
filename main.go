
package main

import (
	"image"
	"fmt"
	"time"
	"strconv"

	"github.com/nunows/goyeelight"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/vova616/screenshot"
	"github.com/jakekausler/prominentcolor"
	"github.com/getlantern/systray"

	"./icons"
	"os"
)

func main() {

	lamp := goyeelight.New("192.168.1.X", "55443")

	var (
		r       string
		prevHex string
	)

	go func() {
		for {
			time.Sleep(500 * time.Millisecond)

			screen, _ := screenshot.CaptureScreen()
			img := image.Image(screen)

			cols, err := prominentcolor.KmeansWithArgs(prominentcolor.ArgumentNoCropping, img)

			if err != nil {
				fmt.Println("Error")
				continue
			}

			col := cols[0].AsString()

			if prevHex == "#"+col {
				continue
			} else {
				prevHex = "#" + col
			}

			fmt.Println(prevHex)

			c, _ := colorful.Hex(prevHex)

			h, s, v := c.Hsv()

			fmt.Println(h)
			fmt.Println(s)
			fmt.Println(v * 100)

			r = lamp.SetBright(strconv.FormatFloat(v*100, 'f', 1, 64), "smooth", "1000")
			fmt.Println(r)

			r = lamp.SetHSV(strconv.FormatFloat(h, 'f', 1, 64), strconv.FormatFloat(s*100, 'f', 1, 64), "smooth", "1000")
			fmt.Println(r)

		}
	}()

	systray.Run(onReady, onExit)

}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Yee-Ambient Light")
	systray.SetTooltip("Yee-Ambient Light")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		os.Exit(0)
	}()
}

func onExit() {
	// clean up here
}