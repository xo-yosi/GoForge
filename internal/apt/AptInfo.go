package apt

import (
	"os"
	"os/exec"
)

func InfoApt(pkg string) error {
	cmd := exec.Command("apt-cache", "show", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}