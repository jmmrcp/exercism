// Build constraint.  It's like when you buy something new and you find a
// tag that says "Inspected by #12."  Nice to know it was inspected, but
// you don't need the tag.  Take it off.
// +build !example

// Package clause.
package gigasecond

import "time"

// Constant declaration.
const testVersion = 4 // find the value in gigasecond_test.go

// API function.  It uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1E9 * time.Second)
}

// Reviewers don't think much of stub comments.  Replace or remove.
