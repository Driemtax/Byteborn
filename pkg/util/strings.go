package util

import "strings"

// Indent: Creates indentation with n times double spaces
func Indent(n int) string {
	return strings.Repeat("  ", n)
}
