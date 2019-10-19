package student

import "github.com/01-edu/z01"

func Raid1d(x, y int) {
	if x <= 0 || y <= 0 {
	} else {
		for j := 0; j < y; j++ {
			for i := 0; i < x; i++ {
				if i == 0 && j == 0 {
					z01.PrintRune(65)
				} else if i == 0 && j == y-1 {
					z01.PrintRune(65)
				} else if i == x-1 && j == 0 {
					z01.PrintRune(67)
				} else if i == x-1 && j == y-1 {
					z01.PrintRune(67)
				} else if j == 0 || j == y-1 {
					z01.PrintRune(66)
				} else if i == 0 || i == x-1 {
					z01.PrintRune(66)
				} else {
					z01.PrintRune(32)
				}

			}
			z01.PrintRune(10)
		}
	}

}
