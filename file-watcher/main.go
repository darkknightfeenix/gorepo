package main

import (
	config "file-watcher/config"
	web "file-watcher/web"
	"fmt"
)

func main() {
	go watchAndReload()
	web.SampleApi()
}

func watchAndReload() {
	filesToWatch := []string{"config/config.yaml"}
	ch := make(chan string)
	go config.WatchFiles(ch, filesToWatch)
	for {
		_, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		web.Initialize()
	}
}
