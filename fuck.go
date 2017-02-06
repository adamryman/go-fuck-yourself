package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	// Dave Cheny
	"github.com/pkg/errors"
)

const maxCalls = 100

var pgMode bool

func init() {
	// pg test
	if os.Args[0] == "goofy" {
		pgMode = true
	}

}

func main() {
	var calls int
	go func() {
		for {
			if calls > 100 {
				fmt.Fprintln(os.Stderr, "May never build...")
				fmt.Fprintln(os.Stderr, `¯\_(ツ)_/¯`)
				os.Exit(1)
			}
			time.Sleep(time.Second)
		}
	}()

	buffer := bytes.NewBuffer(nil)
	var fuck func()
	fuck = func() {
		calls = calls + 1
		cmd := exec.Command("go", os.Args[1:]...)

		cmd.Stderr = buffer
		cmd.Stdout = os.Stdout
		scanner := bufio.NewScanner(buffer)

		err := cmd.Run()
		if err != nil {
			defer fuck()
		}

		linesRemoved := 0
		for scanner.Scan() {
			b := scanner.Bytes()

			file, lineN, err := parseFileLine(b)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			if lineN < 1 {
				continue
			}
			line, err := Delete(file, lineN-linesRemoved)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				linesRemoved = linesRemoved + 1
				if pgMode {
					fmt.Fprintf(os.Stderr, "Lets just not do this: `%s`\n", strings.TrimSpace(line))

				} else {
					fmt.Fprintf(os.Stderr, "FUCK: `%s`\n", strings.TrimSpace(line))
				}
			}
		}
		buffer.Reset()
	}
	fuck()
	fmt.Fprintln(os.Stderr, "You are now a go developer")
}

func parseFileLine(goOut []byte) (string, int, error) {
	fileLine := bytes.Split(goOut, []byte{':'})
	if len(fileLine) < 2 {
		return "", 0, nil
	}
	file := string(fileLine[0])
	lineN, err := strconv.Atoi(string(fileLine[1]))
	if err != nil {
		return "", 0, errors.Wrap(err, "cannot convert line number to int")
	}
	return file, lineN, nil
}

// Delete removes one line from a file by path and returns the line removed
func Delete(path string, line int) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", errors.Wrap(err, "cannot open file to remove line")
	}
	scanner := bufio.NewScanner(f)
	buffer := bytes.NewBuffer(nil)
	for i := 1; i < line; i++ {
		scanner.Scan()
		_, err = buffer.Write(append(scanner.Bytes(), byte('\n')))
		if err != nil {
			return "", errors.Wrap(err, "cannot write to buffer")
		}
	}
	// DO IT
	scanner.Scan()
	fuck := scanner.Text()
	// SEE YA
	for scanner.Scan() {
		_, err = buffer.Write(append(scanner.Bytes(), byte('\n')))
		if err != nil {
			return "", errors.Wrap(err, "cannot write to buffer")
		}
	}
	err = f.Close()
	if err != nil {
		return "", errors.Wrap(err, "cannot close old file")
	}
	// LETS DO IT
	f, err = os.Create(path)
	if err != nil {
		return "", errors.Wrap(err, "cannot overwrite old file with new file")
	}
	_, err = io.Copy(f, buffer)
	if err != nil {
		return "", errors.Wrap(err, "cannot write to new file")
	}

	return fuck, nil
}
