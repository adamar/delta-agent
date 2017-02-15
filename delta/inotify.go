
package delta

impodelta
    "github.com/scyth/go-webproject/gwp/libs/inotify"
    "log"
)



func StartiNotifyEngine() {


	watcher, err := inotify.NewWatcher()
	if err != nil {
    		log.Fatal(err)
	}

	watcher.Watch("/boot")
	watcher.Watch("/dev")
	watcher.Watch("/etc")
	watcher.Watch("/lib")
	watcher.Watch("/lib64")
	watcher.Watch("/proc")
	watcher.Watch("/usr")
	watcher.Watch("/root")
	watcher.Watch("/usr")



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
