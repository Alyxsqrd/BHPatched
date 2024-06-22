package main

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/sqweek/dialog"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			stack := strings.Split(strings.TrimRight(string(debug.Stack()), "\n"), "\n")
			dialog.Message("%s\n\n%s", err, strings.TrimLeft(stack[len(stack)-1], "\t")).Title("Brick Hill Patched").Error()
			panic(err)
		}
	}()

	cache := filepath.Join(os.Getenv("APPDATA"), "Godot", "app_userdata", "Brick Hill", "assets")
	err := os.MkdirAll(cache, 0660)

	if err != nil {
		panic(err)
	}

	assets := filepath.Join(os.Args[0], "..", "assets")
	files, err := os.ReadDir(assets)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		name := file.Name()
		dest, err := os.OpenFile(filepath.Join(cache, strings.TrimSuffix(name, filepath.Ext(name))), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0660)

		if errors.Is(err, os.ErrExist) {
			continue
		} else if err != nil {
			panic(err)
		}

		src, err := os.OpenFile(filepath.Join(assets, name), os.O_RDONLY, 0)

		if err != nil {
			panic(err)
		}

		_, err = io.Copy(dest, src)

		if err != nil {
			panic(err)
		}
	}

	err = exec.Command(filepath.Join(os.Args[0], "..", "workshop", "brick-hill.exe"), "res://scenes/workshop.tscn").Start()

	if err != nil {
		panic(err)
	}

	// Remove this credit... and I will remove your balls (respectfully)
	dialog.Message("Workshop has been successfully patched!\nhttps://github.com/Alyxsqrd/BHPatched").Title("Brick Hill Patched").Info()
}
