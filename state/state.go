package state

import (
	"charged/cmd"
	"charged/config"
	"fmt"
	"strconv"
)

var Levels = map[string]int{}

var Alerts = map[string]bool{}

func Init() {
	Update()
}

func Update() {
	for _, device := range config.Devices {
		level := getLevel(device)
		Levels[device.Name] = level
		if level < device.Low {
			Alerts[device.Name] = true
		} else {
			Alerts[device.Name] = false
		}

	}
	fmt.Println("State", Levels)
}

func getLevel(device config.DeviceModel) int {
	result, err := cmd.Run(device.Command)
	if err != nil {
		fmt.Println("Error executing cmd", err)
		return 0
	}
	level, err := strconv.Atoi(result)
	if err != nil {
		fmt.Println("Error converting result", err)
		return 0
	}
	return level
}
