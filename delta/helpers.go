package delta

import (
	"github.com/nu7hatch/gouuid"
	"strconv"
	"time"
)

func GenUuid() (string, error) {

	u4, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return u4.String(), nil

}

func GenTimeStamp() string {

	nano := time.Now().UnixNano()
	return strconv.FormatInt(nano, 10)

}
