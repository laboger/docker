// Code forked from Docker project
package daemon

import (
	"errors"
	"net"
	"os"
)

// ErrSdNotifyNoSocket is an error returned if no socket was specified.
var ErrSdNotifyNoSocket = errors.New("No socket")

// SdNotify sends a message to the init daemon. It is common to ignore the error.
func SdNotify(state string) error {
	socketAddr := &net.UnixAddr{
		Name: os.Getenv("NOTIFY_SOCKET"),
		Net:  "unixgram",
	}

	if socketAddr.Name == "" {
		return ErrSdNotifyNoSocket
	}

	conn, err := net.DialUnix(socketAddr.Net, nil, socketAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(state))
	return err
}
