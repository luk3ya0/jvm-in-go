package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer func(r *zip.ReadCloser) {
		err := r.Close()
		if err != nil {
			log.Fatalln("closing error rises in reading jar file")
		}
	}(r)

	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			// ignore possible resources leak,
			// due to conditional flow and return clause
			defer func(rc io.ReadCloser) {
				_ = rc.Close()
			}(rc)
			// TODO: using other method to read class file
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
