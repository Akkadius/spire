package questapi

import (
	"fmt"
	"strings"
)

type PerlEvent struct {
	EntityType      string   `json:"entity_type"`      // @eg Player
	EventName       string   `json:"event_name"`       // @eg EVENT_SAY
	EventIdentifier string   `json:"event_identifier"` // @eg EVENT_SAY
	EventVars       []string `json:"event_vars"`       // @eg []{"text", "langid"}
}

type PerlEventEntityMapping struct {
	Entity string // @eg Player
	Event  string // @eg EVENT_SAY
}

func (c *ParseService) parsePerlEvents(files map[string]string) []PerlEvent {

	// we need to parse the event code constants so we can line them up with the actual
	// strings that are used in the perl api
	// constants used in the code != what is used in the quest api in some circumstances
	// this is to map the two together
	// @example source: EVENT_TRADE perl usage: EVENT_ITEM
	eventCodes := []string{}
	perlEventCodes := []string{}
	for fileName, contents := range files {

		// source event names
		if strings.Contains(fileName, "event_codes.h") {
			for _, l := range strings.Split(contents, "\n") {

				// grep
				// @example EVENT_TRADE,		//being given an item or money
				if strings.Contains(l, "EVENT_") && strings.Contains(l, ",") {
					event := ""
					eventCodeSplit := strings.Split(l, ",")
					if len(eventCodeSplit) > 0 {
						event = strings.TrimSpace(eventCodeSplit[0])
						event = strings.ReplaceAll(event, "0", "")
						event = strings.ReplaceAll(event, "=", "")
						event = strings.TrimSpace(event)
					}

					eventCodes = append(eventCodes, event)
				}
			}
		}

		// perl event names
		if strings.Contains(fileName, "embparser.cpp") {
			for _, l := range strings.Split(contents, "\n") {

				// grep
				// @example EVENT_TRADE,		//being given an item or money
				if strings.Contains(l, "EVENT_") && strings.Contains(l, ",") && strings.Contains(l, "\"") {
					event := l
					event = strings.ReplaceAll(event, "\"", "")
					event = strings.ReplaceAll(event, ",", "")
					event = strings.TrimSpace(event)

					perlEventCodes = append(perlEventCodes, event)
				}
			}
		}
	}

	//fmt.Printf("%+v\n", perlEventCodes)
	//fmt.Printf("%+v\n", len(perlEventCodes))
	//fmt.Printf("%+v\n", eventCodes)
	//fmt.Printf("%+v\n", len(eventCodes))

	// first pass; fetch all events and what entity they relate to (Player, NPC, Item etc.)
	eventEntityMappings := []PerlEventEntityMapping{}
	for fileName, contents := range files {
		if strings.Contains(fileName, ".cpp") {
			if strings.Contains(contents, "parse->Event") {
				for _, l := range strings.Split(contents, "\n") {
					if strings.Contains(l, "parse->Event") && strings.Contains(l, "EVENT_") {
						entity := ""
						event := ""

						// @eg parse->EventNPC(EVENT_KILLED_MERIT, this, c, "killed", 0);
						// @result NPC
						entitySplit := strings.Split(l, "parse->Event")
						if len(entitySplit) > 0 {
							entitySplit2 := strings.Split(entitySplit[1], "(")
							if len(entitySplit2) > 0 {
								entity = strings.TrimSpace(entitySplit2[0])
							}
						}

						// @eg parse->EventNPC(EVENT_KILLED_MERIT, this, c, "killed", 0);
						// @result EVENT_KILLED_MERIT
						eventSplit := strings.Split(l, "(EVENT_")
						if len(eventSplit) > 0 {
							eventSplit2 := strings.Split(eventSplit[1], ",")
							if len(eventSplit2) > 0 {
								event = "EVENT_" + strings.TrimSpace(eventSplit2[0])
							}
						}

						if entity == "Encounter" {
							continue
						}

						eventEntityMappings = append(
							eventEntityMappings, PerlEventEntityMapping{
								Entity: entity,
								Event:  event,
							},
						)
					}
				}
			}
		}
	}

	// second pass; parse event vars and tie to entities
	perlEvents := []PerlEvent{}
	hasMapped := []string{}
	for fileName, _ := range files {
		contents := files[fileName]
		if strings.Contains(fileName, "embparser.cpp") {
			currentEvent := ""
			events := []string{}
			eventVars := []string{}
			for _, l := range strings.Split(contents, "\n") {

				// break will start a new buffer
				// we could use the same exported vars for multiple event cases
				if strings.Contains(l, "break") && currentEvent != "" {

					// loop through events scanned in the switch case since they have similar vars
					for _, event := range events {

						// loop through entity mappings and associate and entity with each event
						// found in the mappings
						for _, mapping := range eventEntityMappings {

							// used the mapping entity and event to make sure we don't add
							// the same event and entity pair more than once
							// we do this because in the first pass we may have parsed multiple of the same
							// event / entity pair because its in the code multiple times conditionally
							mapKey := fmt.Sprintf("%s%s", mapping.Entity, event)

							//fmt.Printf("[%v] [%v]\n", mapping.Event, event)
							if mapping.Event == event && !Find(hasMapped, mapKey) {

								// get the perl script-used event name if exists
								// @example source: EVENT_TRADE perl usage: EVENT_ITEM
								finalEvent := event
								eventIndex := indexOf(event, eventCodes)
								if eventIndex >= 0 {
									if len(perlEventCodes) > eventIndex {
										finalEvent = perlEventCodes[eventIndex]
									}
								}

								perlEvents = append(
									perlEvents, PerlEvent{
										EntityType:      mapping.Entity,
										EventName:       finalEvent,
										EventIdentifier: finalEvent,
										EventVars:       eventVars,
									},
								)

								hasMapped = append(hasMapped, mapKey)
							}
						}
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

	}

	return perlEvents
}
