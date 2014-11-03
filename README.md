# go-holiday

This module is a go port of [DateTime::Holiday::Japanese](http://coderepos.org/share/browser/lang/perl/DateTime-Holiday-Japanese)

## SYNOPSIS

```
package main

import (
	"fmt"

	"github.com/mix3/go-holiday"
	_ "github.com/mix3/go-holiday/holiday/jp"
)

func main() {
	hd, err := holiday.New("jp")
	if err != nil {
		panic(err)
	}
	if hd.IsHolidayName(2014, 1, 1) {
		fmt.Println("holiday")
	}
	if hd.IsHolidayName(2014, 1, 5) {
		fmt.Println("holiday")
	}
	if hd.IsHoliday(2014, 1, 1) {
		fmt.Println("with sunday")
	}
	if hd.IsHoliday(2014, 1, 5) {
		fmt.Println("with sunday")
	}
	if name := hd.HolidayName(2014, 1, 1); name != "" {
		fmt.Println(name) // 元旦
	}
	holidays := hd.Holidays(2014, 2015)
	for year, _ := range holidays {
		for md, name := range holidays[year] {
			fmt.Printf("%04d-%s: %s\n", year, md, name)
		}
	}
}
```

output
```
holiday
with sunday
with sunday
元日
2014-05-04: みどりの日
2014-05-06: 振替休日
2014-11-23: 勤労感謝の日
2014-12-23: 天皇誕生日
2014-03-21: 春分の日
2014-04-29: 昭和の日
2014-05-05: こどもの日
.
.
.
```

## Copyright

Copyright (c) 2014 mix3. See [LICENSE](LICENSE) for details.
