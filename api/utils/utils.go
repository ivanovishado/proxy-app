package utils

import (
	env "github.com/joho/godotenv"
)

// LoadEnvVars should load the .env file
func LoadEnvVars() {
	env.Load()
}

// PriorityLevel is the type of the enum
type PriorityLevel int

// PriorityLevels to be used by the prioritization algorithm
const (
	Low PriorityLevel = iota
	Medium
	High
)

// ShiftArray inserts an element to a strings list and shifts the following elements
func ShiftArray(array *[]string, position int, value string) {
	//  extend array by one
	*array = append(*array, "")

	// shift values
	copy((*array)[position+1:], (*array)[position:])

	// insert value
	(*array)[position] = value
}
