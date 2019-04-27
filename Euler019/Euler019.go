package main

func main() {
	sundays := 0

	days := 0
	months := 0

	for months/12 < 100 {
		days += getDays(months%12, 1900+months/12)
		// So basically the baseline is Jan 1st 1900, which was a Monday, but we're interested
		// in the mondays from Jan 1st 1901 to Dec 31st 2000, so ignore the first year
		if months >= 12 {
			// Because the baseline is a Monday, that's 6 days from Sunday
			if (days+6)%7 == 0 {
				sundays++
			}
		}
		months++
	}

	println(sundays)
}

func getDays(month, year int) int {
	switch month {
	case 1:
		if isLeap(year) {
			return 29
		}
		return 28
	case 3, 5, 8, 10:
		return 30
	default:
		return 31
	}
}

func isLeap(year int) bool {
	if year%100 == 0 {
		return year%400 != 0
	}
	return year%4 == 0
}
