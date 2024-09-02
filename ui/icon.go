package ui

import (
	"charged/config"
	"charged/state"
	"fmt"

	"github.com/lxn/walk"
)

var notifyIcon *walk.NotifyIcon

func createNotifyIcon(mw *walk.MainWindow) error {
	var err error

	notifyIcon, err = walk.NewNotifyIcon(mw)
	if err != nil {
		return err
	}

	err = notifyIcon.SetIcon(iconStatusGood)
	if err != nil {
		return err
	}

	err = notifyIcon.SetToolTip("Charged")
	if err != nil {
		return err
	}

	onLeftClick()

	return nil
}

func onLeftClick() {
	notifyIcon.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}
		alerts := false
		statuses := ""
		for _, device := range config.Devices {
			level := state.Levels[device.Name]
			alerts = alerts || state.Alerts[device.Name]
			statuses += fmt.Sprintf("%s - %d%%\n", device.Name, level)
		}
		if alerts {
			notifyIcon.ShowCustom("", statuses, iconStatusAlert)
		} else {
			notifyIcon.ShowCustom("", statuses, iconStatusGood)
		}
	})
}

func updateNotifyIcon(hasAlerts bool) {
	if hasAlerts {
		setNotifyIconAlert()
	} else {
		setNotifyIconGood()
	}
}

func setNotifyIconGood() {
	notifyIcon.SetIcon(iconStatusGood)
	notifyIcon.SetToolTip("Charged")
}

func setNotifyIconAlert() {
	notifyIcon.SetIcon(iconStatusAlert)
	notifyIcon.SetToolTip("Warning!")
}
