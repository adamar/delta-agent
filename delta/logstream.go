package delta

import (
	"github.com/coreos/go-systemd/sdjournal"
	"log"
)

var LogChannel = "logstream"

func (dc *DeltaCore) StartLogStreamEngine() {

	j, err := sdjournal.NewJournal()
	if err != nil {
		log.Println(err)
		return
	}
	err = j.SeekTail()
	if err != nil {
		log.Println(err)
		return
	}
	for {
		n, _ := j.Next()

		if n < 1 {
			j.Wait(sdjournal.IndefiniteWait)
			continue
		}

		m, err := j.GetEntry()
		if err != nil {
			log.Printf("Error: " + err.Error())
		} else {

			key := genKeyName(LogChannel, "log_event")

			if _, ok := m.Fields["_SOURCE_REALTIME_TIMESTAMP"]; ok {
				event := BuildEvent(m.Fields["_SOURCE_REALTIME_TIMESTAMP"], 
                                                    m.Fields["_SOURCE_REALTIME_TIMESTAMP"], 
                                                    key, 
                                                    m.Fields)
				event.PublishEvent(LogChannel)
			} else {
				event := BuildEvent(m.Fields["_SOURCE_MONOTONIC_TIMESTAMP"], 
                                                    m.Fields["_SOURCE_MONOTONIC_TIMESTAMP"], 
                                                    key, 
                                                    m.Fields)
				event.PublishEvent(LogChannel)
			}

		}

	}
}
