package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if err := do(); err != nil {
		panic(err)
	}
}

func Command(cmd string, args ...string) (n *exec.Cmd) {
	n = exec.Command(cmd, args...)
	fmt.Fprintf(os.Stderr, "exec: %s %s\n", cmd, strings.Join(args, " "))

	n.Stdout = os.Stderr
	n.Stderr = os.Stderr

	return
}

func cmdOutput(cmd string, params ...string) (output string, err error) {
	p := Command(cmd, params...)
	p.Stdout = nil

	opb, err := p.Output()
	if err != nil {
		return
	}

	return strings.TrimSpace(string(opb)), nil
}

func goList(params ...string) (output string, err error) {
	return cmdOutput("go", append([]string{"list"}, params...)...)
}

func do() (err error) {
	var goPath = string(os.Args[1])
	var remainingArgs = os.Args[2:]

	modRoot := filepath.Join(goPath, "pkg/mod")

	const gogoslick = "github.com/gogo/protobuf/protoc-gen-gogoslick"
	const binary = "github.com/zemnmez/tab/internal/proto/gen/protoc-gen-something"

	// detect our root module path
	ourPkg, err := goList(".")
	if err != nil {
		return
	}

	rootModule, err := goList("-m")
	if err != nil {
		return
	}

	rootModuleDir, err := goList("-f", "{{.Dir}}", "-m", rootModule)
	if err != nil {
		return
	}

	// we are going to make a temporary folder containing
	// a 'vendor' kind of thing to allow protoc to do fully qualified import
	// paths on modules. yes it is a huge faff.

	depList, err := cmdOutput("go", "mod", "graph")
	if err != nil {
		return
	}

	if err != nil {
		return
	}

	// make the temp folder
	protocPath, err := ioutil.TempDir("", "tab-protoc")
	if err != nil {
		return
	}

	defer os.RemoveAll(protocPath)

	if err = os.MkdirAll(filepath.Join(protocPath, filepath.Dir(rootModule)), 0700); err != nil {
		return
	}

	if err = os.Symlink(rootModuleDir, filepath.Join(protocPath, rootModule)); err != nil {
		return
	}

	// get the version used
	gogoSlickPath, err := goList(
		"-f", "{{.Dir}}",
		gogoslick,
	)

	if err != nil {
		return
	}

	// get our location

	binaryName := filepath.Base(binary)

	//binaryLoc := filepath.Join(modRoot, binary)

	pbBase := filepath.Join(gogoSlickPath, "../protobuf/google")

	if err = os.Symlink(pbBase, filepath.Join(protocPath, "google")); err != nil {
		return
	}

	for _, pair := range bytes.Split([]byte(depList), []byte("\n")) {
		// skip if it's an indirect import

		importedBy := string(bytes.Split(pair, []byte(" "))[0])

		if !strings.HasPrefix(importedBy, rootModule) {
			continue
		}

		importedPkg := string(bytes.Split(pair, []byte(" "))[1])

		var dirs = filepath.Join(protocPath, filepath.Dir(importedPkg))
		fmt.Fprintf(os.Stderr, "making dir path %s\n", dirs)

		if err = os.MkdirAll(dirs, 0700); err != nil {
			return
		}

		var oldpath = filepath.Join(modRoot, importedPkg)
		var newpath = filepath.Join(protocPath, strings.Split(importedPkg, "@")[0])

		fmt.Fprintf(os.Stderr, "ln -s %s %s\n", oldpath, newpath)

		if err = os.Symlink(oldpath, newpath); err != nil {
			return
		}
	}

	Command("go", "get", binary).Run()

	var remapsMap = map[string]string{
		"google/protobuf/timestamp.proto": "github.com/gogo/protobuf/types",
		"google/protobuf/duration.proto":  "github.com/gogo/protobuf/types",
		"google/protobuf/struct.proto":    "github.com/gogo/protobuf/types",
		"google/protobuf/wrappers.proto":  "github.com/gogo/protobuf/types",
		"google/protobuf/any.proto":       "github.com/gogo/googleapis/google/api",
	}

	var remapsFlat []string
	for k, v := range remapsMap {
		remapsFlat = append(remapsFlat, strings.Join([]string{k, v}, "="))
	}

	var remaps = strings.Join(remapsFlat, ",M")

	args := []string{
		"protoc",
		"--proto_path", protocPath,
		"--plugin", fmt.Sprintf(
			"%s=%s",
			binaryName,
			filepath.Join(goPath, "bin", binaryName),
		),
		fmt.Sprintf("--%s_out=M%s:.", strings.TrimPrefix(binaryName, "protoc-gen-"), remaps),
	}

	var files []string
	for _, arg := range remainingArgs {
		var extraFiles []string
		if extraFiles, err = filepath.Glob(arg); err != nil {
			return
		}

		files = append(files, extraFiles...)
	}

	for i, fileName := range files {
		files[i] = filepath.Join(protocPath, ourPkg, fileName)
	}

	args = append(args, files...)

	cmd := Command(args[0], args[1:]...)

	cmd.Dir = protocPath

	return cmd.Run()
}
