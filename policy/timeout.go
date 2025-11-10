// Package policy @author: Violet-Eva @date  : 2025/9/3 @notes :
package policy

import (
	"strings"
	"time"
)

// validitySchedulesParse
// @Description:
// @param input 2006/1/2 15:04:05
// @return output 2006-01-02 15:04:05
// @return err
func validitySchedulesParse(input string) (output string) {
	var (
		year, mount, day     string
		hour, minute, second string
	)
	timeArr := strings.Split(input, " ")
	splitYMD := strings.Split(timeArr[0], "/")
	year = splitYMD[0]
	if len(splitYMD[1]) != 2 {
		mount = "0" + splitYMD[1]
	} else {
		mount = splitYMD[1]
	}
	if len(splitYMD[2]) != 2 {
		day = "0" + splitYMD[2]
	} else {
		day = splitYMD[2]
	}

	splitHMS := strings.Split(timeArr[1], ":")
	if len(splitHMS[0]) != 2 {
		hour = "0" + hour
	} else {
		hour = splitHMS[0]
	}
	if len(splitHMS[1]) != 2 {
		minute = "0" + minute
	} else {
		minute = splitHMS[1]
	}
	if len(splitHMS[2]) != 2 {
		second = "0" + second
	} else {
		second = splitHMS[2]
	}

	output = year + "-" + mount + "-" + day + " " + hour + ":" + minute + ":" + second

	return
}

func (b *Policy) ParseValiditySchedules() (output []string) {
	for _, vs := range b.ValiditySchedules {
		startTime := validitySchedulesParse(vs.StartTime)
		endTime := validitySchedulesParse(vs.EndTime)
		tmpStr := startTime + "~" + endTime + "~" + vs.TimeZone
		output = append(output, tmpStr)
	}
	return
}

func (b *Policy) JudgeTimeout() (isTimeout bool, err error) {
	vss := b.ParseValiditySchedules()
	for _, vs := range vss {
		timeArr := strings.Split(vs, "~")
		var location *time.Location
		var parse time.Time
		location, err = time.LoadLocation(timeArr[2])
		if err != nil {
			return
		}
		parse, err = time.ParseInLocation("2006-01-02 15:04:05", timeArr[1], location)
		if err != nil {
			return
		}
		localTime := parse.Local()
		if time.Now().Local().After(localTime) {
			isTimeout = true
		} else {
			isTimeout = false
		}
	}
	return
}
