package main

import "time"

type Employee struct {
	ID         int
	NAME       string
	ADDRESS    string
	DOB        time.Time
	POSITION   string
	SALARY     int
	MANAGER_Id int
}

func main() {
	var robertHu Employee;
	robertHu.ID         = 202011
	robertHu.NAME       = "RobertHu"
	robertHu.ADDRESS    = "AnHui,HoFei"
	robertHu.POSITION   = "JavaDeveloper"
	robertHu.SALARY     = 8000
	robertHu.MANAGER_Id = 05
}