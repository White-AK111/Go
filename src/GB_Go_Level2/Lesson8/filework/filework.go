package filework

import (
	"encoding/hex"
	"fmt"
	"github.com/White-AK111/GB_Go_Level2/Lesson8/config"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type FileEntity struct {
	Create time.Time
	Name   string
	Path   string
	Hash   string
	Size   int64
}

// getHashOfFile method get hash of file
func (f *FileEntity) getHashOfFile(app config.App) error {

	file, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer file.Close()

	app.HashAlgorithm.Reset()
	if _, err := io.Copy(app.HashAlgorithm, file); err != nil {
		return err
	}

	hashInBytes := app.HashAlgorithm.Sum(nil)
	f.Hash = hex.EncodeToString(hashInBytes)

	return nil
}

func (f *FileEntity) contains(fl []FileEntity, it int) bool {
	for i := it + 1; i < len(fl); i++ {
		if fl[i].Name == f.Name && fl[i].Size == f.Size && fl[i].Hash == f.Hash {
			return true
		}
	}
	return false
}

func NewFileEntity() *FileEntity {
	return &FileEntity{
		Create: time.Now(),
		Name:   "",
		Path:   "",
		Hash:   "",
		Size:   0,
	}
}

func DoDuplicateFiles(app config.App) error {

	fileList, err := findAllFiles(app)
	if err != nil {
		log.Printf("Error on find all files in source path: %s", err)
	}

	sort.Slice(fileList, func(i, j int) bool {
		return fileList[i].Path < fileList[j].Path
	})

	for i, file := range fileList {
		if file.contains(fileList, i) {
			fmt.Printf("Duplicate file: %v\n", file)
		}
	}

	if app.FlagDelete {
		var confirm string
		for strings.ToUpper(confirm) != "Y" && strings.ToUpper(confirm) != "N" {
			fmt.Print("Delete this files? (Y/N): ")
			fmt.Fscan(os.Stdin, &confirm)
		}
		if strings.ToUpper(confirm) == "Y" {
			fmt.Println("Files deleted!")
		}
	}

	return nil
}

func DoRandomCopyFiles(app config.App) error {

	fileList, err := findAllFiles(app)
	if err != nil {
		log.Printf("Error on find all files in source path: %s", err)
	}

	sort.Slice(fileList, func(i, j int) bool {
		return fileList[i].Path < fileList[j].Path
	})

	if app.FlagRandCopy {
		var confirm string
		for strings.ToUpper(confirm) != "Y" && strings.ToUpper(confirm) != "N" {
			fmt.Print("Do random copy files? (Y/N): ")
			fmt.Fscan(os.Stdin, &confirm)
		}
		if strings.ToUpper(confirm) == "Y" {
			fmt.Println("Files copied!")
		}
	}

	return nil
}

func findAllFiles(app config.App) (fileList []FileEntity, err error) {

	// get absolut filepath for source path
	sourcePath, err := filepath.Abs(app.SourcePath)
	if err != nil {
		fmt.Printf("error on get ABS path from source path %q: %v\n", sourcePath, err)
		return fileList, err
	}
	app.SourcePath = sourcePath

	// recursive walk in directory and all subdirectory, find only files
	err = filepath.WalkDir(sourcePath, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			f := NewFileEntity()
			f.Name = info.Name()
			f.Path = path
			fi, err := info.Info()
			if err != nil {
				log.Printf("can't get FileInfo of file %q: %v\n", path, err)
				return err
			}
			f.Create = fi.ModTime()
			if err = f.getHashOfFile(app); err != nil {
				log.Printf("can't get hash of file %q: %v\n", path, err)
				return err
			}
			f.Size = fi.Size()
			fileList = append(fileList, *f)
		}
		return nil
	})

	if err != nil {
		log.Printf("error walking the path %q: %v\n", sourcePath, err)
		return fileList, err
	}

	return fileList, err
}
