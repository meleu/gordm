package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

var binName = "gordm-test"

// ----------------------------------------------------------------------
// build and test the binary
// ----------------------------------------------------------------------

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

// ----------------------------------------------------------------------

func TestGordmCLI(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("generates an integer", func(t *testing.T) {
		got := assertRunCmd(t, cmdPath)
		assertCanConvertToInt(t, got)
	})

	t.Run("accepts '--min NUM' to specify the min value", func(t *testing.T) {
		// using min = max to generate a deterministic output
		maxStr := strconv.Itoa(defaultMax)
		got := assertRunCmd(t, cmdPath, "--min", maxStr)
		assertEqualString(t, got, maxStr)
	})

	t.Run("accepts '--max NUM' to specify the max value", func(t *testing.T) {
		// using max = min to generate a deterministic output
		minStr := strconv.Itoa(defaultMin)
		got := assertRunCmd(t, cmdPath, "--max", minStr)
		assertEqualString(t, got, minStr)
	})

	t.Run("fails if min > max", func(t *testing.T) {
		got := refuteRunCmd(t, cmdPath, "--min", "999", "--max", "0")
		if !strings.HasPrefix(got, "invalid input: ") {
			t.Errorf("expected output to start with \"invalid input: \", got %q", got)
		}
	})
}

// ----------------------------------------------------------------------
// test helper functions
// ----------------------------------------------------------------------

func assertCanConvertToInt(t testing.TB, str string) {
	t.Helper()
	_, err := strconv.Atoi(str)
	if err != nil {
		t.Errorf("expected %q to represent an integer:\n%v", str, err)
	}
}

func assertEqualString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("\n got: %q\nwant: %q", got, want)
	}
}

func assertRunCmd(t testing.TB, args ...string) string {
	cmdPath := args[0]
	cmdArgs := args[1:]

	cmd := exec.Command(cmdPath, cmdArgs...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	// convert []byte to string and remove the '\n' char at the end
	return string(out[:len(out)-1])
}

func refuteRunCmd(t testing.TB, args ...string) string {
	cmdPath := args[0]
	cmdArgs := args[1:]

	cmd := exec.Command(cmdPath, cmdArgs...)
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("Expected an error, got none.")
	}

	// convert []byte to string and remove the '\n' char at the end
	return string(out[:len(out)-1])
}
