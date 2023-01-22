package main

import(
	"os"
	"bufio"
	"github.com/jdkato/prose/v2"
	"fmt"
	"encoding/csv"
)
func ReadBook(path string) ([]string, error) {
	book, err := os.Open(path) 
	var lines []string
	if err != nil {
		return lines, err
	}

	scn := bufio.NewScanner(book)
	for scn.Scan(){
		lines = append(lines, scn.Text())
	}
	book.Close()
	return lines, err
}



func ClassifySetting(text string, settings map[string]Topic, line int){
	doc, _ := prose.NewDocument(text)
    for _, ent := range doc.Entities() {
		fmt.Println(ent.Text)
        if(ent.Label == "GPE"){
			_, contains := settings[ent.Text]
			switch contains{
			case true: 
				s := settings[ent.Text]
				s.mentions += 1
				s.line_num = append(s.line_num, line)
				settings[ent.Text] = s
			
			case false:
				s := Topic{ent.Text, 1, []int{line}}

				settings[s.name] = s
			}

    	}
	}
}

func ReadFoods(path string) (map[string]bool, error){
	datafile,err := os.Open(path)
	if err != nil{
		return make(map[string]bool), err
	}
	csvReader := csv.NewReader(datafile)
	
    data, err := csvReader.ReadAll()
	var all_foods = make(map[string]bool)
	for _,fd := range data {
		all_foods[fd[0]] = true
	}
	return all_foods, err

}


func ClassifyDining(text string, foods map[string]Topic, line int){

}

type Topic struct {
	name string 
	mentions int 
	line_num []int
}

