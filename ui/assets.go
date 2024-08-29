package ui

import "github.com/lxn/walk"

var iconStatusGood *walk.Icon
var iconStatusAlert *walk.Icon
var iconDeviceGood *walk.Icon
var iconDeviceAlert *walk.Icon
var iconExit *walk.Icon

func loadAssets() error {
	var err error

	iconStatusGood, err = walk.Resources.Icon("./assets/status-good.ico")
	if err != nil {
		return err
	}

	iconStatusAlert, err = walk.Resources.Icon("./assets/status-alert.ico")
	if err != nil {
		return err
	}

	iconDeviceGood, err = walk.Resources.Icon("./assets/device-good.ico")
	if err != nil {
		return err
	}

	iconDeviceAlert, err = walk.Resources.Icon("./assets/device-alert.ico")
	if err != nil {
		return err
	}

	iconExit, err = walk.Resources.Icon("./assets/exit.ico")
	if err != nil {
		return err
	}

	return nil
}
