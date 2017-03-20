package out

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
)

// utility function to convert a slice of int64 into a slice of bytes.
// return (nil, error) if slice is nil or if writing int64 in the buffer rise
// an error.
func convert(slice []int64) ([]byte, error) {
	if slice == nil {
		return nil, errors.New("Slice given can not be null")
	}

	// make a buffer of  bytes
	buf := new(bytes.Buffer)
	// for every int64 into slice, convert it in byte and push it in the buffer
	for _, e := range slice {
		if err := binary.Write(buf, binary.LittleEndian, e); err != nil {
			return nil, err
		}
	}

	// Bytes() function return a []byte
	return buf.Bytes(), nil
}

// utility function to check if open file given is a directory or not.
func isDirectory(file *os.File) (error, bool) {
	fileStat, errs := os.Stat(file)
	if errs != nil {
		return nil, errors.New("Error while retriving stats from file")
	}
	return nil, fileStat.IsDir()
}

// utilty function to check if given file is readable or writable.
// nil, err is occurred otherwise.
func openIfCanRW(file string) (*os.File, error) {
	if file == nil {
		return nil, errors.New("Given string cao not be nil.")
	}

	// check if file is readable or writable.
	openFile, err := os.OpenFile(file, os.O_RDWR, 0666)
	if err != nil {
		if os.IsPermission(err) {
			return nil, err
		}
	}

	// check fi given path points to directory
	if err, isDir := isDirectory(openFile); err != nil {
		return nil, err
	} else if isDir {
		return nil, errors.New("Given path is a directory")
	}

	return openFile, nil
}

// this utiliy function writes the header of the file according to structure
// given. The headers for each data structure are described in README.
func writeHeader(structure interface{}, path string) error {
}

func WriteA(slice []int64, path string) error {
	return nil
}

// TODO add doc
func Write(slice []int64, path string) error {
	// TODO check args
	// TODO !!! check if file is not a binary file || directory || permissions

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// create and write file header
	s := fmt.Sprintf("%d\n", len(slice))
	_, err = file.WriteString(s)
	if err != nil {
		return err
	}

	// create and write file body
	err = writeSingleSlice(slice, file)
	if err != nil {
		return err
	}

	// Issue a Sync to flush writes to stable storage
	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}

// reusable method to write a single slice on given file
func writeSingleSlice(slice []int64, file *os.File) error {
	var s string
	for _, v := range slice {
		s = fmt.Sprintf("%d ", v)
		_, err := file.WriteString(s)
		if err != nil {
			return err
		}
	}
	_, err := file.WriteString("\n")
	return err
}
