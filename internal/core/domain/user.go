package domain

import "time"

type PackageType int

const (
	// FreeTrial ...
	FreeTrial PackageType = iota
	// Basic ...
	Basic
	// Pro ...
	Pro
)

// User ...
// strcture to interact with User Model from Repository
type User struct {
	ID             string
	Signature      string
	PackageType    string
	IsActive       bool
	ExpirationTime time.Time
}
