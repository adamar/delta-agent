package delta

import (
	"fmt"
	//"io/ioutil"
	//"syscall"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
	"strconv"
	"github.com/shirou/gopsutil/process"
)

var ProcfsChannel = "procfs"

func (dc *DeltaCore) StartProcFSEngine() {

	// Wait for initla File Descriptors to be opened
	// before monitoring begins
	//
	time.Sleep(2 * time.Second)

	existing := []string{}
	latest := []string{}

	for {

		latest, _ = filepath.Glob("/proc/[0-9]*/fd/[0-9]*")

		// Check if this is the first check
		// if so set the previous FD list
		// to match the current list
		//
		if len(existing) == 0 {

			existing = latest

		}

		// Convert the FD Glob to a String for comparison
		// some bad assumptions live here
		//
		if fmt.Sprintf("%v", existing) != fmt.Sprintf("%v", latest) {

			//pid := strconv.Itoa(os.Getpid())

			ts := GenTimeStamp()

			// Find FDs added and removed
			//
			added, removed := differ(existing, latest)

			for _, a := range added {
				collectData(a, ts, "FD_OPEN")
			}
			for _, r := range removed {
				collectData(r, ts, "FD_CLOSE")
			}

			existing = latest

		}

		time.Sleep(100 * time.Millisecond)

	}

}

// Collect data about a file describ=ptor being
// opened by a process
func collectData(pathe string, ts string, evttype string) {

	// TODO cleanup this function
	// and the functions it calls
	substrings := strings.Split(pathe, "/")

	pid := substrings[2]
	fileDesc := substrings[4]

	data := map[string]string{"pid": pid, "fd": fileDesc}

	// If File descriptor is being opened, gather
	// info on the event
	//
	if evttype == "FD_OPEN" {

    		i, _ := strconv.Atoi("1")

    		pro, err := process.NewProcess(int32(i))
    		if err != nil {
        		panic(err)
    		}

    		cmd, _ := pro.Cmdline()
    		create, _ := pro.CreateTime()

		data["cmdline"] = cmd
		data["start_time"] = strconv.FormatInt(create, 10)
		data["path"], data["type"] = readFD(pathe)
	}

	key := genKeyName(ProcfsChannel, evttype)
	event := BuildEvent(ts, ts, key, data)
	event.PublishEvent(ProcfsChannel)

}

func differ(oldest []string, newest []string) ([]string, []string) {

	added := []string{}
	removed := []string{}
	m := map[string]int{}

	// Convert the slive to a Map
	// for easier comparison
	//
	for _, orig := range oldest {
		m[orig] = 1
	}

	for _, add := range newest {
		m[add] = m[add] + 2
	}

	for k, v := range m {
		if v == 1 {
			removed = append(removed, k)
		} else if v == 2 {
			added = append(added, k)
		}

	}

	return added, removed

}

func readFD(fileDesc string) (string, string) {

	// Read linked File descriptors
	// to get more info on the type
	//
	link, err := os.Readlink(fileDesc)

	if err != nil {
		log.Println(err.Error())
		return "", ""
	}

	typ := linkType(link)

	return link, typ

}


// Return the File type
func linkType(fileDesc string) string {

	//TODO - return with a more appropriate type
	if strings.HasPrefix(fileDesc, "socket") {
		return "socket"
	}

	if strings.HasPrefix(fileDesc, "pipe") {
		return "pipe"
	}

	if strings.HasPrefix(fileDesc, "anon_inode") {
		return "inode"
	}

	if strings.HasPrefix(fileDesc, "/dev") {
		return "dev"
	}

	if strings.HasPrefix(fileDesc, "/proc") {
		return "proc"
	}

	return "file"

}
