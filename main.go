// one
// choose one line from stdin
package main

import (
	"fmt"
	"os"

	ui "github.com/gizak/termui"
)

func main() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}

	strs := []string{
		"[0] github.com/gizak/termui",
		"[3] color output",
		"[4] output.go",
		"[5] random_out.go",
		"[6] dashboard.go",
		"[7] nsf/termbox-go",
	}

	o := &one{
		targets: strs,
		cursor:  2,
		ls:      ui.NewList(),
	}

	o.ls.Items = strs
	o.ls.ItemFgColor = ui.ColorWhite
	o.ls.BorderLabel = "Select A Target"
	o.ls.Height = 15
	o.ls.Y = 0

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, o.ls),
		),
	)

	o.Render()

	down := func(e ui.Event) {
		o.move(1)
	}
	up := func(e ui.Event) {
		o.move(-1)
	}
	quit := func(e ui.Event) {
		o.success = true
		o.exit()
	}

	ui.Handle("/sys/kbd/j", down)
	ui.Handle("/sys/kbd/<down>", down)
	ui.Handle("/sys/kbd/<up>", up)
	ui.Handle("/sys/kbd/k", up)
	ui.Handle("/sys/kbd/q", quit)
	ui.Handle("/sys/kbd/<enter>", quit)
	ui.Loop()
}

type one struct {
	ls       *ui.List
	targets  []string
	rendered []string
	cursor   int
	success  bool
}

func (o *one) move(i int) {
	if o == nil {
		return
	}
	l := len(o.targets)
	tmp := o.cursor + i
	if tmp < 0 {
		tmp += l
	}
	if tmp > l {
		tmp -= l
	}
	o.cursor = tmp
	o.Render()
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

func (o *one) Render() {
	o.rendered = make([]string, 0, len(o.targets))
	for i, s := range o.targets {
		if i == o.cursor {
			o.rendered = append(o.rendered, fmt.Sprintf("[%s](fg-red)", s))
			continue
		}
		o.rendered = append(o.rendered, s)
	}
	o.ls.Items = o.rendered
	ui.Body.Align()
	ui.Render(ui.Body)
}
