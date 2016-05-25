package main

import "fmt"
import "os"
import "log"
import "text/template"
import "strings"

var DIGIT_LESS_INTERVAL    int = 3579
var DIGIT_GREATER_INTERVAL int = 9753
var DIGIT_IN_INTERVAL      int = 9779

type Tango struct {
    Data []string
}


func main() {

  //new_data    := []int {1,2,3}
  //line        := "a=22 b=1 c=0"
  line        := "наименование='курица' вес=3.5кг процентное содержание воды=40%"
  lineAsArray := strings.Split(line, "")
  data_pos    := [][]int{{14,19},{26,28},{59,60}}
  //fmt.Printf("\n %v   %v\n",new_data,line)

  GenTemplate(lineAsArray, data_pos)
  //data   := []

}

//func GenTemplate(lineAsArray []string, data_pos [][]int)( *template.Template) {
func GenTemplate(lineAsArray []string, data_pos [][]int)() {
    //new_line = ``
    new_data    := []string {"говядина","1.3","30"}

    template_string:=GetFixedArrayChars(lineAsArray, data_pos)

    t:=template.Must(template.New("data_replace").Parse(template_string))

    tango:=&Tango{Data:new_data}

    err := t.Execute(os.Stdout, tango)

    if err != nil {
        log.Println("executing template:", err)
    }
    //fmt.Printf("\n Line without data: %v \n", template_string)
}

func GetFixedArrayChars(lineAsArray []string, data_indexes[][]int) (string) {

    word:=""
    template_variable_counter:=0
    fmt.Printf("\n Line Array %v\n",lineAsArray)

    //for i := range  lineAsArray {
    for i := 0; i < len(lineAsArray); i++ {
    //for i,c := range  lineAsArray {
        //last_pair_index := -1
        //if i == start_at {
        //fmt.Printf("%s %s  %d\n",lineAsArray[i],string(c),i)
        on_interval     := false
        invert          := false
        matched_pair    := []int{-1,-1}
        for z := range data_indexes {
            pair:=data_indexes[z]
            if (len(pair)==2 && pair[0]<=pair[1]) {
                if DigitInInterval(i, pair) == DIGIT_IN_INTERVAL {
                    on_interval = true
                    if len(pair) == 2 {
                        //last_pair_index = pair[1]
                        matched_pair = pair
                    }
                    break
                }
            } else {
                temp_pair := []int{pair[1],pair[0]}
                    if DigitInInterval(i, temp_pair) == DIGIT_IN_INTERVAL {
                        on_interval = true
                        invert      =  true
                        if len(pair) == 2 {
                            //last_pair_index = pair[1]
                            matched_pair = temp_pair
                        }
                        break
                    }
            }
        }


        if (on_interval == false) {
            word=word+lineAsArray[i]
        } else {
            if ( invert==false ) {
                word=word+fmt.Sprintf("{{index .Data %d}}",template_variable_counter)
            } else {
                first_delim:=matched_pair[0]
                last_delim :=matched_pair[1]
                word=word+fmt.Sprintf(lineAsArray[first_delim]+"{{index .Data %d}}"+lineAsArray[last_delim],template_variable_counter)
            }
            template_variable_counter+=1
            i = matched_pair[1]// start iterating array from next index of found pair
        }
    }
    fmt.Printf("\nword %s\n",word)
    return word
}

func DigitInInterval(digit int, interval []int) (int) {
    if digit <= interval[1] && digit >= interval[0] {
        return DIGIT_IN_INTERVAL
    }
    if digit < interval[0] {
        return DIGIT_LESS_INTERVAL
    }
    if digit > interval[1] {
        return DIGIT_GREATER_INTERVAL
    }
    return 0
}

