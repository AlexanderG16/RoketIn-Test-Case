package main

import (
	"fmt"
	"strconv"
	// "strings"
	// "reflect"
)

const eHour, eMin, eSec = 24, 60, 60
const rHour, rMin, rSec = 10, 100, 100

func checkValidType(a string) int {
	var res int
	var n int
	for n == 0 {
		fmt.Print("Input ", a, ": ")
		n, er := fmt.Scan(&res)
		if n != 1 {
			fmt.Println(er)
		} else {
			break
		}
	}
	return res
}

func clockToString(hour, minute, second int, planet string) {
	h := strconv.Itoa(hour)
	m := strconv.Itoa(minute)
	s := strconv.Itoa(second)
	if hour < 10 {
		h = fmt.Sprintf("0%d", hour)
	}
	if minute < 10 {
		m = fmt.Sprintf("0%d", minute)
	}
	if second < 10 {
		s = fmt.Sprintf("0%d", second)
	}

	fmt.Printf("On %s %s:%s:%s", planet, h, m, s)
}

func earthtoRocketConvert(hour, minute, second int) (int, int, int) {
	var outHour, outMin, outSec int
	total := (hour * eMin * eSec) + (minute * eSec) + second
	totalRocket := (total * rHour * rMin * rSec) / (eHour * eMin * eSec)

	outSec = totalRocket % rSec
	if outMin = totalRocket / rSec; outMin > rMin {
		outHour = outMin / rMin
		outMin %= rMin
	}

	return outHour, outMin, outSec
}

func main() {
	var inpHour, inpMin, inpSec int
	inpHour = checkValidType("Hours")
	inpMin = checkValidType("Minutes")
	inpSec = checkValidType("Seconds")

	outHour, outMin, outSec := earthtoRocketConvert(inpHour, inpMin, inpSec)

	clockToString(inpHour, inpMin, inpSec, "Earth")
	fmt.Print(", ")
	clockToString(outHour, outMin, outSec, "RoketIn")
}
