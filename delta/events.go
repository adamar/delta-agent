
package delta

import (
    "github.com/adamar/delta-server/models"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "github.com/elgs/jsonql"
)



type Rule struct {
    Name        string `json:"name"`
    Channel     string `json:"channel"`
    Query       string `json:"query"`
}

func ParseRules() []Rule {
    raw, _ := ioutil.ReadFile("./conf.d/rules.json")
    keys := make([]Rule,0)
    json.Unmarshal(raw, &keys)
    //fmt.Printf("%#v", keys)
    return keys
}


func ParseEvents(blerg <-chan interface{})  {

	rules := ParseRules()

        for {
                select {
                case in := <-blerg:
                        var jj = in.(*models.Event)
			matchEvent(jj.EventType, jj.Data, rules)
                }
        }

}

func matchEvent(msgType string, msg string, ruleset []Rule) {
    
    for _, r := range ruleset {
        //fmt.Println(msgType)
        parser, err := jsonql.NewStringQuery(msg)
        if err != nil {
            fmt.Println(err)
        }
        ret, _ := parser.Query(r.Query)
	if ret != nil {
        	fmt.Println("FOUND event matching in rule", r.Query, " with event ", msg)
	}
    }


}

