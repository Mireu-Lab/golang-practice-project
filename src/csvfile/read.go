package csvfile

import (
	"bufio"
	"encoding/csv"
	"os"
)

func Read(FilePath string) [][]string {
	file, _ := os.Open(FilePath)

	rdr := csv.NewReader(bufio.NewReader(file)) // csv reader 생성
	rows, _ := rdr.ReadAll()                    // csv 내용 모두 읽기

	return rows
}
