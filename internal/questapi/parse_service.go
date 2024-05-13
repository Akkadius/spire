package questapi

import (
	"github.com/Akkadius/spire/internal/github"
	gocache "github.com/patrickmn/go-cache"
	"sort"
	"strings"
	"time"
)

type ParseService struct {
	cache      *gocache.Cache
	downloader *github.SourceDownloader
	files      map[string]string
	snippets   map[string]string
}

func NewParseService(
	cache *gocache.Cache,
	downloader *github.SourceDownloader,
) *ParseService {
	return &ParseService{cache: cache, downloader: downloader}
}

type Response struct {
	LastRefreshed time.Time `json:"last_refreshed"`
	PerlApi       `json:"perl"`
	LuaApi        `json:"lua"`
}

type PerlApi struct {
	PerlMethods   map[string][]PerlMethod    `json:"methods"`
	PerlEvents    []PerlEvent                `json:"events"`
	PerlConstants map[string][]PerlConstants `json:"constants"`
}

type LuaApi struct {
	LuaMethods   map[string][]LuaMethod    `json:"methods"`
	LuaEvents    []LuaEvent                `json:"events"`
	LuaConstants map[string][]LuaConstants `json:"constants"`
}

type PerlMethod struct {
	Method     string   `json:"method"`
	Params     []string `json:"params"`
	MethodType string   `json:"method_type"`
	ReturnType string   `json:"return_type"`
	Categories []string `json:"categories"`
}

type LuaMethod struct {
	Method     string   `json:"method"`
	Params     []string `json:"params"`
	MethodType string   `json:"method_type"`
	ReturnType string   `json:"return_type"`
	Categories []string `json:"categories"`
}

var perlMethods = map[string][]PerlMethod{}
var luaMethods = map[string][]LuaMethod{}
var lastRefreshed time.Time
var perlEvents []PerlEvent
var luaEvents []LuaEvent
var perlConstants map[string][]PerlConstants
var luaConstants map[string][]LuaConstants

const lockKey = "quest-api-lock"

// Files return files from memory
func (c *ParseService) Files() map[string]string {
	return c.files
}

// Source source files
func (c *ParseService) Source(
	org string,
	repo string,
	branch string,
	forceRefresh bool,
) map[string]string {

	// return cache if not refresh
	if !forceRefresh && len(c.files) > 0 {
		return c.files
	}

	// set to memory
	c.files = c.downloader.Source(org, repo, branch, forceRefresh).Files

	// return local reference
	return c.files
}

// Source source files
func (c *ParseService) SourceSnippets(
	org string,
	repo string,
	branch string,
	forceRefresh bool,
) map[string]string {

	// return cache if not refresh
	if !forceRefresh && len(c.snippets) > 0 {
		return c.snippets
	}

	// set to memory
	c.snippets = c.downloader.Source(org, repo, branch, forceRefresh).Files

	// return local reference
	return c.snippets
}

// Parse parses methods from our source
func (c *ParseService) Parse(forceRefresh bool) Response {
	// if lock set, return
	_, found := c.cache.Get(lockKey)
	if found {
		return c.apiResponse()
	}

	// pull files in
	c.Source("EQEmu", "Server", "master", forceRefresh)
	c.SourceSnippets("EQEmu", "spire-quest-snippets", "main", forceRefresh)

	// return cached copy
	if len(perlMethods) > 0 && len(luaMethods) > 0 && !forceRefresh {
		return c.apiResponse()
	}

	// set operation lock
	c.cache.Set(lockKey, 1, time.Minute*10)

	// empty maps
	for k := range perlMethods {
		delete(perlMethods, k)
	}
	for k := range luaMethods {
		delete(luaMethods, k)
	}

	// reset
	perlEvents = []PerlEvent{}

	// loop through files
	for fileName, contents := range c.Files() {

		// perl files
		isPerlFile := strings.Contains(fileName, "perl_") || strings.Contains(fileName, "embparser")
		if isPerlFile {
			c.parsePerlMethods(contents, perlMethods)
		}

		// lua files
		isLuaFile := strings.Contains(fileName, "lua_") && strings.Contains(fileName, "cpp")
		if isLuaFile {
			c.parseLuaMethods(contents, fileName, luaMethods)
		}
	}

	// events
	perlEvents = c.parsePerlEvents(c.Files())
	luaEvents = c.parseLuaEvents(c.Files())

	// constants
	perlConstants = c.parsePerlConstants(c.Files())
	luaConstants = c.parseLuaConstants(c.Files())

	// sort perl methods
	for _, methods := range perlMethods {
		sort.Slice(
			methods[:], func(i, j int) bool {
				return methods[i].Method < methods[j].Method
			},
		)
	}
	// sort perl events
	sort.Slice(
		perlEvents, func(i, j int) bool {
			if perlEvents[i].EntityType != perlEvents[j].EntityType {
				return perlEvents[i].EntityType < perlEvents[j].EntityType
			}

			return perlEvents[i].EventName < perlEvents[j].EventName
		},
	)
	// sort perl constants
	for _, constants := range perlConstants {
		sort.Slice(
			constants[:], func(i, j int) bool {
				return constants[i].Constant < constants[j].Constant
			},
		)
	}

	// sort lua methods
	for _, methods := range luaMethods {
		sort.Slice(
			methods[:], func(i, j int) bool {
				return methods[i].Method < methods[j].Method
			},
		)
	}

	// sort lua events
	sort.Slice(
		luaEvents, func(i, j int) bool {
			if luaEvents[i].EntityType != luaEvents[j].EntityType {
				return luaEvents[i].EntityType < luaEvents[j].EntityType
			}

			return luaEvents[i].EventName < luaEvents[j].EventName
		},
	)

	// sort lua constants
	for _, constants := range luaConstants {
		sort.Slice(
			constants[:], func(i, j int) bool {
				return constants[i].Constant < constants[j].Constant
			},
		)
	}

	lastRefreshed = time.Now()

	// delete lock
	c.cache.Delete(lockKey)

	return c.apiResponse()
}

// response method
func (c *ParseService) apiResponse() Response {
	return Response{
		LastRefreshed: lastRefreshed,
		PerlApi: PerlApi{
			PerlMethods:   perlMethods,
			PerlEvents:    perlEvents,
			PerlConstants: perlConstants,
		},
		LuaApi: LuaApi{
			LuaMethods:   luaMethods,
			LuaEvents:    luaEvents,
			LuaConstants: luaConstants,
		},
	}
}

func (c *ParseService) GetSnippets() map[string]string {
	if len(c.snippets) == 0 {
		c.SourceSnippets("EQEmu", "spire-quest-snippets", "main", true)
	}

	return c.snippets
}
