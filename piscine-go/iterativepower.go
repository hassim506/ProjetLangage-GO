package piscine

func IterativePower(nb int, power int) int {
	a := 1
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	for i := 1; i <= power; i++ {
		a = a * nb
	}
	return a
}
