package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	//type digits [10][5]string

	screen.Clear()
	for {
		//fmt.Println("\f")
		screen.MoveTopLeft()
		now := time.Now()
		hour, min, sec := now.Hour(), now.Minute(), now.Second()
		ssec := now.Nanosecond() / 1e8

		//fmt.Printf("hour: %d, min: %d, sec: %d\n", hour, min, sec)

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			separator,
			digits[min/10], digits[min%10],
			separator,
			digits[sec/10], digits[sec%10],
			dot,
			digits[ssec],
		}

		for line := range clock[0] {

			for index, digit := range clock {
				//fmt.Print(digit)
				//fmt.Print(line)
				next := clock[index][line]
				if (digit == separator || digit == dot) && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, " ")
			}
			fmt.Println()
		}
		fmt.Println()

		time.Sleep(time.Second / 10)
	}

}
