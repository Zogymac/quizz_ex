package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type QuestionReader interface {
	ParseQuestions(r io.Reader) ([]Question, error)
}
/*
type CsvReader struct {

}

func (cr *CsvReader) ParseQuestions(r io.Reader) ([]Question, error) {

}*/

type Question struct {
	question string
	answer   string
}

func readCsv(filename string) []Question {
	csvfile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	//fmt.Println(records[0])
	// Допишите код здесь
	m := make([]Question, len(records))
	for i, v := range records {
		m[i].question = v[0]
		m[i].answer = v[1]
	}
	return m
}

func main() {
	total := 0

	//questions := []Question{{"1 + 1", "2"}, {"2 + 2", "4"}}
	questions := readCsv("problems.csv")
	// Пройтись циклом. Вывести вопрос, предложить пользователю ввести ответ.
	// Если ответ правильный, увеличить total.
	// for
	ans := ""
	for _, v := range questions {
		fmt.Println(v.question)
		fmt.Scanln(&ans)
		if ans == v.answer {
			total += 1
		}
	}
	fmt.Println("You got ", total, "/", len(questions))
}
