package main

import (
	"log"
	"os"
)

const (
	preErr = "[ ERR ] "
	preInf = "[ INF ] "
	preWrn = "[ WRN ] "
	preFtl = "[ FTL ] "
	preDbg = "[ DBG ] "
)

var (
	// LogErr is the logger for error outputs
	LogErr = log.New(os.Stderr, preErr, log.LstdFlags)
	// LogInf ist the logger for info outputs
	LogInf = log.New(os.Stdout, preInf, log.LstdFlags)
	// LogWrn ist the logger for warning outputs
	LogWrn = log.New(os.Stdout, preWrn, log.LstdFlags)
	// LogFtl ist the logger for fatal error outputs
	LogFtl = log.New(os.Stderr, preFtl, log.LstdFlags)
	// LogDbg ist the logger for debug outputs
	LogDbg = log.New(os.Stdout, preDbg, log.LstdFlags|log.Lshortfile)
)
