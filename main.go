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
	first := GetFirstDateOfMonth(start)
	last := GetLastDateOfMonth(end)

	// 同月份
	if first.Month() == last.Month() {
		days := GetDays(first, last)

		if start == end {
			mon := fmt.Sprintf("%d%02d", start.Year(), start.Month())
			return float64(a.dateMap[mon].Amount / days)
		}
	} else {

		var totalCount float64 = 0
		startfirst := GetFirstDateOfMonth(start)
		startLast := GetLastDateOfMonth(start)
		days := GetDays(start, startLast)
		// fmt.Println(startfirst, startLast, days)
		mon := fmt.Sprintf("%d%02d", start.Year(), start.Month())

		totalCount = float64((a.dateMap[mon].Amount / GetDays(startfirst, startLast) * days)) + totalCount

		fmt.Println("totalCount = ", totalCount)
		nextFirst := startfirst
		for {
			nextFirst = nextFirst.AddDate(0, 1, 0)
			nextLast := GetLastDateOfMonth(nextFirst)
			if nextLast.Month() == end.Month() {
				days := GetDays(nextFirst, end)
				// fmt.Println(days)
				nextMon := fmt.Sprintf("%d%02d", end.Year(), end.Month())
				return float64(a.dateMap[nextMon].Amount/GetDays(nextFirst, nextLast)*days) + totalCount
			} else {
				nextMon := fmt.Sprintf("%d%02d", nextFirst.Year(), nextFirst.Month())
				fmt.Println(totalCount, nextMon, a.dateMap[nextMon].Amount)
				totalCount = float64(a.dateMap[nextMon].Amount) + totalCount
			}
		}

		//endFirst := GetFirstDateOfMonth(end)

	}

	startMon := fmt.Sprintf("%d%02d", start.Year(), start.Month())
	// endMon := fmt.Sprintf("%d%02d", end.Year(), end.Month())
	fmt.Print(startMon, a.dateMap)

	fmt.Println("first = ", first)
	fmt.Println("last = ", last)

	return float64(a.dateMap[startMon].Amount)
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
