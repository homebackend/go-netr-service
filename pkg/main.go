package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/akamensky/argparse"
	homecommon "github.com/homebackend/go-homebackend-common/pkg"
)

const (
	PROG_NAME = "netrservice"
	CONF_FILE = "/etc/netr/config.yaml"
)

func service(c *string) {
	homecommon.CheckPrerequisites(homecommon.O_ANY, *c, []string{})
	pidFile := homecommon.CreatePidFile()

	defer func() {
		homecommon.StopIpc(PROG_NAME)
		pidFile.Unlock()
	}()

	sigc := homecommon.Signal()

	log.Printf("Service started.")

	for {
		select {
		case s := <-sigc:
			log.Printf("Signal captured: %s", s)
			return
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func status() {
}

func main() {
	parser := argparse.NewParser(os.Args[0], "Netr backend")

	startCommand := parser.NewCommand("start", "Start the netr service")
	stopCommand := parser.NewCommand("stop", "Stop the netr service")
	statusCommand := parser.NewCommand("status", "Show status of the netr service")
	addLocationCommand := parser.NewCommand("add-location", "Add Location")
	delLocationCommand := parser.NewCommand("del-location", "Delete Location")
	alterLocationCommand := parser.NewCommand("alter-location", "Alter Location")
	addCameraCommand := alterLocationCommand.NewCommand("add-camera", "Add Camera Configuration")
	delCameraCommand := alterLocationCommand.NewCommand("del-camera", "Delete Camera Configuration")
	alterCameraCommand := alterLocationCommand.NewCommand("alter-camera", "Alter Camera Configuration")

	c := startCommand.String("c", "configuration-file", &argparse.Options{
		Required: false,
		Default:  CONF_FILE,
		Help:     "Configuration File",
	})

	la := addLocationCommand.StringPositional(&argparse.Options{
		Required: true,
		Help:     "Location to add",
	})

	ld := delLocationCommand.StringPositional(&argparse.Options{
		Required: true,
		Help:     "Location to delete",
	})

	ldr := delLocationCommand.Flag("r", "recurse", &argparse.Options{
		Required: false,
		Default:  false,
		Help:     "Recursively delete location. That is any cameras added will be deleted as well.",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if startCommand.Happened() {
		service(c)
	} else if stopCommand.Happened() {
		homecommon.Stop(PROG_NAME)
	} else if statusCommand.Happened() {
		status()
	} else if addLocationCommand.Happened() {
		log.Printf("Location: %s", *la)
	} else if delLocationCommand.Happened() {
		log.Printf("Location: %s, recursive: %t", *ld, *ldr)

	} else if alterCameraCommand.Happened() {
		if addCameraCommand.Happened() {

		} else if delCameraCommand.Happened() {

		} else if alterCameraCommand.Happened() {

		}
	}
}
