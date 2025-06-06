package pklog

import (
	"log"
	"log/syslog"
	"os"
	"path/filepath"

	"github.com/pezhmankasraee/pklog/v2/baseconfig"
)

const productionLifeCycle string = "PRODUCTION"
const errorMessageEnding string = baseconfig.Reset + " ] "

// CreateLog prints out a log message depends on application life cycle.
// This module reads first the result of environment parameter PKLOG.
// If PKLOG is equal to "PRODUCTION", the logs will be generated in /var/log/syslog otherwise
// the logs will be displayed with color on standard display.
// logLevel accepts four types of log level: FatalError(0), Error (100), Warning (200) and Info (300).
// s is the log message which is of type string.
func CreateLog(logLevel int, s string) {

	softwareLifeCycle := os.Getenv("PKLOG")
	appName := filepath.Base(os.Args[0])

	logwriter, err := syslog.New(syslog.LOG_DEBUG, appName)
	if err != nil {
		panic(" [ panic ] : pklog cannot write on syslog")
	}

	switch logLevel {
	case FatalError:
		if softwareLifeCycle == productionLifeCycle {
			log.SetPrefix("[ Fatal Error ] ")
		} else {
			log.SetPrefix("[ " + baseconfig.FGMagneta + "Fatal Error" + errorMessageEnding)
		}
		break
	case Error:
		if softwareLifeCycle == productionLifeCycle {
			log.SetPrefix("[ Error ] ")
		} else {
			log.SetPrefix("[ " + baseconfig.FGRed + "Error" + errorMessageEnding)
		}
		break
	case Warning:
		if softwareLifeCycle == productionLifeCycle {
			log.SetPrefix("[ Warning ] ")
		} else {
			log.SetPrefix("[ " + baseconfig.FGYellow + "Warning" + errorMessageEnding)
		}
		break
	case Information:
		if softwareLifeCycle == productionLifeCycle {
			log.SetPrefix("[ Info ] ")
		} else {
			log.SetPrefix("[ " + baseconfig.FGTeal + "Info" + errorMessageEnding)
		}
		break
	default:
		panic("[ Panic ] Log Level is invalid")
	}

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.LUTC)
	if logLevel < 100 {
		log.Fatal(s)
	}

	if softwareLifeCycle == productionLifeCycle {
		log.SetOutput(logwriter)
	} else {
		log.SetOutput(os.Stdout)
	}

	log.Println(s)
}
