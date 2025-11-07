package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("git diff --staged failed: %v (%s)", err, stderr.String())
	}

	return out.String(), nil
}
