package filework

import (
	"encoding/hex"
	"github.com/White-AK111/GB_Go_Level2/Lesson8/config"
	"io"
	"os"
	"sync"
	"time"
)

type workerPool struct {
	wg            sync.WaitGroup
	resultChan    chan FileEntity
	semaphoreChan chan struct{}
	mu            sync.Mutex
}

type filesInfo struct {
	allFilesList       []FileEntity
	duplicateFilesList []FileEntity
	deleteFilesList    []FileEntity
	randomFilesList    []FileEntity
	directoryList      []string
}

func newWorkerPool(N int) *workerPool {
	return &workerPool{
		wg:            sync.WaitGroup{},
		resultChan:    make(chan FileEntity, N),
		semaphoreChan: make(chan struct{}, N),
	}
}

type FileEntity struct {
	OriginalFile *FileEntity
	Create       time.Time
	Name         string
	Path         string
	Hash         string
	Size         int64
}

// getHashOfFile method get hash of file
func (f *FileEntity) getHashOfFile(app *config.App) error {

	file, err := os.Open(f.Path)
	if err != nil {
		return err
	}
	defer fileClose(app, file)

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
		if fl[i].Hash == f.Hash {
			f.OriginalFile = &fl[i]
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

func findAllFiles(app *config.App, fInfo *filesInfo) error {
	wp := newWorkerPool(app.CountGoroutine)
	defer wp.wg.Wait()

	wp.wg.Add(1)
	lsFiles(app.SourcePath, app, wp, fInfo)

	return nil
}

func lsFiles(dir string, app *config.App, wp *workerPool, fInfo *filesInfo) {
	// block while full
	wp.semaphoreChan <- struct{}{}

	go func() {
		defer func() {
			wp.mu.Unlock()
			// read to release a slot
			<-wp.semaphoreChan
			wp.wg.Done()
		}()

		wp.mu.Lock()
		file, err := os.Open(dir)
		if err != nil {
			app.ErrorLogger.Println("error opening directory: %s\n", err)
		}

		defer fileClose(app, file)

		files, err := file.Readdir(-1) // Loads all children files into memory.
		if err != nil {
			app.ErrorLogger.Println("error reading directory: %s\n", err)
		}

		for _, f := range files {
			path := dir + "/" + f.Name()
			if f.IsDir() {
				fInfo.directoryList = append(fInfo.directoryList, path)
				wp.wg.Add(1)
				go lsFiles(path, app, wp, fInfo)
			} else {
				fe := NewFileEntity()
				fe.Name = f.Name()
				fe.Path = path
				fe.Create = f.ModTime()
				if err = fe.getHashOfFile(app); err != nil {
					app.ErrorLogger.Printf("can't get hash of file %s: %s\n", path, err)
				}
				fe.Size = f.Size()
				fInfo.allFilesList = append(fInfo.allFilesList, *fe)
			}
		}
	}()
}

func fileClose(app *config.App, file *os.File) {
	err := file.Close()
	if err != nil {
		app.ErrorLogger.Printf("error on defer close file %s: %s\n", file.Name(), err)
	}
}

func byteCopy(app *config.App, source *os.File, destination *os.File) error {
	buf := make([]byte, 1024)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			app.ErrorLogger.Printf("error on byte read from file: %s\n", err)
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			app.ErrorLogger.Printf("error on byte read to file: %s\n", err)
			return err
		}
	}

	return nil
}
