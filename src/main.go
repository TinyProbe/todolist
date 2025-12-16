package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
)

func main() {
    myApp := app.New()
    myWin := myApp.NewWindow("ToDoList")
    myWin.Resize(fyne.NewSize(500.0, 400.0))
    myWin.SetFixedSize(true)
    myWin.CenterOnScreen()
    myWin.SetContent(
        container.NewAppTabs(
            DisplayTab(myWin),
            AddTab(),
        ),
    )
    myWin.ShowAndRun()
}
