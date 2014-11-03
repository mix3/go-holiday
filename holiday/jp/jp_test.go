package jp

import (
	"testing"
	"time"

	"github.com/mix3/go-holiday"
	"github.com/stretchr/testify/assert"
)

func TestFixed(t *testing.T) {
	h, _ := holiday.New("jp")
	assert.Equal(t, "", h.HolidayName(1948, time.Month(1), 1))
	assert.Equal(t, "元日", h.HolidayName(1949, time.Month(1), 1))
	assert.Equal(t, "元日", h.HolidayName(9999, time.Month(1), 1))
	assert.Equal(t, "", h.HolidayName(1948, time.Month(1), 15))
	assert.Equal(t, "成人の日", h.HolidayName(1949, time.Month(1), 15))
	assert.Equal(t, "成人の日", h.HolidayName(1999, time.Month(1), 15))
	assert.Equal(t, "", h.HolidayName(2000, time.Month(1), 15))
	assert.Equal(t, "", h.HolidayName(1966, time.Month(2), 11))
	assert.Equal(t, "建国記念の日", h.HolidayName(1967, time.Month(2), 11))
	assert.Equal(t, "建国記念の日", h.HolidayName(9999, time.Month(2), 11))
	assert.Equal(t, "", h.HolidayName(1948, time.Month(4), 29))
	assert.Equal(t, "天皇誕生日", h.HolidayName(1949, time.Month(4), 29))
	assert.Equal(t, "天皇誕生日", h.HolidayName(1988, time.Month(4), 29))
	assert.Equal(t, "みどりの日", h.HolidayName(1989, time.Month(4), 29))
	assert.Equal(t, "みどりの日", h.HolidayName(2006, time.Month(4), 29))
	assert.Equal(t, "昭和の日", h.HolidayName(2007, time.Month(4), 29))
	assert.Equal(t, "昭和の日", h.HolidayName(9999, time.Month(4), 29))
	assert.Equal(t, "", h.HolidayName(1948, time.Month(5), 3))
	assert.Equal(t, "憲法記念日", h.HolidayName(1949, time.Month(5), 3))
	assert.Equal(t, "憲法記念日", h.HolidayName(9999, time.Month(5), 3))
	assert.Equal(t, "国民の休日", h.HolidayName(2006, time.Month(5), 4))
	assert.Equal(t, "みどりの日", h.HolidayName(2007, time.Month(5), 4))
	assert.Equal(t, "みどりの日", h.HolidayName(9999, time.Month(5), 4))
	assert.Equal(t, "", h.HolidayName(1948, time.Month(5), 5))
	assert.Equal(t, "こどもの日", h.HolidayName(1949, time.Month(5), 5))
	assert.Equal(t, "こどもの日", h.HolidayName(9999, time.Month(5), 5))
	assert.Equal(t, "", h.HolidayName(1995, time.Month(7), 20))
	assert.Equal(t, "海の日", h.HolidayName(1996, time.Month(7), 20))
	assert.Equal(t, "海の日", h.HolidayName(2002, time.Month(7), 20))
	assert.Equal(t, "", h.HolidayName(2003, time.Month(7), 20))
	assert.Equal(t, "", h.HolidayName(1965, time.Month(9), 15))
	assert.Equal(t, "敬老の日", h.HolidayName(1966, time.Month(9), 15))
	assert.Equal(t, "敬老の日", h.HolidayName(2002, time.Month(9), 15))
	assert.Equal(t, "敬老の日", h.HolidayName(2003, time.Month(9), 15)) // たまたま第三月曜日
	assert.Equal(t, "", h.HolidayName(2004, time.Month(9), 15))
	assert.Equal(t, "", h.HolidayName(1965, time.Month(10), 10))
	assert.Equal(t, "体育の日", h.HolidayName(1966, time.Month(10), 10))
	assert.Equal(t, "体育の日", h.HolidayName(1999, time.Month(10), 10))
	assert.Equal(t, "", h.HolidayName(2000, time.Month(10), 10))
	assert.Equal(t, "", h.HolidayName(1947, time.Month(11), 3))
	assert.Equal(t, "文化の日", h.HolidayName(1948, time.Month(11), 3))
	assert.Equal(t, "文化の日", h.HolidayName(9999, time.Month(11), 3))
	assert.Equal(t, "", h.HolidayName(1947, time.Month(11), 23))
}

func TestFloat(t *testing.T) {
	h, _ := holiday.New("jp")
	assert.Equal(t, "", h.HolidayName(1999, time.Month(1), 11))
	assert.Equal(t, "成人の日", h.HolidayName(2000, time.Month(1), 10))
	assert.Equal(t, "", h.HolidayName(2002, time.Month(7), 15))
	assert.Equal(t, "海の日", h.HolidayName(2003, time.Month(7), 21))
	assert.Equal(t, "", h.HolidayName(2001, time.Month(9), 17))
	assert.Equal(t, "敬老の日", h.HolidayName(2002, time.Month(9), 15)) // たまたま
	assert.Equal(t, "敬老の日", h.HolidayName(2003, time.Month(9), 15))
	assert.Equal(t, "", h.HolidayName(1998, time.Month(10), 12))
	assert.Equal(t, "振替休日", h.HolidayName(1999, time.Month(10), 11)) // たまたま
	assert.Equal(t, "体育の日", h.HolidayName(2000, time.Month(10), 9))
}

func TestBetween(t *testing.T) {
	h, _ := holiday.New("jp")
	assert.Equal(t, "", h.HolidayName(1984, time.Month(5), 4))
	assert.Equal(t, "", h.HolidayName(1985, time.Month(5), 4))
	assert.Equal(t, "", h.HolidayName(1986, time.Month(5), 4))
	assert.Equal(t, "振替休日", h.HolidayName(1987, time.Month(5), 4))
	assert.Equal(t, "国民の休日", h.HolidayName(1988, time.Month(5), 4))
	assert.Equal(t, "国民の休日", h.HolidayName(2009, time.Month(9), 22))
}

func TestSpecialHoliday(t *testing.T) {
	h, _ := holiday.New("jp")
	assert.Equal(t, "皇太子明仁親王の結婚の儀", h.HolidayName(1959, time.Month(4), 10))
	assert.Equal(t, "皇太子明仁親王の結婚の儀", h.HolidayName(1959, time.Month(4), 10))
	assert.Equal(t, "昭和天皇の大喪の礼", h.HolidayName(1989, time.Month(2), 24))
	assert.Equal(t, "即位礼正殿の儀", h.HolidayName(1990, time.Month(11), 12))
	assert.Equal(t, "皇太子徳仁親王の結婚の儀", h.HolidayName(1993, time.Month(6), 9))
}

func TestNormal(t *testing.T) {
	h, _ := holiday.New("jp")
	assert.False(t, h.IsHoliday(2006, 4, 3))
	assert.False(t, h.IsHolidayName(2006, 4, 3))
	assert.Equal(t, "", h.HolidayName(2006, 4, 3))

	assert.True(t, h.IsHoliday(2006, 4, 2))
	assert.False(t, h.IsHolidayName(2006, 4, 2))
	assert.Equal(t, "", h.HolidayName(2006, 4, 2))
}

func TestETCHoliday(t *testing.T) {
	d := JPDriver{}
	assert.Equal(t, "20060321", d.vernalEquinox(2006))
	assert.Equal(t, "20060923", d.autumnalEquinox(2006))
}

var testCase = []struct {
	year  int
	month int
	day   int
	name  string
}{
	{1999, 1, 1, "元日"},
	{1999, 1, 15, "成人の日"},
	{1999, 2, 11, "建国記念の日"},
	{1999, 3, 21, "春分の日"},
	{1999, 3, 22, "振替休日"},
	{1999, 4, 29, "みどりの日"},
	{1999, 5, 3, "憲法記念日"},
	{1999, 5, 4, "国民の休日"},
	{1999, 5, 5, "こどもの日"},
	{1999, 7, 20, "海の日"},
	{1999, 9, 15, "敬老の日"},
	{1999, 9, 23, "秋分の日"},
	{1999, 10, 10, "体育の日"},
	{1999, 10, 11, "振替休日"},
	{1999, 11, 3, "文化の日"},
	{1999, 11, 23, "勤労感謝の日"},
	{1999, 12, 23, "天皇誕生日"},

	{2006, 1, 1, "元日"},
	{2006, 1, 2, "振替休日"},
	{2006, 1, 9, "成人の日"},
	{2006, 2, 11, "建国記念の日"},
	{2006, 3, 21, "春分の日"},
	{2006, 4, 29, "みどりの日"},
	{2006, 5, 3, "憲法記念日"},
	{2006, 5, 4, "国民の休日"},
	{2006, 5, 5, "こどもの日"},
	{2006, 7, 17, "海の日"},
	{2006, 9, 18, "敬老の日"},
	{2006, 9, 23, "秋分の日"},
	{2006, 10, 9, "体育の日"},
	{2006, 11, 03, "文化の日"},
	{2006, 11, 23, "勤労感謝の日"},
	{2006, 12, 23, "天皇誕生日"},

	{2009, 1, 1, "元日"},
	{2009, 1, 12, "成人の日"},
	{2009, 2, 11, "建国記念の日"},
	{2009, 3, 20, "春分の日"},
	{2009, 4, 29, "昭和の日"},
	{2009, 5, 3, "憲法記念日"},
	{2009, 5, 4, "みどりの日"},
	{2009, 5, 5, "こどもの日"},
	{2009, 5, 6, "振替休日"},
	{2009, 7, 20, "海の日"},
	{2009, 9, 21, "敬老の日"},
	{2009, 9, 22, "国民の休日"},
	{2009, 9, 23, "秋分の日"},
	{2009, 10, 12, "体育の日"},
	{2009, 11, 3, "文化の日"},
	{2009, 11, 23, "勤労感謝の日"},
	{2009, 12, 23, "天皇誕生日"},

	{2010, 1, 1, "元日"},
	{2010, 1, 11, "成人の日"},
	{2010, 2, 11, "建国記念の日"},
	{2010, 3, 21, "春分の日"},
	{2010, 3, 22, "振替休日"},
	{2010, 4, 29, "昭和の日"},
	{2010, 5, 3, "憲法記念日"},
	{2010, 5, 4, "みどりの日"},
	{2010, 5, 5, "こどもの日"},
	{2010, 7, 19, "海の日"},
	{2010, 9, 20, "敬老の日"},
	{2010, 9, 23, "秋分の日"},
	{2010, 10, 11, "体育の日"},
	{2010, 11, 3, "文化の日"},
	{2010, 11, 23, "勤労感謝の日"},
	{2010, 12, 23, "天皇誕生日"},

	{2011, 1, 1, "元日"},
	{2011, 1, 10, "成人の日"},
	{2011, 2, 11, "建国記念の日"},
	{2011, 3, 21, "春分の日"},
	{2011, 4, 29, "昭和の日"},
	{2011, 5, 3, "憲法記念日"},
	{2011, 5, 4, "みどりの日"},
	{2011, 5, 5, "こどもの日"},
	{2011, 7, 18, "海の日"},
	{2011, 9, 19, "敬老の日"},
	{2011, 9, 23, "秋分の日"},
	{2011, 10, 10, "体育の日"},
	{2011, 11, 3, "文化の日"},
	{2011, 11, 23, "勤労感謝の日"},
	{2011, 12, 23, "天皇誕生日"},
}

func TestCase(t *testing.T) {
	h, _ := holiday.New("jp")
	for _, v := range testCase {
		assert.Equal(t, v.name, h.HolidayName(v.year, time.Month(v.month), v.day))
	}
}

func TestHolidays(t *testing.T) {
	expect := map[int]map[string]string{
		1949: map[string]string{
			"01-01": "元日",
			"01-15": "成人の日",
			"03-21": "春分の日",
			"04-29": "天皇誕生日",
			"05-03": "憲法記念日",
			"05-05": "こどもの日",
			"09-23": "秋分の日",
			"11-03": "文化の日",
			"11-23": "勤労感謝の日",
		},
		1950: map[string]string{
			"01-01": "元日",
			"01-15": "成人の日",
			"03-21": "春分の日",
			"04-29": "天皇誕生日",
			"05-03": "憲法記念日",
			"05-05": "こどもの日",
			"09-23": "秋分の日",
			"11-03": "文化の日",
			"11-23": "勤労感謝の日",
		},
	}
	h, _ := holiday.New("jp")
	assert.Equal(t, expect, h.Holidays(1949, 1950))
}
