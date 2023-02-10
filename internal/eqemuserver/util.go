package eqemuserver

import "strings"

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func contains(slice []string, val string) bool {
	for _, item := range slice {
		if strings.Contains(val, item) {
			return true
		}
	}
	return false
}
