package student

import "github.com/01-edu/z01"

func main() {
	Raid1a(1, 1)

}
func Raid1a(x, y int) {
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			if i == 0 && j == 0 {
				z01.PrintRune(111)
			} else if i == 0 && j == y-1 {
				z01.PrintRune(111)
			} else if i == x-1 && j == 0 {
				z01.PrintRune(111)
			} else if i == x-1 && j == y-1 {
				z01.PrintRune(111)
			} else if j == 0 || j == y-1 {
				z01.PrintRune(45)
			} else if i == 0 || i == x-1 {
				z01.PrintRune(124)
			} else {
				z01.PrintRune(32)
			}

		}
		z01.PrintRune(10)
	}
}
