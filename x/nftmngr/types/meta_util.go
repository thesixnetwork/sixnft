package types

import (
	"strconv"
	"time"
)

func (m *Metadata) getBlockHeight() int64 {
	return m.GetBlockHeightFunction()
}

func (m *Metadata) getBlockTimestamp() time.Time {
	return m.GetBlockTimeFunction()
}

func (m *Metadata) GetBlockHeight() string {
	return strconv.FormatInt(m.getBlockHeight(), 10)
}

// 2022-12-15T07:00:57Z
func (m *Metadata) GetUTCBlockTimestamp(format string) string {
	return m.getBlockTimestamp().UTC().Format(format)
}

// TZ name is from https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
func (m *Metadata) GetBlockTimestampByZone(zone string ,format string) string {
	loc, _ := time.LoadLocation(zone)
	blockTIme := m.getBlockTimestamp().In(loc)
	return blockTIme.Format(format)
}

func (m *Metadata) GetLocalBlockTimestamp(format string) string {
	return m.getBlockTimestamp().Local().Format(format)
}