
package delta

import (
    "github.com/adamar/delta-server/models"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "github.com/elgs/jsonql"
)


type Rules struct {
    Ruleset []Rule
}


type Rule struct {
    Name        string `json:"name"`
    Channel     string `json:"channel"`
    Query       string `json:"query"`
}

func ParseRules() {
    raw, _ := ioutil.ReadFile("./conf.d/rules.json")
    keys := make([]Rule,0)
    json.Unmarshal(raw, &keys)
    fmt.Printf("%#v", keys)
}


func ParseEvents(blerg <-chan interface{})  {

        for {
                select {
                case in := <-blerg:
                        var jj = in.(*models.Event)
                        fmt.Println(jj.Data)
                        fmt.Println(jj.EventType)
                }
        }

}

func matchEvent(msgType string, msg string) {

    fmt.Println(msgType)
    parser, err := jsonql.NewStringQuery(msg)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(parser.Query("exe='/usr/bin/su' || exe='/usr/bin/sudo'"))

}

