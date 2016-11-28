
package delta

import (
    "bytes"
    "encoding/json"
    //"time"
    "github.com/adamar/delta-server/models"
)



func BuildEvent(serial string, msgType string, native string, data map[string]string) *models.Event {

    var uuid, err = GenUuid()
    if err != nil {
        panic(err)
    }

    timestamp := GenTimeStamp()

    flatData := new(bytes.Buffer)
    enc := json.NewEncoder(flatData)
    err = enc.Encode(data)
    if err != nil {
        panic(err)
    }

    evt := &models.Event{
        Uuid:      uuid,
        Serial:    serial,
        TimeStamp: timestamp,
        NativeTimeStamp: native,
        EventType: msgType,
        Data:      flatData.String(),
    }

    return evt

}




