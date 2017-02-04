package fuck

import (
	"bufio"
	"bytes"
	"io"
	"os"

	// DAVE CHENY
	"github.com/pkg/errors"
)

func DeleteLine(path string, line int) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "cannot open incorrect file")
	}
	scanner := bufio.NewScanner(f)
	buffer := bytes.NewBuffer(nil)
	for i := 0; i < line; i++ {
		scanner.Scan()
		scanner.Bytes()
		_, err := buffer.Write(scanner.Bytes())
		if err != nil {
			return errors.Wrap(err, "cannot write to buffer")
		}
	}
	// DO IT
	scanner.Scan()
	// SEE YA
	for scanner.Scan() {
		_, err := buffer.Write(scanner.Bytes())
		if err != nil {
			return errors.Wrap(err, "cannot write to buffer")
		}
	}
	err = f.Close()
	if err != nil {
		return errors.Wrap(err, "cannot close incorrect file")
	}
	// LETS DO IT
	f, err = os.Create(path)
	_, err = io.Copy(f, buffer)
	if err != nil {
		return errors.Wrap(err, "cannot overwrite file with correct one")
	}

	return nil
}
