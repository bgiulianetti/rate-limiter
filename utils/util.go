package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func LoadFile() []byte {
	fileData, err := os.ReadFile("./rules.json")
	if err != nil {
		fmt.Println("error reading from rules.json file::::", err)
	}
	return fileData
}

func FormatDuration(d time.Duration) string {
	if d.Hours() >= 1 {
		return fmt.Sprintf("%.0fh", d.Hours())
	} else if d.Minutes() >= 1 {
		return fmt.Sprintf("%.0fm", d.Minutes())
	} else {
		return fmt.Sprintf("%.0fs", d.Seconds())
	}
}

func SerializeObject(object interface{}) string {
	objectJson, _ := json.Marshal(object)
	return string(objectJson)
}
