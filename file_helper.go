package main

import(
	"os"
	"bufio"
	"github.com/jdkato/prose/v2"
	"fmt"
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

	return lines, err
}



func ClassifySetting(text string, settings map[string]Setting, line int){
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
				s := Setting{ent.Text, 1, []int{line}}

				settings[s.name] = s
			}

    	}
	}
}

type Setting struct {
	name string 
	mentions int 
	line_num []int
}
	