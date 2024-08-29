package ui

import "github.com/lxn/walk"

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

	return nil
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
