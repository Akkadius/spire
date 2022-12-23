package download

import (
	"fmt"
	"github.com/Akkadius/spire/internal/banner"
	"github.com/schollz/progressbar/v3"
	"io"
	"net/http"
	"os"
	"time"
)

func WithProgress(destinationPath, downloadUrl string) error {
	banner.Loading()

	fmt.Printf("[Downloading] URL [%v]\n", downloadUrl)
	fmt.Printf("[Downloading] To [%v]\n\n", destinationPath)

	tempDestinationPath := destinationPath + ".tmp"
	req, _ := http.NewRequest("GET", downloadUrl, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	f, _ := os.OpenFile(tempDestinationPath, os.O_CREATE|os.O_WRONLY, 0644)

	bar := progressbar.NewOptions64(
		resp.ContentLength,
		progressbar.OptionSetDescription("[Downloading]"),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(30),
		progressbar.OptionThrottle(100*time.Millisecond),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr, "\n")
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionSetRenderBlankState(true),
	)

	io.Copy(io.MultiWriter(f, bar), resp.Body)
	os.Rename(tempDestinationPath, destinationPath)

	fmt.Println("")

	return nil
}
