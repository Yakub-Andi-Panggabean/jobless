package actor

import "time"

type Actor struct {
	CreatedDate time.Time
	CreatedBy   string
	UpdatedDate time.Time
	UpdatedBy   string
}
