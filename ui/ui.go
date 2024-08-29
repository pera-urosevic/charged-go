package ui

import (
	"charged/config"
	"charged/state"
	"time"

	"github.com/lxn/walk"
)

var mainWindow *walk.MainWindow

func Init() error {
	var err error

	loadAssets()

	mainWindow, err = createWindow()
	if err != nil {
		return err
	}

	err = createNotifyIcon(mainWindow)
	if err != nil {
		return err
	}
	defer notifyIcon.Dispose()

	err = addDevices(notifyIcon)
	if err != nil {
		return err
	}

	err = addQuit(notifyIcon)
	if err != nil {
		return err
	}

	update()

	err = notifyIcon.SetVisible(true)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(config.Time * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				update()
			}
		}
	}()

	mainWindow.Run()
	return nil
}

func update() {
	state.Update()
	hasAlerts := updateDevices()
	updateNotifyIcon(hasAlerts)
}
