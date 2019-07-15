package gocopy

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// WriteFile from path to path with offset and limit in bytes.
func WriteFile(from, to string, offset, limit int64) (int64, error) {

	fileSource, err := os.OpenFile(from, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fileSource.Close()

	fileDestination, err := os.OpenFile(to, os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer fileDestination.Close()

	length, err := getLengthOfFileInBytes(fileSource)
	if err != nil {
		return 0, err
	}

	if offset >= length {
		return 0, fmt.Errorf("offset %d out of file with size %d", offset, length)
	}

	if limit > length || limit <= 0 {
		return 0, fmt.Errorf("limit %d out of file with size %d", limit, length)
	}
	if limit+offset > length {
		return 0, fmt.Errorf("limit %d and offset %d out of file with size %d", limit, offset, length)
	}

	written, err := write(fileSource, fileDestination, offset, limit)
	if err != nil {
		log.Fatal(err)
	}
	return written, nil
}

func write(fileSource, fileDestination *os.File, offset, limit int64) (int64, error) {
	var step int64 = 10
	var totalBytes int64 = 0

	_, err := fileSource.Seek(offset, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	for totalBytes < limit {

		if totalBytes+step > limit {
			step = limit - totalBytes
		}

		b, err := io.CopyN(fileDestination, fileSource, step)
		if err != nil {
			return 0, err
		}

		totalBytes += b

		prevPercent := (totalBytes - step) * 100 / limit
		currentPercent := (totalBytes * 100)/ limit
		if prevPercent != currentPercent {
			progressWrite(os.Stdout, currentPercent)
		}
	}
	return totalBytes, nil
}
func progressWrite(w io.Writer, percent int64) {
	progress := strings.Repeat("░", int(percent)/5)
	_, err := fmt.Fprint(w, fmt.Sprintf("\r 0  %s %d", progress, percent))
	if err != nil {
		log.Fatal(err)
	}
	if percent == 100 {
		_, err = fmt.Fprintln(w)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getLengthOfFileInBytes(file *os.File) (int64, error) {
	fi, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}
