package config

import (
	"crypto/sha256"
	"flag"
	"hash"
	"log"
	"os"
	"path/filepath"
)

type App struct {
	HashAlgorithm    hash.Hash
	ErrorLogger      *log.Logger
	SourcePath       string
	CountGoroutine   int
	CountRndCopyIter int
	FlagDelete       bool
	FlagRandCopy     bool
}

func NewApp() *App {
	return &App{
		SourcePath:       ".",
		CountGoroutine:   1000,
		CountRndCopyIter: 100,
		FlagDelete:       false,
		FlagRandCopy:     false,
	}
}

func (a *App) Init() {
	const (
		usagePath = "use this flag for set source directory"
		usageRm   = "use this flag for delete duplicate files"
		usageCp   = "use this flag for random copy files"
		usageGo   = "use this flag for set max count of goroutines"
	)

	flag.StringVar(&a.SourcePath, "path", a.SourcePath, usagePath)
	flag.BoolVar(&a.FlagDelete, "rm", a.FlagDelete, usageRm)
	flag.BoolVar(&a.FlagRandCopy, "cp", a.FlagRandCopy, usageCp)
	flag.IntVar(&a.CountGoroutine, "go", a.CountGoroutine, usageGo)

	a.ErrorLogger = NewBuiltinLogger().logger
	a.SourcePath = "/home/white/Files"
	a.HashAlgorithm = sha256.New()
	if err := a.setABSPath(); err != nil {
		a.ErrorLogger.Printf("error on get ABS path from source path %q: %v\n", a.SourcePath, err)
	}
}

func (a *App) setABSPath() error {
	// get absolut filepath for source path
	sourcePath, err := filepath.Abs(a.SourcePath)
	if err != nil {
		return err
	}
	a.SourcePath = sourcePath

	return nil
}

type BuiltinLogger struct {
	logger *log.Logger
}

func NewBuiltinLogger() *BuiltinLogger {
	return &BuiltinLogger{logger: log.New(os.Stdout, "", 5)}
}

func (l *BuiltinLogger) Debug(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *BuiltinLogger) Debugf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}
