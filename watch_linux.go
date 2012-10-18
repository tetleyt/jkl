package main

import (
	"github.com/dersebi/golang_exp/exp/inotify"
	"strings"

	"fmt"
)

func watch(site *Site) {
	fmt.Printf("Listening for changes to %s\n", site.Src)

	// Setup the inotify watcher
	watcher, err := inotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Only the events we care about
	const flags = inotify.IN_MODIFY | inotify.IN_DELETE | inotify.IN_CREATE | inotify.IN_MOVE

	// Get recursive list of directories to watch
	for _, path := range dirs(site.Src) {
		if err := watcher.AddWatch(path, flags); err != nil {
			fmt.Println(err)
			return
		}
	}

	for {
		select {
		case ev := <-watcher.Event:
			// Ignore changes to the _site directoy, hidden, or temp files		
			if !strings.HasPrefix(ev.Name, site.Dest) && !isHiddenOrTemp(ev.Name) {
				fmt.Println("Event:", ev.Name)
				recompile(site)
			}
		case err := <-watcher.Error:
			fmt.Println("inotify error:", err)
		}
	}
}
