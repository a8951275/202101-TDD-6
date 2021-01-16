package main

import (
	"fmt"
	"time"
)

type IDB interface {
	GetAll() []Budget
}

type DB struct {
	Budget []Budget
}

func (d *DB) GetAll() []Budget {
	return d.Budget
}

type Budget struct {
	Date   string
	Amount int
}

type Accounting struct {
	DB      IDB
	dateMap map[string]Budget
}

func (a *Accounting) TotalAmount(start, end time.Time) float64 {
	a.GetAll()

	currentMon := GetFirstDateOfMonth(start)
	var totalCount float64 = 0
	for {
		monLast := GetLastDateOfMonth(currentMon)
		mon := fmt.Sprintf("%d%02d", currentMon.Year(), currentMon.Month())
		monDays := GetDays(currentMon, monLast)
		var diffDays int

		if currentMon.Before(start) {
			diffDays = GetDays(start, monLast)
		} else if monLast.After(end) {
			diffDays = GetDays(currentMon, end)
		} else {
			diffDays = GetDays(currentMon, monLast)
		}

		totalCount = float64((a.dateMap[mon].Amount / monDays * diffDays)) + totalCount

		if monLast.After(end) {
			return totalCount
		}

		currentMon = currentMon.AddDate(0, 1, 0)
	}
}

func (a *Accounting) GetAll() {
	a.dateMap = make(map[string]Budget)

	data := a.DB.GetAll()
	for _, v := range data {
		a.dateMap[v.Date] = v
	}
}

func main() {}

func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

func GetDays(start, end time.Time) int {
	duration := end.Sub(start) + time.Hour*24
	hours := duration.Hours()
	return int(hours / 24)
}
