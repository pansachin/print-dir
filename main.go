package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"time"

	paint "github.com/fatih/color"
)

var (
	debug     bool
	icon      bool
	color     string
	recursive bool
	p         *paint.Color
)

func init() {
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.BoolVar(&icon, "icon", false, "show icon\ninstall nerd-font icon to see the icons\ndefault false")
	flag.StringVar(&color, "color", "", "show color\navailable color: red, blue, green, cyan, yellow\ndefault to system color")
	flag.BoolVar(&recursive, "recursive", false, "show recursivly")

	flag.Parse()
}

// print speaces.
func spaces(indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print("    ")
	}
}

// currentWalk print the content in current dir.
func currentdirWalk(dir string) error {
	dirName := filepath.Base(dir)
	icon := getFolderIcon(dirName)
	printIcon(icon)
	println(dirName)

	list, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, each := range list {
		var icon string
		print("├── ")

		if !each.IsDir() {
			ext := filepath.Ext(each.Name())
			if ext != "" {
				icon = getFileIcon(ext[1:])
			} else {
				icon = getFileIcon(ext)
			}
		} else {
			icon = getFolderIcon(each.Name())
		}
		printIcon(icon)

		println(each.Name())
	}

	return err
}

// recursiveWalk print the content recursivily.
func recursiveWalk(dir string, indent int) error {
	dirName := filepath.Base(dir)
	icon := getFolderIcon(dirName)
	printIcon(icon)
	println(dirName)

	f, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, each := range f {
		spaces(indent)
		print("├── ")
		if !each.IsDir() {
			var icon string
			ext := filepath.Ext(each.Name())
			if ext != "" {
				icon = getFileIcon(ext[1:])
			} else {
				icon = getFileIcon(ext)
			}
			printIcon(icon)
			println(each.Name())
		} else {
			if err := recursiveWalk(path.Join(dir, each.Name()), indent+1); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	// Intialize logger
	l := logging(debug)
	startTime := time.Now().UTC().Unix()
	l.Debug("init values", "debug", debug, "color", color, "recursive", recursive, "icon", icon)

	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Set StdOut color
	setPaint()

	// Default is non recursive walk
	switch recursive {
	case true:
		if err := recursiveWalk(wd, 0); err != nil {
			panic(err)
		}
	default:
		if err := currentdirWalk(wd); err != nil {
			panic(err)
		}
	}

	// Log total time taken
	l.Debug("total time taken", "duration", time.Now().UTC().Unix()-startTime)
}

func logging(debug bool) *slog.Logger {
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}
	l := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	return slog.New(l)
}
