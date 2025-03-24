package main

import (
  "encoding/json"
  "strconv"
  "time"
  "fmt"
  "os"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/widget"
)

func AddTab() *container.TabItem {
  singleTextField := widget.NewEntry()
  singleTextField.Resize(fyne.NewSize(260.0, 36.0))
  multiTextField := widget.NewMultiLineEntry()
  multiTextField.Resize(fyne.NewSize(320.0, 120.0))

  years := generateYears()
  yearSelect := widget.NewSelect(years, nil)
  yearSelect.SetSelected(strconv.Itoa(getYear))
  months := generateMonths()
  monthSelect := widget.NewSelect(months, nil)
  monthSelect.SetSelected("1")
  days := generateDays(time.Now().Year(), 1)
  daySelect := widget.NewSelect(days, nil)
  daySelect.SetSelected("1")

  yearSelect.OnChanged = func(selectedYear string) {
    year, _ := strconv.Atoi(selectedYear)
    month, _ := strconv.Atoi(monthSelect.Selected)
    daySelect.Options = generateDays(year, month)
    daySelect.SetSelectedIndex(0)
  }
  monthSelect.OnChanged = func(selectedMonth string) {
    year, _ := strconv.Atoi(yearSelect.Selected)
    month, _ := strconv.Atoi(selectedMonth)
    daySelect.Options = generateDays(year, month)
    daySelect.SetSelectedIndex(0)
  }
  yearSelect.Resize(fyne.NewSize(80.0, 36.0))
  monthSelect.Resize(fyne.NewSize(60.0, 36.0))
  monthSelect.Move(fyne.NewPos(-40.0, 0.0))
  daySelect.Resize(fyne.NewSize(60.0, 36.0))
  daySelect.Move(fyne.NewPos(-100.0, 0.0))

  resetForms := func() {
    singleTextField.SetText("")
    multiTextField.SetText("")
    yearSelect.SetSelected(strconv.Itoa(getYear))
    monthSelect.SetSelected("1")
    daySelect.SetSelected("1")
  }

  buttonReset := widget.NewButton("Reset", resetForms)
  buttonReset.Resize(fyne.NewSize(80.0, 40.0))
  buttonReset.Move(fyne.NewPos(0.0, 50.0))
  buttonAdd := widget.NewButton("Add", func() {
    BringScheduleData("todolist.json")
    scheduleData = append(scheduleData, map[string]string {
      "title" : singleTextField.Text,
      "description" : multiTextField.Text,
      "year" : yearSelect.Selected,
      "month" : monthSelect.Selected,
      "day" : daySelect.Selected,
    })
    jsonData, err := json.MarshalIndent(scheduleData, "", "  ")
    if err != nil {
      fmt.Println("json parsing error:", err)
      return
    }
    err = os.WriteFile("todolist.json", jsonData, 0644)
    if err != nil {
      fmt.Println("write file error:", err)
      return
    }
    SendNotification("Add Schedule", singleTextField.Text)
    resetForms()
  })
  buttonAdd.Resize(fyne.NewSize(80.0, 40.0))
  buttonAdd.Move(fyne.NewPos(180.0, 50.0))

  contents := container.NewVBox(
    container.NewHBox(
      widget.NewLabel("Title:"),
      container.NewWithoutLayout(singleTextField),
    ),
    container.NewHBox(
      widget.NewLabel("Date:"),
      container.NewWithoutLayout(yearSelect),
      container.NewWithoutLayout(monthSelect),
      container.NewWithoutLayout(daySelect),
    ),
    container.NewVBox(
      widget.NewLabel("Description:"),
      container.NewWithoutLayout(multiTextField),
    ),
    container.NewHBox(
      container.NewWithoutLayout(buttonReset),
      container.NewWithoutLayout(buttonAdd),
    ),
  )
  contents.Move(fyne.NewPos(8.0, 14.0))

  return container.NewTabItem(
    "      Add      ",
    container.NewWithoutLayout(contents),
  )
}

func generateYears() []string {
  currentYear := getYear
  years := []string {}
  for i := currentYear; i <= currentYear + 100; i++ {
    years = append(years, strconv.Itoa(i))
  }
  return years
}

func generateMonths() []string {
  months := []string {}
  for i := 1; i <= 12; i++ {
    months = append(months, strconv.Itoa(i))
  }
  return months
}

func generateDays(year, month int) []string {
  if month == 0 { return []string {} }

  daysInMonth := 31
  switch month {
  case 4, 6, 9, 11:
    daysInMonth = 30
  case 2:
    if (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0) {
      daysInMonth = 29
    } else {
      daysInMonth = 28
    }
  }

  days := []string {}
  for i := 1; i <= daysInMonth; i++ {
    days = append(days, strconv.Itoa(i))
  }
  return days
}
