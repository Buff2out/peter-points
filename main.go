package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// можно заморочиться с буферизацией,
// но только если будет задание с бОльшим числом записей
// а пока можно абстрагироваться и оперативной памяти должно быть "за глаза"
// с (пока что) таким кол-вом записей
func getFromTsv(filename string) {
	res := make([][]string, 20)
	f, _ := os.Open(filename)
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = '\t'

	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	records = records[1:] // reslice to omit header
	for i, record := range records {
		res[i] = append(res[i], record[0])                    // name
		res[i] = append(res[i], record[1][:len(record[1])-2]) // убираем "ч" с конца 2 == 1 - кириллица
		res[i] = append(res[i], record[2])                    // ранг они же очки (приоритет)
		fmt.Println(res[i][0], res[i][1], res[i][2])
	}
}

func main() {
	getFromTsv("test.tsv")
}
