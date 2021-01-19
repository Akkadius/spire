package questapi

import (
	"strings"
)

type PerlEvent struct {
	EventName string   `json:"event_name"`
	EventVars []string `json:"event_vars"`
}

func parsePerlEvents(contents string, fileName string) []PerlEvent {
	perlEvents := []PerlEvent{}
	if strings.Contains(fileName, "embparser.cpp") {
		currentEvent := ""
		events := []string{}
		eventVars := []string{}
		for _, l := range strings.Split(contents, "\n") {

			// break will start a new buffer
			// we could use the same exported vars for multiple event cases
			if strings.Contains(l, "break") && currentEvent != "" {
				for _, event := range events {
					perlEvents = append(
						perlEvents, PerlEvent{
							EventName: event,
							EventVars: eventVars,
						},
					)
				}

				events = []string{}
				eventVars = []string{}
			}

			// keep track of the current event
			if strings.Contains(l, "case EVENT") {
				line := l
				line = strings.ReplaceAll(line, "case", "")
				line = strings.ReplaceAll(line, ":", "")
				line = strings.ReplaceAll(line, "{", "")
				line = strings.TrimSpace(line)

				// set the current event in the buffer
				currentEvent = line

				events = append(events, currentEvent)
			}

			// parse vars
			if strings.Contains(l, "ExportVar") && currentEvent != "" {
				quoteSplit := strings.Split(l, "\"")
				if len(quoteSplit) > 1 {
					eventVars = append(eventVars, strings.TrimSpace(quoteSplit[1]))
				}
			}
		}
	}

	return perlEvents
}
