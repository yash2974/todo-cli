package google

import (
	"TODOCLI/utility"
	"TODOCLI/auth"
)

func ReadCalender() {
	utility.PrintCMD("hey google test", "Green")
	auth.Calender()
}