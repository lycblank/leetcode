package main

func main() {

}

func dayOfTheWeek(day int, month int, year int) string {
	// 1971年1月1日为星期五
	weeks := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	days := (year - 1971)*365 + (year-1971-1)/4
	if (year-1971-1)%4 != 0 {
		days++
	}
	for i := 1; i < month; i++ {
		days += GetMonthDay(year, i)
	}
	days += day-1
	days %= 7
	return weeks[(5+days)%7]
}

func IsLeap(year int) bool {
	return (year % 4 == 0 && year % 100 != 0) || (year % 400 == 0)
}

func GetMonthDay(year int, month int) int {
	monthDays := []int{31,28,31,30,31,30, 31, 31, 30, 31, 30, 31}
	days := monthDays[month-1]
	if month == 2 && IsLeap(year) {
		days += 1
	}
	return days
}



