package apt

import (
	"os"
	"os/exec"
)

func InstallApt(pkg string) error {
    cmd := exec.Command("sudo", "apt-get", "install", "-y", pkg)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return cmd.Run()
}