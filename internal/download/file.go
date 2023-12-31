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
	banner.Loading()

	var attemptStr string
	if attempt > 1 {
		attemptStr = fmt.Sprintf("| Attempt (%v of %v)", attempt, maxRetries)
	}

	fmt.Printf("[Download]  URL | %v %v\n", downloadUrl, attemptStr)
	fmt.Printf("[Download] File | %v\n", destinationPath)

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
		progressbar.OptionSetDescription("[Download]"),
		progressbar.OptionSetWriter(os.Stderr),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(50),
		progressbar.OptionThrottle(10*time.Millisecond),
		progressbar.OptionOnCompletion(func() {
			fmt.Fprint(os.Stderr)
		}),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "| ",
			BarEnd:        " |",
		}),
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

	fmt.Printf("\n\n")

	return nil
}
