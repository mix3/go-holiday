package holiday

import (
	"fmt"
	"time"

	"github.com/mix3/go-holiday/driver"
)

var drivers = make(map[string]driver.Driver)

func Register(name string, driver driver.Driver) {
	if driver == nil {
		panic("holiday: Register driver is nil")
	}
	if _, dup := drivers[name]; dup {
		panic("holiday: Register called twice for driver" + name)
	}
	drivers[name] = driver
}

type Holiday struct {
	driver driver.Driver
}

func New(driverName string) (*Holiday, error) {
	driveri, ok := drivers[driverName]
	if !ok {
		return nil, fmt.Errorf("holiday: unknown driver %q (forgotten import?)", driverName)
	}
	holiday := &Holiday{driveri}
	return holiday, nil
}

var cache = map[string]time.Time{}

func GetTime(year int, month time.Month, day int) (time.Time, string) {
	ymd := fmt.Sprintf("%04d%02d%02d", year, month, day)
	if v, ok := cache[ymd]; ok {
		return v, ymd
	}
	t := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	cache[ymd] = t
	return t, ymd
}

func ok(str string) bool {
	return str != ""
}

func (hd *Holiday) IsHoliday(year int, month time.Month, day int) bool {
	t, _ := GetTime(year, month, day)
	if t.Weekday() == 0 {
		return true
	}
	return hd.driver.Holiday(year, month, day) != ""
}

func (hd *Holiday) IsHolidayName(year int, month time.Month, day int) bool {
	return hd.driver.Holiday(year, month, day) != ""
}

func (hd *Holiday) HolidayName(year int, month time.Month, day int) string {
	return hd.driver.Holiday(year, month, day)
}

func (hd *Holiday) Holidays(startYear, endYear int) map[int]map[string]string {
	ret := map[int]map[string]string{}
	for y := startYear; y <= endYear; y++ {
		ret[y] = map[string]string{}
	}

	if endYear < startYear {
		return ret
	}

	st, sYmd := GetTime(startYear, 1, 1)
	_, eYmd := GetTime(endYear+1, 1, 1)
	for {
		st, sYmd = GetTime(st.Date())
		if eYmd <= sYmd {
			return ret
		}
		y, m, d := st.Date()
		md := fmt.Sprintf("%02d-%02d", m, d)
		if name := hd.driver.Holiday(y, m, d); ok(name) {
			ret[y][md] = name
		}
		st = st.AddDate(0, 0, 1)
	}

	return ret
}
