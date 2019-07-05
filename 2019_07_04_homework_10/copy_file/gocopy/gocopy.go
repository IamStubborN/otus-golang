package gocopy

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func WriteFile(from, to string, offset, limit int64) (int, error) {

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

	buffer := make([]byte, limit)

	_, err = fileSource.ReadAt(buffer, offset)
	if err != nil {
		return 0, err
	}
	written, err := write(fileSource, fileDestination, buffer)
	if err != nil {
		log.Fatal(err)
	}
	return written, nil
}

func write(src io.ReaderAt, dst io.WriterAt, buffer []byte) (int, error) {

	off := 0
	step := 8
	totalBytes := 0
	length := len(buffer)
	for totalBytes < length {
		b, err := dst.WriteAt(buffer[totalBytes:off], int64(totalBytes))
		if err != nil {
			return 0, err
		}

		if off+step > length {
			off = length
		} else {
			off += step
		}

		totalBytes += b

		prevPercent := (totalBytes - step) * 100 / length
		currentPercent := (totalBytes) * 100 / length
		if prevPercent != currentPercent {
			progressWrite(os.Stdout, currentPercent)
		}
		time.Sleep(3 * time.Millisecond) // Для наглядности
	}
	return totalBytes, nil
}
func progressWrite(w io.Writer, percent int) {
	progress := strings.Repeat("░", percent/5)
	fmt.Fprint(w, fmt.Sprintf("\r 0  %s %d", progress, percent))
	if percent == 100 {
		fmt.Fprintln(w)
	}
}

func getLengthOfFileInBytes(file *os.File) (int64, error) {
	fi, err := file.Stat()
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}
