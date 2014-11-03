package driver

import "time"

type Driver interface {
	Holiday(year int, month time.Month, day int) string
}
