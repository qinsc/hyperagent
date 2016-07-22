package host

import (
	"fmt"
	"os"
)

func GetHostName() string {
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Println("Get hostname failed.")
	}
	return hostName
}
