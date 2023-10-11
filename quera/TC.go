package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFilePath := flag.String("f", "", "path to the go file")
	outputFilePath := flag.String("o", "", "path to the output file")
	flag.Parse()

	modifiedCode := manipulateFile(*inputFilePath)

}

func commentCode(filePath, startTag, endTag string) {
	file, _ := os.Open(filePath)
	defer file.Close()

	var newLines []string
	commentBlock := false
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == endTag {
			commentBlock = false
			continue
		}

		if strings.TrimSpace(line) == startTag {
			commentBlock = true
			continue
		}

		if commentBlock {
			newLines = append(newLines, "//"+line)
		} else {
			newLines = append(newLines, line)
		}
	}
	fmt.Println(newLines)
}

func uncommentCode(filePath, startTag, endTag string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	var newLines []string
	uncommentBlock := false
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == endTag {
			uncommentBlock = false
			continue
		}

		if strings.TrimSpace(line) == startTag {
			uncommentBlock = true
			continue
		}

		if uncommentBlock {
			for i := 0; i < len(line); i++ {
				if strings.HasPrefix(line[i:], "//") {
					line = line[i+2:]
					i = -1
				}
			}
		}
		newLines = append(newLines, line)
	}
	fmt.Println(newLines)
}

func deleteCommentCode(filePath, startTag, endTag string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	var newLines []string
	deleteCommentBlock := false

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == startTag {
			deleteCommentBlock = true
		}

		if !deleteCommentBlock {
			newLines = append(newLines, line)
		}

		if strings.TrimSpace(line) == endTag {
			deleteCommentBlock = false
		}
	}
	fmt.Println(newLines)
}

func renameCommentCode(filePath, renameTag string) {
	file, _ := os.Open(filePath)
	defer file.Close()

	var buffer bytes.Buffer
	var oldName string
	var newName string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, renameTag) {
			parts := strings.Split(line, " ")
			if len(parts) >= 4 {
				oldName = parts[3]
				newName = parts[4]
			}
			continue
		}
		buffer.WriteString(line + "\n")
	}

	completeCode := buffer.String()
	completeCode = strings.ReplaceAll(completeCode, oldName, newName)
	fmt.Println(completeCode)
}

func manipulateFile(inputFile string) string {
	file, _ := os.Open(inputFile)
	defer file.Close()

	var buffer bytes.Buffer
	var oldName string
	var newName string
	commentBlock, uncommentBlock, deleteBlock := false, false, false

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		switch strings.TrimSpace(line) {
		case "// TC end_comment":
			commentBlock = false
			continue
		case "// TC start_comment":
			commentBlock = true
			continue
		case "// TC end_uncomment":
			uncommentBlock = false
			continue
		case "// TC start_uncomment":
			uncommentBlock = true
			continue
		case "// TC end_delete":
			deleteBlock = false
			continue
		case "// TC start_delete":
			deleteBlock = true
			continue
		}

		if commentBlock {
			line = "// " + line
		}

		if uncommentBlock {
			for i := 0; i < len(line); i++ {
				if strings.HasPrefix(line[i:], "//") {
					line = line[i+2:]
					i = -1
				}
			}
		}

		if deleteBlock {
			continue
		}

		if strings.Contains(line, "// TC rename") {
			parts := strings.Split(line, " ")
			if len(parts) >= 4 {
				oldName = parts[3]
				newName = parts[4]
			}
			continue
		}
		buffer.WriteString(line + "\n")
	}

	completeCode := buffer.String()
	completeCode = strings.ReplaceAll(completeCode, oldName, newName)

	return completeCode
}
