
package delta

import (
	"github.com/cskr/pubsub"
)

type DeltaCore struct {
	pubsub	pubsub.PubSub
}


func Start() *DeltaCore {

	return &DeltaCore{}

}

	
