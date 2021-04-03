package questapi

import (
	"strings"
)

type LuaEventHandler struct {
	EventHandler string   `json:"event_handler"` // @eg handle_spell_fade
	EventVars    []string `json:"event_vars"`    // @eg []{"target", "buff_slot", "caster_id"}
}

type LuaEvent struct {
	EntityType      string   `json:"entity_type"`      // @eg Player
	EventName       string   `json:"event_name"`       // @eg EVENT_SAY
	EventIdentifier string   `json:"event_identifier"` // @eg say or event_say(e)
	EventVars       []string `json:"event_vars"`       // @eg []{"text", "langid"}
}

type LuaEventMapping struct {
	ScriptEventIdentifier string // @eg say = event_say(e)
	Event                 string // @eg EVENT_SAY
}

func (c *ParseService) parseLuaEvents(files map[string]string) []LuaEvent {

	// #1 Handler needs to be mapped from lua_parser_events.cpp eg: void handle_npc_popup
	// #2 Parse lua_general.cpp#L2984 to map event identifiers (event_say(e)) to sub events EVENT_SAY
	// #3 Parse lua_parser.cpp#L153 to map sub events to handles
	// #4 For every event not already mapped (Doesn't have a specific handler) add it to the list

	// #1 Handler needs to be mapped from lua_parser_events.cpp eg: void handle_npc_popup
	// build list of handlers
	// @example
	// (questapi.LuaEventHandler) {
	//  EventHandler: (string) (len=17) "handle_spell_fade",
	//  EventVars: ([]string) (len=3 cap=4) {
	//   (string) (len=6) "target",
	//   (string) (len=9) "buff_slot",
	//   (string) (len=9) "caster_id"
	//  }
	luaEventHandlers := []LuaEventHandler{}
	for fileName := range files {
		if strings.Contains(fileName, "lua_parser_events.cpp") {
			currentEvent := ""
			events := []string{}
			eventVars := []string{}
			for _, l := range strings.Split(files[fileName], "\n") {

				// break will start a new buffer
				// we could use the same exported vars for multiple event cases
				if strings.Contains(l, "handle_") && currentEvent != "" {
					for _, event := range events {
						luaEventHandlers = append(
							luaEventHandlers, LuaEventHandler{
								EventHandler: event,
								EventVars:    eventVars,
							},
						)
					}

					// reset
					events = []string{}
					eventVars = []string{}
				}

				// keep track of the current event
				if strings.Contains(l, "handle_") {
					line := l

					handleSplit := strings.Split(line, "handle_")
					if len(handleSplit) > 0 {

						// the handle needs to be mapped
						// parse lua_general.cpp#L2984 first to get event names to sub events
						// Then parse lua_parser.cpp#L153 to map sub events to handles
						eventSplit := strings.Split(handleSplit[1], "(")
						event := "handle_" + eventSplit[0]
						currentEvent = event
					}

					// if event set, append
					if currentEvent != "" {
						events = append(events, currentEvent)
					}
				}

				// parse vars
				if strings.Contains(l, "lua_setfield") && currentEvent != "" {
					quoteSplit := strings.Split(l, "\"")
					if len(quoteSplit) > 1 {
						e := strings.TrimSpace(quoteSplit[1])
						if !Find(eventVars, e) {
							eventVars = append(eventVars, e)
						}
					}
				}
			}
		}
	}

	// #2 Parse lua_general.cpp#L2984 to map event identifiers (event_say(e)) to sub events EVENT_SAY
	//
	// ([]questapi.LuaEventMapping) (len=78 cap=128) {
	// (questapi.LuaEventMapping) {
	//  ScriptEventIdentifier: (string) (len=3) "say",
	//  Event: (string) (len=9) "EVENT_SAY"
	// },
	// (questapi.LuaEventMapping) {
	//  ScriptEventIdentifier: (string) (len=5) "trade",
	//  Event: (string) (len=11) "EVENT_TRADE"
	// },
	luaEventMappings := []LuaEventMapping{}
	for fileName := range files {
		if strings.Contains(fileName, "lua_general.cpp") {
			for _, l := range strings.Split(files[fileName], "\n") {
				if strings.Contains(l, "luabind::value") && strings.Contains(l, "EVENT_") {
					identifier := ""
					event := ""

					// @sample luabind::value("spell_effect", static_cast<int>(EVENT_SPELL_EFFECT_CLIENT)),
					// @grabs spell_effect
					quoteSplit := strings.Split(l, "\"")
					if len(quoteSplit) > 0 {
						identifier = strings.TrimSpace(quoteSplit[1])
					}

					// @sample luabind::value("spell_effect", static_cast<int>(EVENT_SPELL_EFFECT_CLIENT)),
					// @grabs EVENT_SPELL_EFFECT_CLIENT
					eventSplit := strings.Split(l, ">(")
					if len(eventSplit) > 0 {
						event = eventSplit[1]
						event = strings.ReplaceAll(event, ")", "")
						event = strings.ReplaceAll(event, ",", "")
					}

					// @identifier spawn
					// @event EVENT_SPAWN
					luaEventMappings = append(
						luaEventMappings, LuaEventMapping{
							ScriptEventIdentifier: identifier,
							Event:                 event,
						},
					)
				}
			}
		}
	}

	// #3 Parse lua_parser.cpp#L153 to map sub events to handles
	luaEvents := []LuaEvent{}
	for fileName := range files {
		if strings.Contains(fileName, "lua_parser.cpp") {
			for _, l := range strings.Split(files[fileName], "\n") {
				// @grep ItemArgumentDispatch[EVENT_AUGMENT_ITEM] = handle_item_augment;
				if strings.Contains(l, "ArgumentDispatch[EVENT_") {

					// vars
					entity := ""
					event := ""
					handler := ""

					// @example ItemArgumentDispatch[EVENT_WEAPON_PROC] = handle_item_proc;
					// @result Item
					dispatchSplit := strings.Split(l, "ArgumentDispatch")
					if len(dispatchSplit) > 0 {
						entity = strings.TrimSpace(dispatchSplit[0])
					}

					// @example ItemArgumentDispatch[EVENT_WEAPON_PROC] = handle_item_proc;
					// @result EVENT_WEAPON_PROC
					eventSplit := strings.Split(l, "Dispatch[")
					if len(eventSplit) > 0 {
						eventSplit2 := strings.Split(eventSplit[1], "]")
						if len(eventSplit2) > 0 {
							event = eventSplit2[0]
						}
					}

					// @example ItemArgumentDispatch[EVENT_WEAPON_PROC] = handle_item_proc;
					// @result handle_item_proc
					handleSplit := strings.Split(l, " =")
					if len(handleSplit) > 0 {
						handler = strings.TrimSpace(handleSplit[1])
						handler = strings.ReplaceAll(handler, ";", "")
					}

					// map everything together here to create master list of events

					// get sub event identifiers
					identifier := ""
					for _, mapping := range luaEventMappings {
						if mapping.Event == event {
							identifier = mapping.ScriptEventIdentifier
						}
					}

					// map handler to exported event vars
					eventVars := []string{}
					for _, eventHandler := range luaEventHandlers {
						if eventHandler.EventHandler == handler {
							eventVars = eventHandler.EventVars
						}
					}

					e := LuaEvent{
						EntityType:      entity,
						EventName:       event,
						EventIdentifier: "event_" + identifier,
						EventVars:       eventVars,
					}

					luaEvents = append(luaEvents, e)
				}
			}
		}
	}

	// #4 For every event not already mapped (Doesn't have a specific handler) add it to the list
	for fileName := range files {
		content := files[fileName]
		// grep: parse->EventPlayer(EVENT_LEVEL_UP, this, "", 0);
		if strings.Contains(content, "parse->Event") {
			for _, l := range strings.Split(files[fileName], "\n") {
				if strings.Contains(l, "parse->Event") {
					parseSplit := strings.Split(l, "parse->Event")
					if len(parseSplit) > 0 {
						eventLine := parseSplit[1]
						eventTypeSplit := strings.Split(eventLine, "(")
						// grab event type and then grab the sub event
						// grep: Player(EVENT_LEVEL_UP, this, "", 0);
						if len(eventTypeSplit) > 0 {
							eventType := eventTypeSplit[0]
							eventArgs := eventTypeSplit[1]
							eventArgsSplit := strings.Split(eventArgs, ",")
							// grep: EVENT_LEVEL_UP, this, "", 0);
							// grab EVENT_LEVEL_UP
							if len(eventArgsSplit) > 0 {
								event := strings.TrimSpace(eventArgsSplit[0])

								// make sure the event doesn't already exist
								eventExists := false
								for _, luaEvent := range luaEvents {
									if luaEvent.EventName == event && luaEvent.EntityType == eventType {
										eventExists = true
										break
									}
								}

								// if event doesn't exist we need to map the EVENT_NAME to a script event_name(e)
								// since they are not usually 1:1
								if !eventExists {
									identifier := ""
									for _, mapping := range luaEventMappings {
										if mapping.Event == event {
											identifier = mapping.ScriptEventIdentifier
										}
									}

									if len(identifier) > 0 {
										e := LuaEvent{
											EntityType:      eventType,
											EventName:       event,
											EventIdentifier: "event_" + identifier,
											EventVars:       []string{},
										}

										luaEvents = append(luaEvents, e)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return luaEvents
}
