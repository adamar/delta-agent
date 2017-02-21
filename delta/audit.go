package delta

import (
	"fmt"
	"github.com/cskr/pubsub"
	"github.com/mozilla/libaudit-go"
	"io/ioutil"
	"syscall"
)

var PubSub *pubsub.PubSub

var AuditChannel = "audit"

func (dc *DeltaCore) StartAuditEngine() {

	s, err := libaudit.NewNetlinkConnection()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	defer s.Close()

	// enable audit in kernel
	err = libaudit.AuditSetEnabled(s, 1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// check if audit is enabled
	status, err := libaudit.AuditIsEnabled(s)
	if err == nil && status == 1 {
		fmt.Printf("Enabled Audit\n")
	} else if err == nil && status == 0 {
		fmt.Printf("Audit Not Enabled\n")
		return
	} else {
		fmt.Printf("%v\n", err)
		return
	}

	// set the maximum number of messages
	// that the kernel will send per second
	err = libaudit.AuditSetRateLimit(s, 450)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// set max limit audit message queue
	err = libaudit.AuditSetBacklogLimit(s, 16438)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// register current pid with audit
	err = libaudit.AuditSetPID(s, syscall.Getpid())
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// delete all rules that are previously present in kernel
	err = libaudit.DeleteAllRules(s)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// set audit rules
	// specify rules in JSON format (for example see: https://github.com/arunk-s/gsoc16/blob/master/audit.rules.json)
	out, _ := ioutil.ReadFile("conf.d/audit.rules.json")
	err = libaudit.SetRules(s, out)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// create a channel to indicate libaudit to stop collecting messages
	done := make(chan bool)
	// spawn a go routine that will stop the collection after 5 seconds
	//go func(){
	//	time.Sleep(time.Second*5)
	//	done <- true
	//}()
	// collect messages and handle them in a function
	libaudit.GetAuditMessages(s, callback, &done)
}

// provide a function to handle the messages
func callback(msg *libaudit.AuditEvent, ce error, args ...interface{}) {
	if ce != nil {
		fmt.Printf("%v\n", ce)
	} else if msg != nil {

		key := genKeyName(AuditChannel, msg.Type)
		event := BuildEvent(msg.Serial, msg.Timestamp, key, msg.Data)
		event.PublishEvent(AuditChannel)

	}
}
