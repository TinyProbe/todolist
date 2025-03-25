package main

import (
  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/app"
  "fyne.io/fyne/v2/container"
)

func main() {
  myApp := app.New()
  myWin := myApp.NewWindow("ToDo-List")
  myWin.Resize(fyne.NewSize(480.0, 360.0))

  tabs := container.NewAppTabs(
    DisplayTab(myWin),
    AddTab(),
  )
  tabs.SetTabLocation(container.TabLocationTop)

  myWin.SetContent(tabs)
  myWin.ShowAndRun()
}
