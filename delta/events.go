
package delta

import (
    "github.com/adamar/delta-server/models"
    "fmt"
    "encoding/json"
    "io/ioutil"
)


type Rules struct {
    Ruleset []Rule
}


type Rule struct {
    Name        string `json:"name"`
    Channel     string `json:"channel"`
    Query       string `json:"query"`
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




