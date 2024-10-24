package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/browser"
)

func main() {
	if len(os.Args) > 2 {
		panic(fmt.Errorf("too many arguments"))
	}

	if len(os.Args) == 2 {
		repository := os.Args[1]
		openRepo(repository)
	}

	if len(os.Args) == 1 {
		repositoryOriginOutput, err := exec.Command("git", "config", "--local", "remote.origin.url").Output()
		checkError(err)

		repositoryOrigin := string(repositoryOriginOutput)
		repository := strings.Split(strings.TrimSpace(repositoryOrigin), ":")[1]
		openRepo(repository)
	}
}

func openRepo(repository string) {
	browser.OpenURL(fmt.Sprintf("https://github.com/%s", repository))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// hola
