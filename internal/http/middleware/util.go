package middleware

import "strings"

func stringContains(slice []string, val string) bool {
	for _, item := range slice {
		if strings.Contains(val, item) {
			return true
		}
	}
	return false
}
