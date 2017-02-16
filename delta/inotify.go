
package delta

import (
    "github.com/scyth/go-webproject/gwp/libs/inotify"
    "log"
)

var InotifyChannel = "inotify"

func StartiNotifyEngine() {


	watcher, err := inotify.NewWatcher()
	if err != nil {
    		log.Fatal(err)
	}

	locs := []string{"/boot", "/dev", "/etc", "/lib", "/lib64", "/proc", "/usr", "/root", "/usr", "/home"}	

	for _, loc := range locs {
		watcher.Watch(loc)
	}

        key := genKeyName(InotifyChannel, "event")
 
	for {
    		select {
    		case ev := <-watcher.Event:

		ts := GenTimeStamp()

    		switch ev.Mask {
    		case 2:
			data := map[string]string{"file": ev.Name, "action": "file modified"}
        		event := BuildEvent(ts, ts, key, data)
        		event.PublishEvent(InotifyChannel)

    		case 4:
			data := map[string]string{"file": ev.Name, "action": "file attributes modified"}
        		event := BuildEvent(ts, ts, key, data)
        		event.PublishEvent(InotifyChannel)

    		case 100:
			data := map[string]string{"file": ev.Name, "action": "file created"}
        		event := BuildEvent(ts, ts, key, data)
        		event.PublishEvent(InotifyChannel)

    		case 200, 400:
			data := map[string]string{"file": ev.Name, "action": "file deleted"}
        		event := BuildEvent(ts, ts, key, data)
        		event.PublishEvent(InotifyChannel)

    		case 512:
			data := map[string]string{"file": ev.Name, "action": "file deleted"}
        		event := BuildEvent(ts, ts, key, data)
        		event.PublishEvent(InotifyChannel)

    			}

    		}

	}




}
