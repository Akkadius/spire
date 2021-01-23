package questapi

import (
	"github.com/Akkadius/spire/internal/github"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
	"time"
)

type ParseService struct {
	logger     *logrus.Logger
	cache      *gocache.Cache
	downloader *github.GithubSourceDownloader
	files      map[string]string
}

func NewParseService(
	logger *logrus.Logger,
	cache *gocache.Cache,
	downloader *github.GithubSourceDownloader,
) *ParseService {
	return &ParseService{logger: logger, cache: cache, downloader: downloader}
}

type QuestApiResponse struct {
	LastRefreshed time.Time `json:"last_refreshed"`

	PerlApi `json:"perl_api"`
	LuaApi  `json:"lua_api"`
}

type PerlApi struct {
	PerlMethods   map[string][]PerlMethod `json:"methods"`
	PerlEvents    []PerlEvent             `json:"events"`
	PerlConstants []PerlConstants         `json:"constants"`
}

type LuaApi struct {
	LuaMethods   map[string][]LuaMethod `json:"methods"`
	LuaEvents    []LuaEvent             `json:"events"`
	LuaConstants []LuaConstants         `json:"constants"`
}

type PerlMethod struct {
	Method     string   `json:"method"`
	Params     []string `json:"params"`
	MethodType string   `json:"method_type"`
	ReturnType string   `json:"return_type"`
	MethodFull string   `json:"method_full"`
	Categories []string `json:"categories"`
}

type LuaMethod struct {
	Method     string   `json:"method"`
	Params     []string `json:"params"`
	MethodType string   `json:"method_type"`
	ReturnType string   `json:"return_type"`
	MethodFull string   `json:"method_full"`
	Categories []string `json:"categories"`
}

var perlMethods = map[string][]PerlMethod{}
var luaMethods = map[string][]LuaMethod{}
var lastRefreshed time.Time
var perlEvents []PerlEvent
var luaEvents []LuaEvent
var perlConstants []PerlConstants
var luaConstants []LuaConstants

const lockKey = "quest-api-lock"

// return files from memory
func (c *ParseService) Files() map[string]string {
	return c.files
}

// source files
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
	c.files = c.downloader.Source(org, repo, branch, forceRefresh)

	// return local reference
	return c.files
}

// parses methods from our source
func (c *ParseService) Parse(forceRefresh bool) QuestApiResponse {

	// pull files in
	c.Source("EQEmu", "Server", "master", forceRefresh)

	// return cached copy
	if len(perlMethods) > 0 && len(luaMethods) > 0 && !forceRefresh {
		return c.apiResponse()
	}

	// if lock set, return
	_, found := c.cache.Get(lockKey)
	if found {
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
	sort.Slice(
		perlConstants, func(i, j int) bool {
			if perlConstants[i].ConstantType != perlConstants[j].ConstantType {
				return perlConstants[i].ConstantType < perlConstants[j].ConstantType
			}

			return perlConstants[i].Constant < perlConstants[j].Constant
		},
	)

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
	sort.Slice(
		luaConstants, func(i, j int) bool {
			if luaConstants[i].ConstantType != luaConstants[j].ConstantType {
				return luaConstants[i].ConstantType < luaConstants[j].ConstantType
			}

			return luaConstants[i].Constant < luaConstants[j].Constant
		},
	)

	lastRefreshed = time.Now()

	// delete lock
	c.cache.Delete(lockKey)

	return c.apiResponse()
}

// response method
func (c *ParseService) apiResponse() QuestApiResponse {
	return QuestApiResponse{
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
