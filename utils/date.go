/*
 * Copyright © 2016-2022 Iury Braun
 * Copyright © 2017-2022 Neo7i
 */

package utils

import (
	"fmt"
	//"log"
    //"math"
    "time"
    //"strconv"
    //"strings"
)

func MonthIntToString(month int32) string {
	names := []string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}
    
    if month > 0 {
        return names[month-1]
    } else {
        return ""
    }
}

func LastMonthFirstAndLastDays() (time.Time, time.Time) {
    /* ******************** */
    now := time.Now()
    currentYear, currentMonth, _ := now.Date()
    //currentLocation := now.Location()
    currentTimestamp := time.Now().UTC()
    currentLocation := currentTimestamp.Location()
    
    var firstOfDate, lastOfDate time.Time
    start_date := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
    start_date = start_date.AddDate(0, -1, 0)
    end_date := start_date.AddDate(0, 1, -1)
    
    //currentYear, currentMonth, _ := currentTimestamp.Date()
    firstDateOfYear, firstDateOfMonth, firtDateOfDay := start_date.Date()
    lastDateOfYear, lastDateOfMonth, lastDateOfDay := end_date.Date()
    
    firstOfDate = time.Date(firstDateOfYear, firstDateOfMonth, firtDateOfDay, 0, 0, 0, 0, currentLocation)
    lastOfDate = time.Date(lastDateOfYear, lastDateOfMonth, lastDateOfDay, 23, 59, 59, 999999999, currentLocation)
    
    /*pipeline := []bson.M{}
    pipeline = append(pipeline, bson.M{
            "$match": bson.M{"network.master": args_username, "date": bson.M{"$gte": firstOfDate, "$lt": lastOfDate}},
        })*/
    /* ******************** */
    
    return firstOfDate, lastOfDate
}

func CurrentMonthFirstAndLastDays() (time.Time, time.Time) {
    /* ******************** */
    now := time.Now()
    currentYear, currentMonth, _ := now.Date()
    //currentLocation := now.Location()
    currentTimestamp := time.Now().UTC()
    currentLocation := currentTimestamp.Location()
    
    var firstOfDate, lastOfDate time.Time
    start_date := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
    end_date := start_date.AddDate(0, 1, -1)
    
    //currentYear, currentMonth, _ := currentTimestamp.Date()
    firstDateOfYear, firstDateOfMonth, firtDateOfDay := start_date.Date()
    lastDateOfYear, lastDateOfMonth, lastDateOfDay := end_date.Date()
    
    firstOfDate = time.Date(firstDateOfYear, firstDateOfMonth, firtDateOfDay, 0, 0, 0, 0, currentLocation)
    lastOfDate = time.Date(lastDateOfYear, lastDateOfMonth, lastDateOfDay, 23, 59, 59, 999999999, currentLocation)
    
    /*pipeline := []bson.M{}
    pipeline = append(pipeline, bson.M{
            "$match": bson.M{"network.master": args_username, "date": bson.M{"$gte": firstOfDate, "$lt": lastOfDate}},
        })*/
    /* ******************** */
    
    return firstOfDate, lastOfDate
}

func CurrentMonthFirstAndLastTimeOfDay() (time.Time, time.Time) {
    now := time.Now()
    currentYear, currentMonth, currentDay := now.Date()
    currentTimestamp := time.Now().UTC()
    currentLocation := currentTimestamp.Location()
    
	var stat_date, end_date time.Time
	curr_date := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, currentLocation)
	dateYear, dateMonth, dateDay := curr_date.Date()
	
	stat_date = time.Date(dateYear, dateMonth, dateDay, 0, 0, 0, 0, time.Local)
	end_date = time.Date(dateYear, dateMonth, dateDay, 23, 59, 59, 999999999, time.Local)
    
    return stat_date, end_date
}

func StartAndEndDateFromDate(start_date, end_date time.Time) (time.Time, time.Time) {
	start_dateYear, start_dateMonth, start_dateDay := start_date.Date()
	end_dateYear, end_dateMonth, end_dateDay := end_date.Date()
	currentTimestamp := time.Now().UTC()
	currentLocation := currentTimestamp.Location()

	var stat_date_r, end_date_r time.Time
	start_date_date := time.Date(start_dateYear, start_dateMonth, start_dateDay, 0, 0, 0, 0, currentLocation)
	dateYear, dateMonth, dateDay := start_date_date.Date()

	stat_date_r = time.Date(dateYear, dateMonth, dateDay, 0, 0, 0, 0, time.Local)
	end_date_r = time.Date(end_dateYear, end_dateMonth, end_dateDay, 23, 59, 59, 999999999, time.Local)

	return stat_date_r, end_date_r
}


func GetMonthStartAndEndFromYearAndMonth(year, month int) (time.Time, time.Time, error) {
	if month < 1 || month > 12 {
		return time.Time{}, time.Time{}, fmt.Errorf("mês inválido: %d. O mês deve estar entre 1 e 12", month)
	}

	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0).Add(-time.Nanosecond)

	return startDate, endDate, nil
}

func ExtractComponentsDate(date time.Time) (year, month, day int) {
	year = date.Year()
	month = int(date.Month()) // Month() retorna um tipo time.Month, convertemos para int
	day = date.Day()
	return
}

/**
type Month uint32

const (
    // iota automatically increases by 1 on each line
	January  Month = 1 << iota // 1 << 0
	February                   // 1 << 1
	March                      // 1 << 2
	April                      // 1 << 3
	May                        // 1 << 4
	June                       // ...
	July
	August
	September
	October
	November
	December

	AllMonths = January | February | March | April | May | June | July |
		August | September | October | November | December
)

// String returns the name of a month
//
// This attaches a method String to Month constant
//
// Tip: You can use Stringer to create string representation of enums automatically
//      https://godoc.org/golang.org/x/tools/cmd/stringer
func (m Month) String() string {
	//names := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	names := []string{"Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho", "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro"}

	// if it's power of 2, then it's a singular Month value
	if m&(m-1) == 0 {
		// returns the index of the bit to select the month name from `names`
		// array
		return names[int(math.Log2(float64(m)))]
	} else {
		// for composites, just return its number
		// exercise for you: return the names of the months in a composite Month
		return strconv.Itoa(int(m))
	}
}*/
