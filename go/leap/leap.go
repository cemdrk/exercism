// Package leap provides utility functions for leap year.
package leap

// IsLeapYear decides if a leap year or not in the Gregorian calendar.
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
