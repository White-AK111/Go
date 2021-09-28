package millionfiles

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// createFile create file in directory
func createFile(i int, dir string) {
	ext := ".txt"
	fileName := strconv.Itoa(i) + ext
	f, err := os.Create(filepath.Join(dir, fileName))
	if err != nil {
		log.Fatalf("Error on create file: %s", err)
	}
	defer f.Close()
	fmt.Printf("File created:%s\n", fileName)
}

// CreateMillionFiles create 1M files in directory
func CreateMillionFiles(dir string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered", r)
			err = errors.New("panic on create files")
		}
	}()

	err = removeContents(dir)
	if err != nil {
		log.Fatalf("Error on delete files: %s", err)
	}

	for i := 1; i <= 1000000; i++ {
		createFile(i, dir)
	}

	return err
}

// removeContents remove all files from directory
func removeContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
