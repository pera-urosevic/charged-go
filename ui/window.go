package ui

import "github.com/lxn/walk"

func createWindow() (*walk.MainWindow, error) {
	mw, err := walk.NewMainWindow()
	return mw, err
}
