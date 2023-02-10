package unzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Unzip struct {
	Src   string
	Dest  string
	debug bool
}

func New(src string, dest string) Unzip {
	return Unzip{Src: src, Dest: dest, debug: false}
}

func (uz Unzip) Extract() error {

	fmt.Println("[Zip] Extraction of [" + uz.Src + "] started!")

	r, err := zip.OpenReader(uz.Src)
	if err != nil {
		return err
	}

	err = os.MkdirAll(uz.Dest, 0755)
	if err != nil {
		return err
	}

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}

		path := filepath.Join(uz.Dest, f.Name)
		if !strings.HasPrefix(path, filepath.Clean(uz.Dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: Illegal file path", path)
		}

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(path, f.Mode())
			if err != nil {
				return err
			}
		} else {
			err := os.MkdirAll(filepath.Dir(path), f.Mode())
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
		if uz.debug {
			fmt.Println("Extracting file: " + f.Name)
		}
	}

	err = r.Close()
	if err != nil {
		return err
	}

	fmt.Printf("[Zip] Extracted (%v) files in [%v] to [%v]!\n\n", len(r.File), uz.Src, uz.Dest)

	return nil
}
