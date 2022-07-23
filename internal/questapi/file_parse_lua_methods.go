package questapi

import (
	"strings"
)

// parses lua methods
func (c *ParseService) parseLuaMethods(contents string, fileName string, luaMethods map[string][]LuaMethod) {

	// first pass over file to build a list of definitions and map then to the actual functions themselves
	// definitions are not always 1:1 with the actual source function names
	luaDefinitionMap := map[string]string{}
	for _, l := range strings.Split(contents, "\n") {
		if strings.Contains(l, ".def(\"") || strings.Contains(l, "luabind::def(\"") {
			l = strings.TrimSpace(l)
			l = strings.ReplaceAll(l, ".def(\"", "")
			l = strings.ReplaceAll(l, ".luabind::def(\"", "")

			// split to pull out definition name
			// ex: .def("GetHighestLevel", (int(Lua_Raid::*)(void))&Lua_Raid::GetHighestLevel)
			// get: GetHighestLevel
			definitionSplit := strings.Split(l, "\",")
			definition := strings.TrimSpace(definitionSplit[0])

			// split to pull out method name to key off of
			// ex: GetHighestLevel", (int(Lua_Raid::*)(void))&Lua_Raid::GetHighestLevel)
			// get: GetHighestLevel
			methodSplit := strings.Split(l, "&")

			// matches .def("GetHighestLevel", (int(Lua_Raid::*)(void))&Lua_Raid::GetHighestLevel)
			if len(methodSplit) > 1 {
				methodNameSplit := strings.Split(methodSplit[1], "::")
				if len(methodNameSplit) > 1 {

					methodName := strings.ReplaceAll(methodNameSplit[1], ")", "")

					// add to map
					luaDefinitionMap[methodName] = definition

					//fmt.Println(l)
					//fmt.Printf("Definition: %+v\n", definition)
					//fmt.Printf("MethodName: %+v\n", methodName)
					//fmt.Printf("\n")
				}

				// matches lua_genera.cpp
				// luabind::def("world_wide_set_entity_variable_client", (void(*)(const char*,const char*))&lua_world_wide_set_entity_variable_client),
				if len(methodNameSplit) == 0 {

					methodName := methodSplit[1]
					methodName = strings.ReplaceAll(methodName, "),", "")
					methodName = strings.ReplaceAll(methodName, ")", "")

					// add to map
					luaDefinitionMap[methodName] = definition
				}
			}
		}
	}

	for _, l := range strings.Split(contents, "\n") {
		if strings.Contains(l, "const char *") {
			l = strings.ReplaceAll(l, "const char *", "string ")
		}

		if strings.Contains(l, "const char*") {
			l = strings.ReplaceAll(l, "const char*", "string")
		}

		if strings.Contains(l, "const char") {
			l = strings.ReplaceAll(l, "const char", "string")
		}

		// transform 1 off statements like the following
		// luabind::adl::object lua_get_zone_time(lua_State *L) {
		if strings.Contains(l, "{") {
			l = strings.ReplaceAll(l, "luabind::adl::object", "object")
			l = strings.ReplaceAll(l, "State *L, ", "")
			l = strings.ReplaceAll(l, "State *L", "")
		}
		//fmt.Println(l)

		lineSplit := strings.Split(strings.TrimSpace(l), " ")

		// int Lua_Inventory::CalcSlotFromMaterial(int material)
		// here we first grep for lua_
		if len(lineSplit) > 1 && len(lineSplit[1]) > 5 {

			methodIdentifier := strings.ToLower(lineSplit[1][:4])
			if methodIdentifier == "lua_" &&
				!strings.Contains(l, "luabind") &&
				!strings.Contains(l, "return ") &&
				!strings.Contains(l, "#ifdef") &&
				!strings.Contains(l, "struct ") {

				//fmt.Println("")
				//fmt.Println(strings.TrimSpace(l))
				//fmt.Println(methodIdentifier)

				// luabind::adl::object lua_get_zone_time(lua_State *L) {

				// pull off the return type
				// <Lua_Client> Lua_EntityList::GetClientByCharID(uint32 char_id) {
				returnType := strings.TrimSpace(lineSplit[0])
				l = strings.ReplaceAll(l, returnType+" ", "")

				// strip remainders
				// From: Lua_EntityList::GetClientByCharID(uint32 char_id) {
				// To: EntityList::GetClientByCharID(uint32 char_id)
				l = strings.ReplaceAll(l, "Lua_", "")
				l = strings.ReplaceAll(l, " {", "")
				l = strings.ReplaceAll(l, "std::string ", "string ")
				l = strings.ReplaceAll(l, " const", "")

				if strings.Contains(l, "::") {
					split := strings.Split(l, "::")
					methodType := strings.TrimSpace(split[0])
					methodFull := strings.TrimSpace(split[1])

					methodName := ""
					methodNameSplit := strings.Split(methodFull, "(")
					if len(methodNameSplit) > 1 {
						methodName = strings.TrimSpace(methodNameSplit[0])
					}

					//fmt.Printf("Method name is [%v]\n", methodName)

					// Lua_Object Lua_EntityList::GetObjectByDBID(uint32 db_id) {
					//	Object EntityList::GetObjectByDBID(uint32 db_id)
					//	Return Type: Lua_Object
					//	Method Type: Object EntityList

					// check if there is a definition mapping that is different from the
					// method name parsed from the function itself, this is what is
					// actually used in the quest api
					if val, ok := luaDefinitionMap[methodName]; ok {
						//fmt.Printf("Found override map replacing [%v] in [%v] for [%v]\n", methodFull, methodName, val)
						methodFull = strings.ReplaceAll(methodFull, methodName, val)
					}

					// params
					methodFull = strings.ReplaceAll(methodFull, "lua_", "")
					paramSplit := strings.Split(methodFull, "(")
					params := strings.ReplaceAll(paramSplit[1], ")", "")
					method := paramSplit[0]

					// split and zero if empty
					p := strings.Split(strings.TrimSpace(params), ",")
					for i := range p {
						p[i] = strings.TrimSpace(p[i])
					}
					if params == "" {
						p = []string{}
					}

					// add to functions
					luaMethods[methodType] = append(
						luaMethods[methodType], LuaMethod{
							Method:     method,
							Params:     p,
							MethodType: methodType,
							ReturnType: returnType,
							Categories: nil,
						},
					)

					//fmt.Println(strings.TrimSpace(l))
					//fmt.Printf("Return Type: %+v\n", returnType)
					//fmt.Printf("Method: %+v\n", method)
					//fmt.Printf("Method Type: %+v\n", methodType)
					//fmt.Printf("\n")
				}

				// matches: lua_add_expedition_lockout_by_char_id(uint32 char_id, string expedition_name, string event_name, uint32 seconds)
				if strings.Contains(fileName, "lua_general") {
					//fmt.Println(l)

					methodFull := l
					methodType := "eq"
					methodName := ""
					methodNameSplit := strings.Split(methodFull, "(")
					if len(methodNameSplit) > 1 {
						methodName = strings.TrimSpace(methodNameSplit[0])
					}

					// check if there is a definition mapping that is different from the
					// method name parsed from the function itself, this is what is
					// actually used in the quest api
					if val, ok := luaDefinitionMap[methodName]; ok {
						//fmt.Printf("Found override map replacing [%v] in [%v] for [%v]\n", methodFull, methodName, val)
						methodFull = strings.ReplaceAll(methodFull, methodName, val)
					}

					// params
					methodFull = strings.ReplaceAll(methodFull, "lua_", "")
					paramSplit := strings.Split(methodFull, "(")
					params := strings.ReplaceAll(paramSplit[1], ")", "")
					method := paramSplit[0]

					// split and zero if empty
					p := strings.Split(strings.TrimSpace(params), ",")
					for i := range p {
						p[i] = strings.TrimSpace(p[i])
					}
					if params == "" {
						p = []string{}
					}

					// add to functions
					luaMethods[methodType] = append(
						luaMethods[methodType], LuaMethod{
							Method:     method,
							Params:     p,
							MethodType: methodType,
							ReturnType: returnType,
							Categories: nil,
						},
					)
				}

				//fmt.Println(methodIdentifier)
			}
		}
	}
}
