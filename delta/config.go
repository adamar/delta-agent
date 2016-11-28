
package delta

import (
	"io/ioutil"
	"encoding/json"
	"os/user"
	)


type Triggers struct {
    Triggers []Trigger `json:"triggers"`
}

type Trigger struct {
    EventType string `json:"eventtype"`
    Subscriber string `json:"subscriber"`
}

func LoadTriggers() *Triggers {

        content, err := ioutil.ReadFile("triggers.d/triggers.json")
        if err != nil {
                panic(err)
        }

        triggers := Triggers{}

        err = json.Unmarshal(content, &triggers)
        if err != nil {
                //return errors.Wrap(err, "SetRules failed")
                panic(err)
        }

        return &triggers

}


func PassedPreflighChecks() bool {

	user, err := user.Current()
	if err != nil {
		return false
	}

	if user.Uid != "0" {
		return false
	}

	return true

}




