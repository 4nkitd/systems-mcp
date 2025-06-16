package cmd

import (
	"fmt"
	"os"

	"github.com/4nkitd/systems-mcp/internal/log"
)

var (
	log_dir   string = "log_dir"
	transport string = "transport"
	host      string = "host"
	port      string = "port"
)

func ParseLogDir(dir string) string {
	if dir == "" {

		var dir_err error
		dir, dir_err = os.Getwd()
		if dir_err != nil {
			log.Write("ERROR", fmt.Sprintf("Error getting current directory: %v", dir_err))
			dir = "."
		}
	}
	return dir
}
