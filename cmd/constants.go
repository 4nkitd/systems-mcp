package cmd

import (
	"fmt"
	"os"

	"github.com/4nkitd/mcp/internal/log"
)

var (
	key       string = "key"
	secret    string = "secret"
	log_dir   string = "log_dir"
	transport string = "transport"
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
