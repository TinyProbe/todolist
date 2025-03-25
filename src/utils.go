package main

import (
  "encoding/json"
  "os"
  "sort"
  "time"
  "fmt"
  "os/exec"
  "runtime"
)

var scheduleData []map[string]string
var getYear int = time.Now().Year()

func SendNotification(title, message string) {
  switch runtime.GOOS {
  case "linux":
    cmd := exec.Command("notify-send", title, message)
    cmd.Run()
  case "darwin":
    cmd := exec.Command("osascript", "-e", fmt.Sprintf("%s: %s", title, message))
    cmd.Run()
  default:
    fmt.Println("This OS doesn't support notifications.")
  }
}

func BringScheduleData(filename string) bool {
  jsonData, err := os.ReadFile(filename)
  if err != nil {
    fmt.Println("read file error:", err)
    return false
  }
  err = json.Unmarshal(jsonData, &scheduleData)
  if err != nil {
    fmt.Println("json parsing error:", err)
    return false
  }
  return true
}

func SortScheduleData() {
  sort.Slice(scheduleData, func(i, j int) bool {
    if scheduleData[i]["year"] == scheduleData[j]["year"] {
      if scheduleData[i]["month"] == scheduleData[j]["month"] {
        if scheduleData[i]["day"] == scheduleData[j]["day"] {
          return scheduleData[i]["title"] < scheduleData[j]["title"]
        } else { return scheduleData[i]["day"] < scheduleData[j]["day"] }
      } else { return scheduleData[i]["month"] < scheduleData[j]["month"] }
    } else { return scheduleData[i]["year"] < scheduleData[j]["year"] }
  })
}
