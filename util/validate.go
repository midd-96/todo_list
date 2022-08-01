package util

import (
	"strings"
)

func ValidateEmail(email string) bool {

	if strings.Contains(email, "@") && strings.Contains(email, ".") {
		return true
	} else {
		return false
	}

}

func ValidatePhone(phone string) bool {

	count := strings.Split(phone, "")
	//fmt.Println("count: ", len(count))
	if len(count) == 10 {
		return true
	} else {
		return false
	}

}
