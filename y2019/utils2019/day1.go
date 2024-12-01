package utils2019

func GetFuelFromMass(mass int) int {
	return mass/3 - 2
}

func GetTotalFuelFromMass(mass int) int {
	total := 0

	fuel := GetFuelFromMass(mass)

	for fuel > 0 {
		total += fuel
		fuel = GetFuelFromMass(fuel)
	}

	return total
}
