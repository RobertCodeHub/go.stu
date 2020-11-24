package main

import (
	"fmt"
	"time"
)

type WeekDay int

const(
	Monday WeekDay = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type Flags uint

const(
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}


func main() {

	day := [...]string{Sunday: "SUN", Monday: "MON", Tuesday: "TUE", Wednesday: "WED", Thursday: "THU", Friday : "FRI", Saturday : "SAT"}
	fmt.Println(Monday, day[Monday])

	//获取到当前的日期
	t := time.Now()

	if WeekDay(t.Weekday()) == Monday {
		fmt.Println("Monday")
	}else if WeekDay(t.Weekday()) == Thursday {
		fmt.Println("Thursday")
	}else if WeekDay(t.Weekday()) == Wednesday {
		fmt.Println("Wednesday")
	}else if WeekDay(t.Weekday()) == Thursday {
		fmt.Println("Thursday")
	}else if WeekDay(t.Weekday()) == Friday {
		fmt.Println("Friday")
	}else if WeekDay(t.Weekday()) == Saturday {
		fmt.Println("Saturday")
	}
}