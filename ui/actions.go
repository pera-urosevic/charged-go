package ui

import (
	"charged/config"
	"charged/state"
	"fmt"

	"github.com/lxn/walk"
)

var deviceActions = map[string]*walk.Action{}

func addQuit(ni *walk.NotifyIcon) error {
	quitAction := walk.NewAction()
	if err := quitAction.SetText("Exit"); err != nil {
		return err
	}
	quitAction.SetImage(iconExit)
	quitAction.Triggered().Attach(func() { walk.App().Exit(0) })
	if err := ni.ContextMenu().Actions().Add(quitAction); err != nil {
		return err
	}
	return nil
}

func addDevices(ni *walk.NotifyIcon) error {
	for _, device := range config.Devices {
		deviceAction := walk.NewAction()
		level := state.Levels[device.Name]
		label := fmt.Sprintf("%s - %d%%", device.Name, level)
		if err := deviceAction.SetText(label); err != nil {
			return err
		}
		deviceAction.SetImage(iconDeviceGood)
		if err := ni.ContextMenu().Actions().Add(deviceAction); err != nil {
			return err
		}
		deviceActions[device.Name] = deviceAction
	}
	return nil
}

func updateDevices() bool {
	hasAlerts := false
	for _, device := range config.Devices {
		level := state.Levels[device.Name]
		deviceActions[device.Name].SetText(fmt.Sprintf("%s - %d%%", device.Name, level))
		alert := state.Alerts[device.Name]
		var deviceActionIcon walk.Image
		if alert {
			deviceActionIcon = iconDeviceAlert
		} else {
			deviceActionIcon = iconDeviceGood
		}
		deviceActions[device.Name].SetImage(deviceActionIcon)
		hasAlerts = hasAlerts || alert
	}
	return hasAlerts
}
