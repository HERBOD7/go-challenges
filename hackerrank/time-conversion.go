//input: 07:05:45PM
//output: 19:05:45

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	am := strings.Split(s, "AM")
	pm := strings.Split(s, "PM")
	result := ""
	if len(am) > 1 {
		amItems := strings.Split(am[0], ":")
		hr, _ := strconv.Atoi(amItems[0])
		if hr == 12 {
			result = fmt.Sprintf("00:%s:%s", amItems[1], amItems[2])
		} else {
			if hr < 10 {
				result = fmt.Sprintf("0%d:%s:%s", hr, amItems[1], amItems[2])
			} else {
				result = fmt.Sprintf("%d:%s:%s", hr, amItems[1], amItems[2])
			}
		}
	} else if len(pm) > 1 {
		pmItems := strings.Split(pm[0], ":")
		hr, _ := strconv.Atoi(pmItems[0])
		if hr < 12 {
			result = fmt.Sprintf("%d:%s:%s", hr+12, pmItems[1], pmItems[2])
		} else {
			if hr < 10 {
				result = fmt.Sprintf("0%d:%s:%s", hr, pmItems[1], pmItems[2])
			} else {
				result = fmt.Sprintf("%d:%s:%s", hr, pmItems[1], pmItems[2])
			}
		}
	}
	return result
}
