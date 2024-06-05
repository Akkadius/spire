package unzip

import (
	"archive/zip"
	"fmt"
	"github.com/Akkadius/spire/internal/logger"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Unzipper struct {
	logger *logger.AppLogger
}

func NewUnzipper(logger *logger.AppLogger) *Unzipper {
	return &Unzipper{
		logger: logger,
	}
}

func (uz *Unzipper) Extract(src string, dest string) error {
	uz.logger.Info().Str("src", src).Str("dest", dest).Msg("Extracting zip file")

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		return err
	}

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		path := filepath.Join(dest, f.Name)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: Illegal file path", path)
		}

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
			if err != nil {
				return err
			}
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}

			if err := f.Close(); err != nil {
				if err != nil {
					return err
				}
			}
		}

		if err := rc.Close(); err != nil {
			return err
		}

		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}

		uz.logger.Debug().Str("file", f.Name).Msg("Extracting file")
	}

	err = r.Close()
	if err != nil {
		return err
	}

	uz.logger.Info().Any("files count", len(r.File)).Str("src", src).Str("dest", dest).Msg("Extracted zip file")

	return nil
}
