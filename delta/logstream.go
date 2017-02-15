package delta

import (
	"github.com/coreos/go-systemd/sdjournal"
	"log"
)

func StartLogStreamEngine() {

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

			if _, ok := m.Fields["_SOURCE_REALTIME_TIMESTAMP"]; ok {
				event := BuildEvent(m.Fields["_SOURCE_REALTIME_TIMESTAMP"], m.Fields["_SOURCE_REALTIME_TIMESTAMP"], "LOG_EVENT", m.Fields)
				event.PublishEvent("LogEvent")
			} else {
				event := BuildEvent(m.Fields["_SOURCE_MONOTONIC_TIMESTAMP"], m.Fields["_SOURCE_MONOTONIC_TIMESTAMP"], "LOG_EVENT", m.Fields)
				event.PublishEvent("LogEvent")
			}

		}

	}
}
