package pwdx

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func Pwdx(pid int) Output {
	pidStr := strconv.Itoa(pid)
	cmd := exec.Command("pwdx", pidStr)
	output, err := cmd.Output()
	outputStr := string(output)
	if err != nil {
		return initOutput(outputStr, pidStr, err)
	}
	dir, err := parse(outputStr)
	return initOutput(dir, pidStr, err)
}

func parse(s string) (string, error) {
	parsed := strings.Split(s, ":")
	if len(parsed) != 2 {
		return "", fmt.Errorf("Wrong number of arguments")
	}
	return strings.TrimSpace(parsed[1]), nil
}

func ManyPwdx(pids ...int) Outputs {
	answer := make(Outputs, 0)
	for _, pid := range pids {
		answer = append(answer, Pwdx(pid))
	}
	return answer
}
