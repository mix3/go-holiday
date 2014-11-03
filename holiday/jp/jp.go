package jp

import (
	"fmt"
	"time"

	"github.com/mix3/go-holiday"
)

type JPDriver struct {
}

func init() {
	holiday.Register("jp", &JPDriver{})
}

type staticHoliday struct {
	start int
	end   int
	name  string
}

var fixHoliday = map[string][]staticHoliday{
	"0101": []staticHoliday{{1949, 0, "元日"}},
	"0115": []staticHoliday{{1949, 1999, "成人の日"}},
	"0211": []staticHoliday{{1967, 0, "建国記念の日"}},
	"0429": []staticHoliday{
		{2007, 0, "昭和の日"},
		{1989, 2006, "みどりの日"},
		{1949, 1988, "天皇誕生日"},
	},
	"0503": []staticHoliday{{1949, 0, "憲法記念日"}},
	"0504": []staticHoliday{{2007, 0, "みどりの日"}},
	"0505": []staticHoliday{{1949, 0, "こどもの日"}},
	"0720": []staticHoliday{{1996, 2002, "海の日"}},
	"0915": []staticHoliday{{1966, 2002, "敬老の日"}},
	"1010": []staticHoliday{{1966, 1999, "体育の日"}},
	"1103": []staticHoliday{{1948, 0, "文化の日"}},
	"1123": []staticHoliday{{1948, 0, "勤労感謝の日"}},
	"1223": []staticHoliday{{1989, 0, "天皇誕生日"}},
}

var specialHoliday = map[string]string{
	"19590410": "皇太子明仁親王の結婚の儀",
	"19890224": "昭和天皇の大喪の礼",
	"19901112": "即位礼正殿の儀",
	"19930609": "皇太子徳仁親王の結婚の儀",
}

const (
	COMING_OF_AGE        = "成人の日"
	VERNAL_EQUINOX       = "春分の日"
	NATIONAL_HOLIDAY     = "国民の休日"
	MARINE               = "海の日"
	RESPECT_FOR_THE_AGED = "敬老の日"
	AUTUMNAL_EQUINOX     = "秋分の日"
	HEALTH_AND_SPORTS    = "体育の日"
	SUBSTITUTE_HOLIDAY   = "振替休日"
)

func ok(str string) bool {
	return str != ""
}

func weekdayOfMonth(year int, month time.Month, day int) int {
	t, ymd := holiday.GetTime(year, month, day)
	s, _ := holiday.GetTime(year, month, 1)
	num := 0
	for {
		if ymd < s.Format("20060102") {
			break
		}
		if s.Weekday() == t.Weekday() {
			num++
		}
		s = s.Add(time.Hour * 24)
	}
	return num
}

func (jpd JPDriver) Holiday(year int, month time.Month, day int) string {
	if v := jpd.basicHoliday(year, month, day); v != "" {
		return v
	}
	if v := jpd.changeHoliday(year, month, day); v != "" {
		return v
	}
	if v := jpd.betweenHoliday(year, month, day); v != "" {
		return v
	}
	if v := jpd.specialHoliday(year, month, day); v != "" {
		return v
	}
	return ""
}

func (jpd JPDriver) basicHoliday(year int, month time.Month, day int) string {
	_, ymd := holiday.GetTime(year, month, day)

	if ymd == jpd.vernalEquinox(year) {
		return VERNAL_EQUINOX
	}
	if ymd == jpd.autumnalEquinox(year) {
		return AUTUMNAL_EQUINOX
	}

	if list, ok := fixHoliday[fmt.Sprintf("%02d%02d", month, day)]; ok {
		for _, v := range list {
			if v.start <= year && (0 == v.end || year <= v.end) {
				return v.name
			}
		}
	}

	return jpd.floatHoliday(year, month, day)
}

func (jpd JPDriver) floatHoliday(year int, month time.Month, day int) string {
	t, _ := holiday.GetTime(year, month, day)

	wom := weekdayOfMonth(year, month, day)
	if month == 1 && t.Weekday() == 1 && wom == 2 && 2000 <= year {
		return COMING_OF_AGE
	}
	if month == 7 && t.Weekday() == 1 && wom == 3 && 2003 <= year {
		return MARINE
	}
	if month == 9 && t.Weekday() == 1 && wom == 3 && 2003 <= year {
		return RESPECT_FOR_THE_AGED
	}
	if month == 10 && t.Weekday() == 1 && wom == 2 && 2000 <= year {
		return HEALTH_AND_SPORTS
	}

	return ""
}

func (jpd JPDriver) betweenHoliday(year int, month time.Month, day int) string {
	t, ymd := holiday.GetTime(year, month, day)

	if ymd < "19851227" {
		return ""
	}
	if t.Weekday() == 0 {
		return ""
	}
	if ok(jpd.changeHoliday(year, month, day)) {
		return ""
	}

	prev := t.Add(time.Hour * -24)
	if !ok(jpd.basicHoliday(prev.Date())) {
		return ""
	}

	next := t.Add(time.Hour * 24)
	if !ok(jpd.basicHoliday(next.Date())) {
		return ""
	}

	return NATIONAL_HOLIDAY
}

func (jpd JPDriver) changeHoliday(year int, month time.Month, day int) string {
	t, ymd := holiday.GetTime(year, month, day)

	if ymd < "19730412" {
		return ""
	}

	for {
		t = t.Add(time.Hour * -24)
		if !ok(jpd.basicHoliday(t.Date())) {
			return ""
		}
		if int(t.Weekday()) == 0 {
			return SUBSTITUTE_HOLIDAY
		}
		if year < 2007 {
			return ""
		}
	}

	return ""
}

func (jpd JPDriver) specialHoliday(year int, month time.Month, day int) string {
	_, ymd := holiday.GetTime(year, month, day)
	if v, ok := specialHoliday[ymd]; ok {
		return v
	}
	return ""
}

func (jpd JPDriver) vernalEquinox(year int) string {
	var x, y float64
	if 1900 <= year && year <= 1979 {
		x = 20.8357
		y = 1983.0
	} else if 1980 <= year && year <= 2099 {
		x = 20.8431
		y = 1980.0
	} else if 2100 <= year && year <= 2150 {
		x = 21.8510
		y = 1980.0
	} else {
		return ""
	}
	day := int(x + float64(0.242194)*float64(year-1980) - float64(int((float64(year)-y)/4)))
	return fmt.Sprintf("%04d%02d%02d", year, 3, day)
}

func (jpd JPDriver) autumnalEquinox(year int) string {
	var x, y float64
	if 1900 <= year && year <= 1979 {
		x = 23.2588
		y = 1983.0
	} else if 1980 <= year && year <= 2099 {
		x = 23.2488
		y = 1980.0
	} else if 2100 <= year && year <= 2150 {
		x = 24.2488
		y = 1980.0
	} else {
		return ""
	}
	day := int(x + float64(0.242194)*float64(year-1980) - float64(int((float64(year)-y)/4)))
	return fmt.Sprintf("%04d%02d%02d", year, 9, day)
}
