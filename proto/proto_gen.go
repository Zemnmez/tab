package main

import (
	"os/exec"
	"path/filepath"
	"os"
	"strings"
	"fmt"
)

func main() {
	if err := do(); err != nil { panic(err) }
}

func Command(cmd string, args ...string) (n *exec.Cmd) {
	n = exec.Command(cmd, args...)
	fmt.Fprintf(os.Stderr, "exec: %s %s\n", cmd, strings.Join(args, " "))

	n.Stdout = os.Stderr
	n.Stderr = os.Stderr

	return
}

func do() (err error) {
	var goPath = string(os.Args[1])
	const binary = "github.com/gogo/protobuf/protoc-gen-gogoslick"

	// get the version used
	binPath, err := exec.Command(
		"go",
		 "list",
		 "-f", "{{.Dir}}",
		  binary,
		).Output()

if err != nil { return }

	
	binaryName := filepath.Base(binary)

	modRoot := filepath.Join(goPath, "pkg/mod")

	//binaryLoc := filepath.Join(modRoot, binary)

	pbBase := filepath.Join(string(binPath), "../protobuf")

	Command("go", "get", binary).Run()

	args := []string {
		"protoc",
		"-I=.",
		"--plugin", fmt.Sprintf(
			"%s=%s",
			binaryName, 
			filepath.Join(goPath, "bin", binaryName),
		),
		"-I", pbBase,
		"-I", modRoot, 
		fmt.Sprintf("--%s_out=.", strings.TrimPrefix(binaryName, "protoc-gen-")),
	}
	
	args = append(args, os.Args[2:]...)

	cmd := Command(args[0], args[1:]...)



	return cmd.Run()
}