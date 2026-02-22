package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layer []string, time int) int {
	i := len(layer)
	if time == 0 {
		time = 2
	}
	return i * time
}

// TODO: define the 'Quantities()' function
func Quantities(layer []string) (int, float64) {
	weight := 50
	liquid := 0.2
	noodle := 0
	sauce := 0.0

	for _, v := range layer {
		if v == "noodles" {
			noodle++
		} else if v == "sauce" {
			sauce = sauce + 1.0
		}
	}
	return (weight * noodle), (sauce * liquid)
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend []string, mine []string) {
	mine[len(mine)-1] = friend[len(friend)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, number int) []float64 {
	var list []float64
	for _, v := range quantities {
		list = append(list, (v * (float64(number) / 2)))
	}
	return list
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
