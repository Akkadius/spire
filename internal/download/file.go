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

const maxRetries = 3

func WithProgress(destinationPath, downloadUrl string) error {
	attempt := 1
	for {
		err := download(destinationPath, downloadUrl, attempt)
		if err != nil {
			if attempt < maxRetries {
				attempt++
				continue
			}
			return err
		}
		break
	}

	return nil
}

func download(destinationPath, downloadUrl string, attempt int) error {
	fmt.Printf("-------------------------------------------------------------------------------------------\n")
	banner.Loading()
	fmt.Printf("-------------------------------------------------------------------------------------------\n")

	fmt.Printf("[Downloading]  URL | %v | Attempt (%v of %v)\n", downloadUrl, attempt, maxRetries)
	fmt.Printf("[Downloading] File | %v\n", destinationPath)

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
		progressbar.OptionThrottle(10*time.Millisecond),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr)
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

	err = f.Sync()
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	// check if file destinationPath exists
	if _, err := os.Stat(destinationPath); os.IsNotExist(err) {
		return fmt.Errorf("could not download file [%v]", destinationPath)
	}

	fmt.Printf("\n-------------------------------------------------------------------------------------------")
	fmt.Printf("\n")

	return nil
}
