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
