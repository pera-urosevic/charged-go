package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var Devices []DeviceModel = []DeviceModel{}

var Time time.Duration = 60 * time.Second

type DeviceModel struct {
	Name    string `json:"name"`
	Low     int    `json:"low"`
	High    int    `json:"high"`
	Command string `json:"command"`
}

type ConfigModel struct {
	Devices []DeviceModel `json:"devices"`
	Time    time.Duration `json:"time"`
}

func configPath() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\\charged.json", dirname), nil
}

func Load() error {
	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var Config = ConfigModel{}
	err = json.Unmarshal(data, &Config)
	if err != nil {
		return err
	}
	Devices = Config.Devices
	Time = Config.Time

	return nil
}
