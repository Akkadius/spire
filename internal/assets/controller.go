package assets

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/anaskhan96/soup"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"strings"
)

type Controller struct {
	db *database.Resolver
}

func NewController(
	db *database.Resolver,
) *Controller {
	return &Controller{
		db: db,
	}
}

func (f *Controller) Routes() []*routes.Route {
	return []*routes.Route{
		routes.RegisterRoute(http.MethodGet, "assets/zone-images/:zone_long_name", f.zoneImages, nil),
	}
}

func (f *Controller) zoneImages(c echo.Context) error {
	zone, err := url.PathUnescape(c.Param("zone_long_name"))
	if err != nil {
		fmt.Println(err)
	}

	resp, err := soup.Get("https://everquest.allakhazam.com/db/zone.html?ztype=0")
	if err != nil {
		fmt.Println(err)
	}

	// bit of a hack, but helps match some zone names
	zone = strings.ReplaceAll(zone, "The ", "")
	zone = strings.ReplaceAll(zone, "City of ", "")

	fmt.Println(zone)

	images := []string{}
	doc := soup.HTMLParse(resp)
	for _, link := range doc.FindAll("a") {
		if strings.Contains(link.Attrs()["href"], "zstrat") && strings.Contains(link.Text(), zone) {
			fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])

			// grab zone level
			r, err := soup.Get("https://everquest.allakhazam.com" + link.Attrs()["href"])
			if err != nil {
				fmt.Println(err)
			}

			d := soup.HTMLParse(r)
			for _, l := range d.FindAll("a") {
				if strings.Contains(l.Attrs()["href"], "scenery") {
					images = append(images, l.Attrs()["href"])
				}
			}
		}
	}

	return c.JSON(http.StatusOK, echo.Map{"images": images})
}
