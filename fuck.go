package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/adamryman/kit/file/line"
	// Dave Cheny
	"github.com/pkg/errors"
)

const maxCalls = 100

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

	var fuck func()
	fuck = func() {
		calls = calls + 1
		cmd := exec.Command("go", os.Args[1:]...)

		buffer := bytes.NewBuffer(nil)
		writer := io.MultiWriter(buffer, os.Stderr)
		cmd.Stderr = writer
		cmd.Stdout = os.Stdout
		scanner := bufio.NewScanner(buffer)
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			defer fuck()
		}

		linesRemoved := 0
		for scanner.Scan() {
			b := scanner.Bytes()

			file, lineN, err := parseFileLine(b)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			if lineN < 1 {
				continue
			}
			err = lineDeleteWrapper(file, lineN-linesRemoved)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				linesRemoved = linesRemoved + 1
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

func lineDeleteWrapper(path string, lineN int) error {
	err := line.Delete(path, lineN)
	if err != nil {
		return err
	}
	return nil
}
