package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// will accept the name eg. github.com/justinbather/project-name

	// get git project name from command-line eg. github.com/justinbather/project-name
	// parse folder name eg. project-name
	args := os.Args[1:]

	packageName := args[0]

	folderName := strings.Split(packageName, "/")[2]

	err := os.Mkdir(folderName, 0750)
	if err != nil {
		log.Fatalf("Error creating directory. Err: %s", err)
	}

	err = os.Chdir(folderName)
	if err != nil {
		log.Fatalf("Error changing dirs. Err: %s", err)
	}

	gitCmd := exec.Command("git", "init")
	err = gitCmd.Run()
	if err != nil {
		log.Fatalf("Error doing git command. Err: %s", err)
	}

	goCmd := exec.Command("go", "mod", "init", packageName)
	err = goCmd.Run()
	if err != nil {
		log.Fatalf("Error creating go project. Err: %s", err)
	}

	// Open the file with write-only mode, creating it if it doesn’t exist.
	file, err := os.Create("main.go")
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write content to the file
	content := "package main\n// This file was created by scaffold.\nfunc main() {\n\t//stuff goes here:)\n}"
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalf("Error writing to file. Err: %s", err)
	}
}
