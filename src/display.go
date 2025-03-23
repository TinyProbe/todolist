package main

import (
  "fyne.io/fyne/v2/container"
)

func DisplayTab() *container.TabItem {
  contents := container.NewVBox(
    container.NewHBox(),
  )

  return container.NewTabItem(
    "   Display   ",
    container.NewWithoutLayout(contents),
  )
}
