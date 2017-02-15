
package delta

import (
    "github.com/scyth/go-webproject/gwp/libs/inotify"
    "log"
)



func StartiNotifyEngine() {


	watcher, err := inotify.NewWatcher()
	if err != nil {
    		log.Fatal(err)
	}

	locs := []string{"/boot", "/dev", "/etc", "/lib", "/lib64", "/proc", "/usr", "/root", "/usr", "/home"}	

	for _, loc := range locs {
		watcher.Watch(loc)
	}

        //event := BuildEvent(msg.Serial, msg.Type, msg.Timestamp, msg.Data)
        //event.PublishEvent("SystemCall")

	for {
    		select {
    		case ev := <-watcher.Event:
    		switch ev.Mask {
    		case 2:
        		log.Println("File Modified: ", ev.Name)
    		case 4:
        		log.Println("File Attributes Changed: ", ev.Name)
    		case 100:
        		log.Println("File created: ", ev.Name)
    		case 200, 400:
        		log.Println("File Deleted", ev.Name)
    		case 512:
        		log.Println("File Deleted? ", ev.Name)

    			}

    		}

	}




}
