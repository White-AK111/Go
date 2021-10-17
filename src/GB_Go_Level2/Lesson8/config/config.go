package config

import (
	"crypto/sha256"
	"flag"
	"hash"
	"log"
	"sync"
)

type App struct {
	Pool          chan struct{}
	WG            sync.WaitGroup
	MU            sync.Mutex
	HashAlgorithm hash.Hash
	ErrorLogger   *log.Logger
	SourcePath    string
	FlagDelete    bool
	FlagRandCopy  bool
}

func NewApp() *App {
	return &App{
		WG:           sync.WaitGroup{},
		SourcePath:   ".",
		FlagDelete:   false,
		FlagRandCopy: false,
	}
}

func (a *App) Init() {
	const (
		usagePath = "use this flag for set source directory"
		usageRm   = "use this flag for delete duplicate files"
		usageCp   = "use this flag for random copy files"
		usageGo   = "use this flag for set max count of goroutines"
	)

	countGo := 1

	flag.StringVar(&a.SourcePath, "path", a.SourcePath, usagePath)
	flag.BoolVar(&a.FlagDelete, "rm", a.FlagDelete, usageRm)
	flag.BoolVar(&a.FlagRandCopy, "cp", a.FlagRandCopy, usageCp)
	flag.IntVar(&countGo, "go", 1, usageGo)

	a.Pool = make(chan struct{}, countGo)
	a.SourcePath = "/home/white/Files"
	a.HashAlgorithm = sha256.New()
}
