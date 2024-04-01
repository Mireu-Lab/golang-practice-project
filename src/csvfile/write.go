package csvfile

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func FileColLen(FilePath string) (int, error) {
	file, err := os.Open(FilePath) // 파일 오픈
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}

	rdr := csv.NewReader(bufio.NewReader(file)) // csv reader 생성

	rows, err := rdr.ReadAll() // csv 내용 모두 읽기
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}

	return len(rows), nil
}

func Write(FilePath string, column []string) error {
	file, err := os.OpenFile(FilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	defer file.Close()

	var data [][]string
	ColLen, _ := FileColLen(FilePath)
	column = append(column, strconv.Itoa(ColLen))
	data = append(data, column)

	w := csv.NewWriter(file)
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}
