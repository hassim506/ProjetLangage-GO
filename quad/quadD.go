package piscine

import (
	"github.com/01-edu/z01"
)

func QuadD(x, y int) {
	if x > 0 && y > 0 {

		drawWidth(x)
		drawHeight(x, y)
		if y > 1 {
			drawWidth2(x)
		}

	}
}

func drawHeight(x, y int) {
	if y > 2 {
		for i := 0; i < y-2; i++ {
			if x == 1 {
				z01.PrintRune('B')
				z01.PrintRune('\n')

			} else {
				z01.PrintRune('B')

				for i := 0; i < x-2; i++ {
					z01.PrintRune(' ')
				}
				z01.PrintRune('B')
				z01.PrintRune('\n')
			}
		}
	}
}

func drawWidth(x int) {
	if x == 1 {
		z01.PrintRune('A')
		z01.PrintRune('\n')
	} else if x == 2 {
		z01.PrintRune('A')
		z01.PrintRune('C')
		z01.PrintRune('\n')

	} else {
		z01.PrintRune('A')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
		z01.PrintRune('\n')

	}
}

func drawWidth2(x int) {
	if x == 1 {
		z01.PrintRune('A')
		z01.PrintRune('\n')
	} else if x == 2 {
		z01.PrintRune('A')
		z01.PrintRune('C')
		z01.PrintRune('\n')

	} else {
		z01.PrintRune('A')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
		z01.PrintRune('\n')

	}
}
