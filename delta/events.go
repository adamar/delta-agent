
package delta

import (
    "github.com/adamar/delta-server/models"
    "fmt"
)


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



