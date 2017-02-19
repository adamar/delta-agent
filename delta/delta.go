
package delta

import (
	"os/user"
	"github.com/cskr/pubsub"
)

type DeltaCore struct {
	pubsub	pubsub.PubSub
}


func Start() (*DeltaCore, error) {

	dc := &DeltaCore{}
	err := dc.PreflightChecks()
	if err != nil {
		return nil, err
	}
	return dc, nil

}

	
func (dc *DeltaCore) PreflightChecks() bool {

        user, err := user.Current()
        if err != nil {
                return false
        }

        if user.Uid != "0" {
                return false
        }

        return true

}
