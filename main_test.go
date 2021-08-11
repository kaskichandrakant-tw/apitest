package apitest

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Waiting for container to start")
	exec.Command("sh","test_setup.sh","start").Output()

	fmt.Println("Running tests")
	code := m.Run()

	fmt.Println("Waiting for container to stop")
	exec.Command("sh","test_setup.sh","stop").Output()
	os.Exit(code)
}