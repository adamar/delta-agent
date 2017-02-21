
package delta

import (
	"os/user"
	"encoding/gob"
	//"github.com/cskr/pubsub"
	"github.com/adamar/delta-server/models"
	"github.com/valyala/gorpc"
)

type DeltaCore struct {
	Rpc     *gorpc.Client
}


func Start() (*DeltaCore, error) {

	dc := &DeltaCore{}
	err := dc.PreflightChecks()
	if err != nil {
		return nil, err
	}

        gob.RegisterName("Response", models.Response{})
        gob.RegisterName("Event", models.Event{})

	dc.NewRPClient()

	go dc.StartAuditEngine()
	go dc.StartProcFSEngine()
	go dc.StartiNotifyEngine()
	go dc.StartLogStreamEngine()

	return dc, nil

}

	
func (dc *DeltaCore) PreflightChecks() error {

        user, err := user.Current()
        if err != nil {
                return err
        }

        if user.Uid != "0" {
                return nil
        }

        return nil

}
