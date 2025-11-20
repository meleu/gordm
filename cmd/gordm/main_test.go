package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
)

var binName = "gordm-test"

func TestMain(m *testing.M) {
	fmt.Println("Building the CLI...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		log.Fatalf("cannot build the CLI %q: %s", binName, err)
	}

	fmt.Println("Running the tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	_ = os.Remove(binName)
	os.Exit(result)
}

func TestGordmCLI(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("generates an integer", func(t *testing.T) {
		cmd := exec.Command(cmdPath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		// convert []byte to string and remove the '\n' char at the end
		outputStr := string(out[:len(out)-1])
		assertCanConvertToInt(t, outputStr)
	})
}

func assertCanConvertToInt(t testing.TB, str string) {
	t.Helper()
	_, err := strconv.Atoi(str)
	if err != nil {
		t.Errorf("expected %q to represent an integer:\n%v", str, err)
	}
}
