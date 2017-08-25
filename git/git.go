package git

import (
	"fmt"

	"github.com/TobiasBales/shell-prompt/exec"
	"github.com/logrusorgru/aurora"
)

func getBranchName() string {
	stdout, _, _ := exec.Execute("git", "rev-parse", "--abbrev-ref", "HEAD")
	return stdout
}

func branchIndicator() string {
	branch := getBranchName()
	if len(branch) > 0 {
		return aurora.Green(branch).String()
	}

	return ""
}

func isDirty() bool {
	stdout, _, _ := exec.Execute("git", "status", "--porcelain")
	return len(stdout) > 0
}

func dirtyIndicator() string {
	dirty := isDirty()
	if dirty {
		return aurora.Green("*").String()
	}

	return ""
}

// Indicator returns the git status of the current working dir
func Indicator() string {
	branch := branchIndicator()
	dirty := dirtyIndicator()

	if len(branch) == 0 && len(dirty) == 0 {
		return ""
	}

	return fmt.Sprintf(" %v%v", branch, dirty)
}
