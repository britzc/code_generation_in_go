//go:generate stringer -type=Season

package main

type Season int

const (
	Spring Season = iota
	Summer
	Autumn
	Winter
)
