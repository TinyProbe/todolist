package main

import (
  "fmt"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/dialog"
  "fyne.io/fyne/v2/widget"
)

func DisplayTab(mainWin fyne.Window) *container.TabItem {
  BringScheduleData("todolist.json")
  selected := -1
  list := widget.NewList(
    func() int { return len(scheduleData) },
    func() fyne.CanvasObject { return widget.NewLabel("Schedules") },
    func(id widget.ListItemID, cell fyne.CanvasObject) {
      label := cell.(*widget.Label)
      label.SetText(scheduleData[id]["year"] + "-" +
          scheduleData[id]["month"] + "-" +
          scheduleData[id]["day"] + "  " +
          scheduleData[id]["title"])
      if id == selected {
        label.Wrapping = fyne.TextWrapBreak
        label.TextStyle.Bold = true
      } else {
        label.TextStyle.Bold = false
      }
      label.Refresh()
    },
  )
  list.OnSelected = func(id widget.ListItemID) {
    selected = id
    list.Refresh()
    fmt.Println("selected row:", scheduleData[id])
  }
  list.Resize(fyne.NewSize(320.0, 240.0))
  list.Move(fyne.NewPos(0.0, 8.0))

  buttonRemove := widget.NewButton("Remove", func() {
    if selected >= 0 {
      SendNotification("Remove Schedule", scheduleData[selected]["title"])
      result := []map[string]string {}
      for i := 0; i < len(scheduleData); i++ {
        if i == selected { continue }
        result = append(result, scheduleData[i])
      }
      scheduleData = result
      list.Refresh()
    }
    if len(scheduleData) == 0 { selected = 0 }
  })
  buttonRemove.Resize(fyne.NewSize(80.0, 40.0))
  buttonRemove.Move(fyne.NewPos(0.0, 0.0))
  buttonInspect := widget.NewButton("Inspect", func() {
    if selected >= 0 {
      dialog.ShowInformation(
          "Description", scheduleData[selected]["description"], mainWin)
    }
  })
  buttonInspect.Resize(fyne.NewSize(80.0, 40.0))
  buttonInspect.Move(fyne.NewPos(16.0, 0.0))

  contents := container.NewVBox(
    container.NewHBox(
      container.NewWithoutLayout(buttonRemove),
      container.NewWithoutLayout(buttonInspect),
    ),
    container.NewWithoutLayout(list),
  )
  contents.Move(fyne.NewPos(8.0, 14.0))

  return container.NewTabItem(
    "   Display   ",
    container.NewWithoutLayout(contents),
  )
}
