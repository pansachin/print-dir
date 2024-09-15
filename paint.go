package main

import paint "github.com/fatih/color"

// setPaint sets stdout color.
func setPaint() {
	p = paint.New()
	switch color {
	case "red":
		p.Add(paint.FgRed)
	case "blue":
		p.Add(paint.FgBlue)
	case "green":
		p.Add(paint.FgGreen)
	case "cyan":
		p.Add(paint.FgCyan)
	case "yellow":
		p.Add(paint.FgYellow)
	}
}

func print(name string) {
	p.Print(name)
}

func println(name string) {
	p.Println(name)
}
