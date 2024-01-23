package piscine

func UltimateDivMod(a *int, b *int) {
	valeur_a := *a
	valeur_b := *b

	*a = valeur_a / valeur_b
	*b = valeur_a % valeur_b
}
