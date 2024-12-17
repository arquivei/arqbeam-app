package app

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

// inferBinaryName will return the name of the current running binary if it was
// compiled. Otherwise, if it is running with "go run", for example, it will return
// an empty string
func inferBinaryName() string {
	bin := os.Args[0]
	if strings.Contains(bin, "go-build") || strings.Contains(bin, "__debug__") {
		// Program is been executed by "go run"
		log.Warn().Msg("[arqbeam-app] Unable to get current binary name. This is expected if the code is being executed by 'go run'.")
		return ""
	}
	log.Trace().Str("worker_binary", bin).Msg("[arqbeam-app] Binary name inferred.")

	return bin
}
