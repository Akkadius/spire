package questapi

import "unicode"

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// finds slice index in a slice
func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func getSpecialEventTypes(eventType string) []string {
	if eventType == "BotMerc" {
		return []string{"Bot", "Merc"}
	} else if eventType == "BotMercNPC" {
		return []string{"Bot", "Merc", "NPC"}
	} else if eventType == "MercNPC" {
		return []string{"Merc", "NPC"}
	} else if eventType == "Mob" {
		return []string{"Bot", "Merc", "NPC", "Player"}
	}

	return []string{eventType}
}

func isSpecialEventType(eventType string) bool {
	specialEventTypes := []string{"BotMerc", "BotMercNPC", "MercNPC", "Mob"}
	for _, specialEventType := range specialEventTypes {
		if eventType == specialEventType {
			return true
		}
	}

	return false
}
