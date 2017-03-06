package templating

import "fmt"
import "jumper/cuda"

func GenTemplate(lineAsArray []string, data_indexes[][]int) (string) {

    word:=""
    template_variable_counter:=0
    //fmt.Printf("\n Line Array %v\n",lineAsArray)

    //for i := range  lineAsArray {
    for i := 0; i < len(lineAsArray); i++ {
    //for i,c := range  lineAsArray {
        //last_pair_index := -1
        //if i == start_at {
        //fmt.Printf("%s %d\n",lineAsArray[i],i)
        on_interval     := false
        invert          := false
        matched_pair    := []int{-1,-1}
        for z := range data_indexes {
            pair:=data_indexes[z]
            if (len(pair)==2 && pair[0]<=pair[1]) {
                if cuda.DigitInInterval(i, pair) == cuda.DIGIT_IN_INTERVAL {
                    on_interval = true
                    if len(pair) == 2 {
                        //last_pair_index = pair[1]
                        matched_pair = pair
                    }
                    break
                }
            } else {
                temp_pair := []int{pair[1],pair[0]}
                    if cuda.DigitInInterval(i, temp_pair) == cuda.DIGIT_IN_INTERVAL {
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
    //fmt.Printf("\nword %s\n",word)
    return word
}

