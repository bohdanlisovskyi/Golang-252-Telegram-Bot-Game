package algo

const incMult = 0.1

//ComputeIncreaseOfPeople - return increase of people
func ComputeIncreaseOfPeople(currentAmount, maxAmount int) int {
	inc := int(float64(currentAmount) * incMult)
	if currentAmount+inc > maxAmount {
		return maxAmount - currentAmount
	}
	return inc
}
