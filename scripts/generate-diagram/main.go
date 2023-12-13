package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime/debug"
	"strings"
)

const (
	readmePath                = "README.md"
	generateDiagramReadmePath = "scripts/generate-diagram/README.md"
	diagramPath               = "docs/diagrams"
)

// regex for removing spaces
var space = regexp.MustCompile(`\s+`)

// FocusDiagram is a data structure for generates targeted package
type FocusDiagram struct {
	Name  string
	Focus string
}

// diagrams is the targeted package
var diagrams = []FocusDiagram{
	{
		Name:  "main",
		Focus: "",
	},
	{
		Name:  "di",
		Focus: "internal/di",
	},
	{
		Name:  "handler",
		Focus: "internal/adapters/grpchandler",
	},
	{
		Name:  "usecases",
		Focus: "internal/usecases",
	},
	{
		Name:  "datastore",
		Focus: "internal/infrastructure/datastore",
	},
}

func main() {
	changedFiles, newGenerateDiagramReadme, err := getChangedFile()
	if err != nil {
		log.Fatalf("failed to get changed file: %v", err)
	}

	mainPath, err := getMainPath()
	if err != nil {
		log.Fatalf("failed to get main path: %v", err)
	}

	// create a diagram directory if not exist
	if _, err := os.Stat(diagramPath); os.IsNotExist(err) {
		if err := os.MkdirAll(diagramPath, 0755); err != nil {
			log.Fatalf("failed to create dir: %v", err)
		}
	}

	// readmeDiagram is a readme content for diagram
	readmeDiagram := "\n"

	i := 1
	for _, diagram := range diagrams {
		err := generateDiagram(mainPath, diagramPath, diagram.Name, diagram.Focus, changedFiles)
		if err != nil {
			log.Fatalf("failed to generate diagram: %v", err)
		}

		// creates the list of diagram on the readme
		readmeDiagram += fmt.Sprintf("%d. [%s diagram](docs/diagrams/%s.png)\n", i, diagram.Name, diagram.Name)

		i++
	}

	newReadme, err := generateNewReadme(readmeDiagram)
	if err != nil {
		log.Fatalf("failed to generate new readme: %v", err)
	}

	err = writeNewReadme(newReadme, readmePath)
	if err != nil {
		log.Fatalf("failed to write new readme: %v", err)
	}

	err = writeNewReadme(newGenerateDiagramReadme, generateDiagramReadmePath)
	if err != nil {
		log.Fatalf("failed to new diagram readme: %v", err)
	}

	log.Println("### Generating all diagram finished!! ###")
}

// getChangedFile gets the file changed based on last commit
func getChangedFile() ([]string, string, error) {
	newGenerateDiagramReadme := ""

	// Finds the current commit id
	cmd, err := exec.Command("bash", "-c", "git rev-parse HEAD").Output()
	if err != nil {
		return nil, "", err
	}

	lastCommit := space.ReplaceAllString(string(cmd), "")

	// Reads the content from generate diagram README
	readmeFile, err := os.Open(generateDiagramReadmePath)
	if err != nil {
		return nil, "", err
	}
	defer readmeFile.Close()

	scanner := bufio.NewScanner(readmeFile)

	previousCommit := ""
	for scanner.Scan() {
		text := scanner.Text() + "\n"

		// Checks line by line to find the latest generated version from the commit id
		if strings.Contains(scanner.Text(), "<!-- version:") {
			previousCommit = strings.Split(scanner.Text(), ":")[1]
			text = fmt.Sprintf("<!-- version:%s: -->\n", lastCommit)
		}

		newGenerateDiagramReadme += text
	}

	if previousCommit == "" || lastCommit == "" {
		log.Fatal("previous or last commit not found!")
	}

	var files []string

	fmt.Println(">>> Previous commit: " + previousCommit) // 05a33196589d44a4e67d1d2cd22d29e82efd9650
	fmt.Println(">>> Last commit: " + lastCommit)         // 3093b1e2f74708b9040b717519044c645489e204
	// 27ab8af3a625db4f92e55832b6fd0708c5e7de99
	// Get changed file between previous latest commit id and current commit id
	cmd, err = exec.Command("bash", "-c", "git diff --name-only "+previousCommit+" "+lastCommit).Output()
	if err != nil {

		return nil, "", err
	}

	// Puts the files into array string
	scanner = bufio.NewScanner(strings.NewReader(string(cmd)))
	for scanner.Scan() {
		files = append(files, scanner.Text())
	}

	return files, newGenerateDiagramReadme, err
}

// generateNewReadme generates the content of the new readme (including the old readme + diagram readme)
func generateNewReadme(readmeDiagram string) (string, error) {
	readmeFile, err := os.Open(readmePath)
	if err != nil {
		return "", err
	}
	defer readmeFile.Close()

	scanner := bufio.NewScanner(readmeFile)

	skipLine := false
	newReadme := ""

	for scanner.Scan() {
		if skipLine {
			// continue reads the content if end diagram label found
			if scanner.Text() == "<!-- end diagram doc -->" {
				skipLine = false

				newReadme += scanner.Text() + "\n"
			}

			continue
		}

		newReadme += scanner.Text()

		// start skips the content if start diagram label found
		// and put the readme diagram on the content
		if scanner.Text() == "<!-- start diagram doc -->" {
			skipLine = true

			newReadme += readmeDiagram
		}

		newReadme += "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return newReadme, nil
}

// writeNewReadme writes the new readme file.
func writeNewReadme(newReadme, path string) error {
	newReadmeFile, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}

	defer newReadmeFile.Close()

	_, err = newReadmeFile.WriteString(newReadme)

	if err != nil {
		return err
	}

	return nil
}

// getMainPath gets the main path info
func getMainPath() (string, error) {
	// read the build info to get the module name
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return "", errors.New("failed to read build info")
	}

	return bi.Main.Path, nil
}

// generateDiagram generates the diagram based on the targeted package
func generateDiagram(mainPath, diagramDir, diagramName, focus string, changedFiles []string) error {
	fmt.Println(">>> Running generating " + diagramName + " diagram...")

	doGenerate := false

	// checks if targeted focus is exist on the changed files
	for i := range changedFiles {
		targetFocus := focus

		if focus == "" {
			targetFocus = "cmd/main.go"
		}

		if strings.Contains(changedFiles[i], targetFocus) {
			doGenerate = true
		}
	}

	// skip if the target diagram not exist from changed file
	if !doGenerate {
		fmt.Println(">>> " + diagramName + " diagram not changed, skipped!!")
		return nil
	}

	// output location for the diagram file
	outputFile := "-file " + diagramDir + "/" + diagramName
	// format the output to png
	outputFormat := "-format png"
	// main package in cmd dir
	mainPackage := mainPath + "/cmd"

	focusPackage := ""

	if focus != "" {
		focusPackage = "-focus " + mainPath + "/" + focus
	}

	// wrap a command to generates the diagram using go-callvis [https://github.com/ofabry/go-callvis]
	command := fmt.Sprintf("go-callvis %s %s %s %s",
		outputFile,
		outputFormat,
		focusPackage,
		mainPackage,
	)

	// execute the command
	err := exec.Command("bash", "-c", command).Run()
	if err != nil {
		return err
	}

	// remove unnecessary gv file
	err = os.Remove(diagramDir + "/" + diagramName + ".gv")
	if err != nil {
		return err
	}

	// execute the compress image command using pngquant
	err = exec.Command("bash", "-c", "pngquant "+diagramDir+"/"+diagramName+".png --output "+diagramDir+"/"+diagramName+".png --force").Run()
	if err != nil {
		return err
	}

	fmt.Println(">>> " + diagramName + " diagram generated!!")

	return nil
}
