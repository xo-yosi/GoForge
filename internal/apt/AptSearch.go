package apt

import (
	"os"
	"os/exec"
)

func SearchApt(pkg string) error {
	cmd := exec.Command("apt-cache", "search", pkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
