package types

import "time"

func (m *Metadata) GetBlockTimestamp() time.Time {
	return m.GetBlockTimeFunction()
}
