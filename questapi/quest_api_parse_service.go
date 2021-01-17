package questapi

import (
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"sort"
	"strings"
	"time"
)

type ParseService struct {
	logger *logrus.Logger
	cache  *gocache.Cache
}

func NewParseService(logger *logrus.Logger, cache *gocache.Cache) *ParseService {
	return &ParseService{logger: logger, cache: cache}
}

type QuestApiMethods struct {
	PerlApi map[string][]PerlMethod `json:"perl_api"`
	LuaApi  map[string][]LuaMethod  `json:"lua_api"`
}

type PerlMethod struct {
	Method     string   `json:"method"`
	MethodType string   `json:"method_type"`
	ReturnType string   `json:"return_type"`
	Categories []string `json:"categories"`
}

type LuaMethod struct {
	Method     string   `json:"method"`
	MethodType string   `json:"method_type"`
	ReturnType string   `json:"return_type"`
	Categories []string `json:"categories"`
}

var perlMethods = map[string][]PerlMethod{}
var luaMethods = map[string][]LuaMethod{}

const lockKey = "quest-api-lock"

// parses methods from our source
func (c *ParseService) Parse(forceRefresh bool) QuestApiMethods {

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

	fs := memfs.New()
	_, err := git.Clone(
		memory.NewStorage(), fs, &git.CloneOptions{
			URL:               "https://github.com/EQEmu/Server.git",
			ReferenceName:     "refs/heads/master",
			RecurseSubmodules: 0,
			Depth:             1,
		},
	)
	if err != nil {
		c.logger.Fatal(err)
	}

	zonePath := "./zone"

	// empty maps
	for k := range perlMethods {
		delete(perlMethods, k)
	}
	for k := range luaMethods {
		delete(luaMethods, k)
	}

	// read through memory file system
	files, err := fs.ReadDir(zonePath)
	if err != nil {
		c.logger.Fatal(err)
	}

	// loop through files
	for _, f := range files {

		// perl files
		isPerlFile := strings.Contains(f.Name(), "perl_") ||
			strings.Contains(f.Name(), "embparser_api")

		if isPerlFile {
			parsePerlMethods(fs, f.Name(), perlMethods)
		}

		// lua files
		isLuaFile := strings.Contains(f.Name(), "lua_") &&
			strings.Contains(f.Name(), "cpp")

		if isLuaFile {
			parseLuaMethods(fs, f.Name(), luaMethods)
		}
	}

	// sort
	for _, methods := range perlMethods {
		sort.Slice(
			methods[:], func(i, j int) bool {
				return methods[i].Method < methods[j].Method
			},
		)
	}
	for _, methods := range luaMethods {
		sort.Slice(
			methods[:], func(i, j int) bool {
				return methods[i].Method < methods[j].Method
			},
		)
	}

	// delete lock
	c.cache.Delete(lockKey)

	return c.apiResponse()
}

func (c *ParseService) apiResponse() QuestApiMethods {
	return QuestApiMethods{
		PerlApi: perlMethods,
		LuaApi:  luaMethods,
	}
}
