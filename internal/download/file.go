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

	//tempDestinationPath := destinationPath + ".tmp"
	req, err := http.NewRequest("GET", downloadUrl, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// force remove file before creating if exists
	if _, err := os.Stat(destinationPath); err == nil {
		err := os.Remove(destinationPath)
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(destinationPath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}

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

	_, err = io.Copy(io.MultiWriter(f, bar), resp.Body)
	if err != nil {
		return err
	}

	err = resp.Body.Close()
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	//err = os.Rename(tempDestinationPath, destinationPath)
	//if err != nil {
	//	return err
	//}

	fmt.Printf("\n")

	return nil
}
