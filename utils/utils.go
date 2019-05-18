package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func HasFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func GenerateAssetHash(path string) (string, error) {
	f, err := os.Open("./public/" + path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	hash := fmt.Sprintf("%x", h.Sum(nil))
	fileExtPos := strings.LastIndex(path, ".")
	pathWithHash := path[0:fileExtPos] + "-" + hash + path[fileExtPos:]

	return pathWithHash, nil
}

func CopyFileWithHash(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return err
	}

	hash := fmt.Sprintf("%x", h.Sum(nil))
	//TODO: naieve, will break if file name has no . and other upper directory has .
	fileExtPos := strings.LastIndex(path, ".")
	pathWithHash := path[0:fileExtPos] + "-" + hash + path[fileExtPos:]

	return CopyFile(path, pathWithHash)
}

// Copies file source to destination dest.
func CopyFile(source string, dest string) (err error) {
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(source)
		if err == nil {
			err = os.Chmod(dest, si.Mode())
		}

	}

	return
}

// Recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
func CopyDir(source string, dest string) (err error) {
	// get properties of source dir
	fi, err := os.Stat(source)
	if err != nil {
		log.Println("Error", err)
		return err
	}

	if !fi.IsDir() {
		return &CustomError{"Source is not a directory"}
	}

	// ensure dest dir does not already exist

	_, err = os.Open(dest)
	if os.IsNotExist(err) {
		err = os.MkdirAll(dest, fi.Mode())
		if err != nil {
			return err
		}
	}

	entries, err := ioutil.ReadDir(source)
	if err != nil {
		log.Fatal("Cannot copy", err)
	}

	for _, entry := range entries {

		sfp := source + "/" + entry.Name()
		dfp := dest + "/" + entry.Name()
		if entry.IsDir() {
			err = CopyDir(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		}

	}
	return
}

// A struct for returning custom error messages
type CustomError struct {
	What string
}

// Returns the error message defined in What as a string
func (e *CustomError) Error() string {
	return e.What
}
