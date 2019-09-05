package selfquery

import (
	"flag"
	"os/exec"
)

type Server struct {

}

func New(basePath string, fileGlobs ...string) (gql Server, err error) {
	confFile, err := ioutil.TempFile("", "selfquerygen")
	if err != nil { return }

	defer confFile.Close()

	if err = yaml.NewEncoder(confFile).Encode(config.Config {
			SchemaFilename: config.StringList(fileGlobs),
			Exec: config.PackageConfig {
				Filename: "generated.go",
			},

			Model: config.PackageConfig {
				Filename: "models_gen.go"
			}
	}); err != nil { return }


	cmd := exec.Command(
		"go",
		"run",
		"github.com/99designs/gqlgen",
		"--config", confFile.Name(),
	)

	cmd.Dir = basePath
	cmd.Stderr = os.Stderr
}


func Query(yamlFile string) {

}