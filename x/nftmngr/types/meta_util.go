package types

import (
	"time"
	"strconv"
)

func (m *Metadata) getBlockTimestamp() time.Time {
	return m.GetBlockTimeFunction()
}

func (m *Metadata) GetBlockTimestamp(format string) string {
	return m.getBlockTimestamp().Format(format)
}

// 2022-12-15T07:00:57Z
func (m *Metadata) GetUTCBlockTimestamp(format string) string {
	return m.getBlockTimestamp().UTC().Format(format)
}

// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31
func (m *Metadata) GetDay() string {
	return strconv.FormatInt(int64(m.getBlockTimestamp().Day()),10)
}

// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12
func (m *Metadata) GetHour() int {
	return m.getBlockTimestamp().Hour()
}

// 1, 2, 3, ... 58, 59, 60
func (m *Metadata) GetMinute() int {
	return m.getBlockTimestamp().Minute()
}

// 1, 2, 3, ... 58, 59, 60
func (m *Metadata) GetSecond() int {
	return m.getBlockTimestamp().Second()
}

// monday, tuesday, wednesday, thursday, friday, saturday, sunday
func (m *Metadata) GetWeekday() string {
	return m.getBlockTimestamp().Weekday().String()
}

// mon, tue, wed, thu, fri, sat, sun
func (m *Metadata) GetWeekdayShort() string {
	return m.getBlockTimestamp().Weekday().String()[0:3]
}

// 1, 2, 3, 4, 5, 6, 7
func (m *Metadata) GetWeekdayNumber() int {
	return int(m.getBlockTimestamp().Weekday())
}

// january, february, march, april, may, june, july, august, september, october, november, december
func (m *Metadata) GetMonth() string {
	return m.getBlockTimestamp().Month().String()
}

// jan, feb, mar, apr, may, jun, jul, aug, sep, oct, nov, dec
func (m *Metadata) GetMonthShort() string {
	return m.getBlockTimestamp().Month().String()[0:3]
}

// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12
func (m *Metadata) GetMonthNumber() int {
	return int(m.getBlockTimestamp().Month())
}