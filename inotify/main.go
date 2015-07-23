package main

import "golang.org/x/exp/inotify"
import "log"

func main() {

	path := "/tmp"

	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.AddWatch(path, inotify.IN_CLOSE_WRITE)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case ev := <-watcher.Event:
			log.Println("event:", ev)
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}

}
