package config

import (
	"log"
	"os"
	"time"
)

func WatchFiles(ch chan string, filesToWatch []string) {
	prevModTimes := make([]time.Time, len(filesToWatch))

	for index, file := range filesToWatch {
		if fileInfo, err := os.Stat(file); err != nil {
			log.Fatalln("Cannot read config file")
		} else {
			prevModTimes[index] = fileInfo.ModTime()
		}
	}

	for {
		checkForModification(ch, filesToWatch, prevModTimes)
		time.Sleep(2 * time.Second)
	}
}

func checkForModification(ch chan string, filesToWatch []string, prevModTimes []time.Time) {
	for index, file := range filesToWatch {
		if fileInfo, err := os.Stat(file); err != nil {
			log.Fatalln("Cannot read config file")
		} else if fileInfo.ModTime().After(prevModTimes[index]) {
			log.Println("File Modified")
			ch <- file + " modified"
			prevModTimes[index] = fileInfo.ModTime()
		}
	}
}
