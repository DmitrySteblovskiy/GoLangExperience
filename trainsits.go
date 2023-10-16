package main

import (
	"fmt"
	"math"
)

// ввод: ? d
func ask(a int)

func sitsearch() int {
	sites := [54] int

	left := 0
	right := 53
	sector := ""
	binarity := ""
	bin := ""
	var p int
	d := 0

	for i:=0; i < 52; i++ {
		sites[i+1] = sites [i] + 1
	}

	ask(0)
	fmt.Scanln("%s %s", &binarity, &sector)

	if sector == "main" {
		left = 0
		right = 35
		d := 18
	}

	if sector == "side" {
		left = 36
		right = 53
		d = -9
	}

	for math.Abs(float64(d)) > 1 {
		if sector == "main" {
			if d < 0 {
				d = d * -1
				right = left + d - 1
			} else {
				right = right - d - 1
			}
			// I use indexes instead of actual site number, which is +1
			if (binarity == "high") && (right % 2 == 0) {
				right = right - 1
			}
			if (binarity == "low") && (right % 2)  != 0 {
				right = right - 1
			}
			ask(d)
			fmt.Scanln("%s %s", &bin, &sector)
			d = d / 2
		}
		if sector == "side" {
			if d > 0 {
				d = d * -1
			}
			left = left - d
		}
		if (binarity == "high") && (right % 2 == 0) {
			right = right - 1
		}
		if (binarity == "low") && (right % 2)  != 0 {
			right = right - 1
		}
		ask(d)
		fmt.Scanln("%s %s", &bin, &sector)
		d = d / 2
	}

	if binarity == "high" {
		if right % 2 != 0 {
			return sites[left]
		} else {
			return sites[right]
		}
	}
	if binarity == "low" == 0 {
		if right % 2  == 0 {
			return sites[right]
		} else {
			return sites[left]
		}
	}
	return p
}

func main() {
	ask(sitsearch())
}
/*
	if sector == "side" {
		left = 36
		right = 53
		d = -9
		for math.Abs(float64(d)) > 2 {
			if sector == "main" {
				if d < 0 {
					d = d * -1
					right = left + d - 1
				} else {
					right = d
				}
				//ввод: ? d
				fmt.Scanln("%s %s", &bin, &sector)
				d = d / 2
			}
			if sector == "side" {
				if d > 0 {
					d = d * -1
					left = d
				}
				//ввод: ? d
				fmt.Scanln("%s %s", &bin, &sector)
				d = d / 2

			}
		}
	}

	/*for math.Abs(float64(d)) > 2 {
		if sector == "main" {
			if d < 0 {
				d = d * -1
				right  = left + d -1
			} else {
				right = d
			}
			//ввод: ? d
			fmt.Scanln("%s %s", &bin, &sector)
			d = d / 2
		}
		if sector == "side" {
			if d > 0 {
				d = d * -1
				left = d
			}
			//ввод: ? d
			fmt.Scanln("%s %s", &bin, &sector)
			d = d / 2

		}
	}*/
