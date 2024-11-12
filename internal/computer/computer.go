package computer

import (
	"log"
	"os/exec"
)

func ShutdownComputer() {
	var cmd *exec.Cmd
	cmd = exec.Command("commands/shutdown-script.sh")
	err := cmd.Run()
	if err != nil {
		log.Panic(err)
	}
}

func SleepComputer() {
	var cmd *exec.Cmd
	cmd = exec.Command("commands/sleep-script.sh")
	err := cmd.Run()
	if err != nil {
		log.Panic(err)
	}
}
