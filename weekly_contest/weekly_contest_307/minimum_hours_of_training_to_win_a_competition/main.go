package main

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	currentEnergy := initialEnergy
	currentExperience := initialExperience

	hours := 0
	for i := 0; i < len(energy); i++ {
		if currentEnergy <= energy[i] {
			diff := (energy[i] - currentEnergy) + 1
			currentEnergy += diff;
			hours += diff
		}
		if currentExperience <= experience[i] {
			diff := (experience[i] - currentExperience) + 1
			currentExperience += diff
			hours += diff
		}
		currentEnergy -= energy[i]
		currentExperience += experience[i]
	}
	return hours
}
