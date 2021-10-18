package filework

import (
	"fmt"
	"github.com/White-AK111/GB_Go_Level2/Lesson8/config"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

func DoDuplicateFiles(app *config.App) error {
	fInfo := filesInfo{}
	fInfo.directoryList = append(fInfo.directoryList, app.SourcePath)

	err := findAllFiles(app, &fInfo)
	if err != nil {
		app.ErrorLogger.Printf("error on find all files in source path: %s", err)
		return err
	}

	sort.Slice(fInfo.allFilesList, func(i, j int) bool {
		return fInfo.allFilesList[i].Path < fInfo.allFilesList[j].Path
	})

	for i, file := range fInfo.allFilesList {
		if file.contains(fInfo.allFilesList, i) {
			fmt.Printf("Duplicate file: %s	Original file: %s\n", file.Path, file.OriginalFile.Path)
			fInfo.duplicateFilesList = append(fInfo.duplicateFilesList, file)
		}
	}

	fmt.Printf("Total files: %d\n", len(fInfo.allFilesList))
	fmt.Printf("Duplicate files (without original file): %d\n", len(fInfo.duplicateFilesList))

	if app.FlagDelete {
		if len(fInfo.duplicateFilesList) == 0 {
			fmt.Println("No files for delete!")
		} else {
			var confirm string
			for strings.ToUpper(confirm) != "Y" && strings.ToUpper(confirm) != "N" {
				fmt.Print("Delete this duplicate files? (Y/N): ")
				_, err = fmt.Fscan(os.Stdin, &confirm)
				if err != nil {
					app.ErrorLogger.Printf("error on get approval from console: %s\n", err)
					return err
				}
			}
			if strings.ToUpper(confirm) == "Y" {
				err = deleteFiles(app, &fInfo)
				if err != nil {
					app.ErrorLogger.Printf("error on delete files: %s\n", err)
					return err
				}
				fmt.Println("Files deleted!")
			}
		}
	}

	return nil
}

func deleteFiles(app *config.App, fInfo *filesInfo) error {
	wp := newWorkerPool(app.CountGoroutine)
	defer wp.wg.Wait()

	for _, file := range fInfo.duplicateFilesList {
		wp.wg.Add(1)
		go func(file FileEntity) {
			defer func() {
				wp.mu.Unlock()
				// read to release a slot
				<-wp.semaphoreChan
				fInfo.deleteFilesList = append(fInfo.deleteFilesList, file)
				wp.wg.Done()
			}()
			// block while full
			wp.semaphoreChan <- struct{}{}
			wp.mu.Lock()
			if err := os.Remove(file.Path); err != nil {
				app.ErrorLogger.Printf("error on delete file %s: %s\n", file.Path, err)
			}
		}(file)
	}

	return nil
}

func DoRandomCopyFiles(app *config.App) error {
	fInfo := filesInfo{}
	fInfo.directoryList = append(fInfo.directoryList, app.SourcePath)

	err := findAllFiles(app, &fInfo)
	if err != nil {
		app.ErrorLogger.Printf("error on find all files in source path: %s", err)
		return err
	}

	sort.Slice(fInfo.allFilesList, func(i, j int) bool {
		return fInfo.allFilesList[i].Path < fInfo.allFilesList[j].Path
	})

	sort.Slice(fInfo.directoryList, func(i, j int) bool {
		return fInfo.directoryList[i] < fInfo.directoryList[j]
	})

	err = copyFiles(app, &fInfo)
	if err != nil {
		app.ErrorLogger.Printf("error on copy files: %s", err)
		return err
	}

	fmt.Printf("Count created random copy files: %d\n", len(fInfo.randomFilesList))
	fmt.Printf("Total files after random copy: %d\n", len(fInfo.allFilesList)+len(fInfo.randomFilesList))

	return nil
}

func copyFiles(app *config.App, fInfo *filesInfo) error {
	wp := newWorkerPool(app.CountGoroutine)
	rCount := rand.Intn(app.CountRndCopyIter)
	defer wp.wg.Wait()

	for i := 0; i < rCount; i++ {
		wp.wg.Add(1)
		go func(i int) {
			defer func() {
				wp.mu.Unlock()
				// read to release a slot
				<-wp.semaphoreChan
				wp.wg.Done()
			}()
			// block while full
			wp.semaphoreChan <- struct{}{}
			wp.mu.Lock()

			rand.Seed(time.Now().UnixNano())
			rFile := rand.Intn(len(fInfo.allFilesList) - 1)
			rDir := rand.Intn(len(fInfo.directoryList) - 1)

			pathNewFile := fInfo.directoryList[rDir] + "/copy_" + fInfo.allFilesList[rFile].Name
			if _, err := os.Stat(pathNewFile); os.IsNotExist(err) {
				source, err := os.Open(fInfo.allFilesList[rFile].Path)
				if err != nil {
					app.ErrorLogger.Printf("error on open file: %s\n", err)
				}
				defer fileClose(app, source)

				destination, err := os.Create(pathNewFile)
				if err != nil {
					app.ErrorLogger.Printf("error on create file: %s\n", err)
				}
				defer fileClose(app, destination)

				_ = byteCopy(app, source, destination)
				fRand := fInfo.allFilesList[rFile]
				fRand.OriginalFile = &fInfo.allFilesList[rFile]
				fRand.Path = pathNewFile
				fRand.Name = "copy_" + fInfo.allFilesList[rFile].Name
				fInfo.randomFilesList = append(fInfo.randomFilesList, fRand)
			}
		}(i)
	}

	return nil
}
