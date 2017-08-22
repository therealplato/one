// one
// choose one line from stdin
package main

import (
	"fmt"
	"os"

	ui "github.com/gizak/termui"
)

var inputstring = `foo@localhost
bar@as.internal.df.com
somehost
`

type one struct {
	ls      *ui.List
	input   string
	targets []string
	cursor  int
	success bool
}

func (o *one) move(i int) {
	l := len(o.targets)
	tmp := o.cursor + i
	if tmp < 0 {
		tmp += l
	}
	if tmp > l {
		tmp -= l
	}
	ui.Render(o.ls)
}
func (o *one) exit() {
	ui.StopLoop()
	ui.Close()
	if o.success {
		if len(o.targets) > o.cursor {
			fmt.Println(o.targets[o.cursor])
		}
		os.Exit(0)
	}
	os.Exit(1)
}

func main() {
	// ui.NewMarkdownTxBuilder()
	err := ui.Init()
	if err != nil {
		panic(err)
	}

	o := &one{input: inputstring}

	strs := []string{
		"[0] github.com/gizak/termui",
		"[3] [color output](fg-white,bg-green)",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] nsf/termbox-go"}
	o.targets = strs
	o.cursor = 2

	ls := ui.NewList()
	ls.Items = strs
	ls.ItemFgColor = ui.ColorYellow
	ls.BorderLabel = "List"
	ls.Height = 7
	ls.Width = 25
	ls.Y = 0

	ui.Render(ls)
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		ui.Close()
	})

	// ui.Render(p)
	ui.Handle("/sys/kbd", func(e ui.Event) {
		ek, ok := e.Data.(ui.EvtKbd)
		if !ok {
			return
		}
		ls.Items = []string{ek.KeyStr}
		ui.Render(ls)
		if ek.KeyStr == "<down>" {
			o.move(-1)
		}
		if ek.KeyStr == "<up>" {
			o.move(+1)
		}
		if ek.KeyStr == "enter" {
			o.success = true
			o.exit()
		}
	})
	ui.Handle("/sys/kbd/q", func(e ui.Event) {
		o.success = true
		o.exit()
	})

	// event handler...
	ui.Loop()
}
