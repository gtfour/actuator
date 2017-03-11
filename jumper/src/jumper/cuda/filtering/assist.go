package filtering

import "fmt"
import "jumper/cuda/commons"

var LEFT_DIRECTION         int = 1100
var RIGHT_DIRECTION        int = 1001
var DIGIT_LESS_INTERVAL    int = 3579
var DIGIT_GREATER_INTERVAL int = 9753
var DIGIT_IN_INTERVAL      int = 9779

func Shifter(interval [][]int)(ninterval [][]int) {
    //
    // what this fucking function does ?? ??
    // as i see it  skips nested intervals (intervals which is smaller than another existing interval) 
    //
    var skipped []int
    for i:= range interval {
        if commons.IsDigitIn(i,skipped) == false {
            parent_int_part:=interval[i]
            for z:= range interval {
                if z == i {continue} //do not compare interval with itself
                int_part:=interval[z]
                if len(int_part)!=2{continue}

                first:=int_part[0]
                last :=int_part[1]

                if commons.DigitInInterval(first,parent_int_part) == DIGIT_IN_INTERVAL && commons.DigitInInterval(last,parent_int_part) == DIGIT_IN_INTERVAL {
                    skipped=append(skipped, z)
                }

            }

        }
    }
    for x:= range interval {
        if commons.IsDigitIn(x,skipped) == false {
            ninterval=append(ninterval, interval[x])
        }
    }
    return
    //
    //
    //
}

func AlumaPaster(delims [][]int, data [][]int, strada [][]int)(ndelims [][]int, ndata [][]int){
    //
    //  strada should be inserted in data array
    //  delims with indexes included in strada will be ignored
    //  data  with indexes included  in strada will be ignored
    //  fmt.Printf("  delims: %v\n  data: %v\n strada: %v\n",delims,data,strada)
    //
    // 
    // hello
    //  kind of another strange and magic function :) 
    //  i guess i don't know what this function does again
    //
    var last_delim_index int
    var last_data_index  int
    //
    delims_last_elem := delims[(len(delims)-1)]
    data_last_elem   := data[(len(data)-1)]
    if len(delims_last_elem)==2 && len(data_last_elem)==2 {
        last_delim_index = delims_last_elem[1]
        last_data_index  = data_last_elem[1]
    }
    // --
    // -- range over strada
    // --
    for i := range strada {
        // --
        // -- delims array and data array will be updated inside each strada iteration 
        // -- 
        ndelims := [][]int{}
        indexes := strada[i]
        if len(indexes)!=2 { continue }  //{ break ; return delims, data }
        first := indexes[0] // first strada elem
        last  := indexes[1] // last  strada elem
        //
        //
        //
        if first > last {  // // ?? wut  ... so seems i still remeber why: this trick need's when strada describes an emty position  : ""
                           // // example     : ""
                           // // description : strada for emptyness between two quotes will be described by reversed indexes 
                           // // quotes indexes(quotes will be determined as delims ): 5 and 6 
                           // // strada for above example( strada describes data indexes. In this case data is empty  ): [6,5]
            first = indexes[1]
            last  = indexes[0]
        }
        // -- 
        // -- range over delims . 
        // --
        for de := range  delims {
            delim        := delims[de]
            if len(delim)!=2 { continue }
            first_delim       := delim[0]
            last_delim        := delim[1]
            first_state       := commons.DigitInInterval(first, delim)
            last_state        := commons.DigitInInterval(last, delim)
            first_delim_state := commons.DigitInInterval(first_delim, indexes)
            last_delim_state  := commons.DigitInInterval(last_delim,  indexes)
            //fmt.Printf("\nfirst %v | firststate %v | laststate %v | strada %v | delim %v | firstdelimstate %v | lastdelimstate %v \n ",first,first_state,last_state, strada,delim,first_delim_state, last_delim_state)
            //
            // comparing delims array with current strada
            //
            if first_state == DIGIT_IN_INTERVAL && last_state == DIGIT_IN_INTERVAL {
               // split current delim to two new delims without strada indexes
               fmt.Printf("\nStrada on delim interval\n")
               new_delim_first := make([]int, 2)
               new_delim_last  := make([]int, 2)
               diff_first      := first - first_delim
               diff_last       := last_delim  - last
               if diff_first>0 { new_delim_first[0] = first_delim ; new_delim_first[1] = first - 1 ; ndelims=append(ndelims, new_delim_first) }
               if diff_last >0 { new_delim_last[0]  = last +1     ; new_delim_last[1]  = last_delim; ndelims=append(ndelims, new_delim_last)  }
               // if diff_first == 0 && diff_last == 0 {   }
            } else if first_state == DIGIT_IN_INTERVAL {
                new_delim    := make([]int,2)
                diff_first   := first - first_delim
                if diff_first > 0{
                    new_delim[0]= first_delim
                    new_delim[1]= first-1
                    ndelims=append(ndelims, new_delim)
                }
            } else if last_state == DIGIT_IN_INTERVAL {
                new_delim    := make([]int,2)
                diff_last    := last_delim - last
                if diff_last > 0 {
                    new_delim[0]= last +1
                    new_delim[1]= last_delim
                    ndelims=append(ndelims, new_delim)
                }
            //
            //  ???  Pay attention
            //
            } else if first_delim_state == DIGIT_IN_INTERVAL && last_delim_state == DIGIT_IN_INTERVAL {
                //
                //
                //

            } else {
                //
                ndelims=append(ndelims, delim)
                //
            }
        }
        //
        // stop changing existing delims
        //
        delims = ndelims // replace existing delims set with new delims set . new delims set has been updated according to strada  
    }
    ndelims                =  delims     // wtf  ? unfortunately  forgot why needs this reverse :(  
    last_matched_strada_id := -1         // what ? 

    for da := range  data {
        data_part:=data[da]
        if len(data_part)!=2 { continue }
        first_data := data_part[0]
        last_data  := data_part[1]
        var includes      bool
        var insert_strada bool
        for i := range strada {
            indexes:=strada[i]
            if len(indexes)!=2 { continue }
            first_state          := commons.DigitInInterval(first_data, indexes)
            last_state           := commons.DigitInInterval(last_data, indexes)
            if first_state == DIGIT_IN_INTERVAL && last_state == DIGIT_IN_INTERVAL{
                includes = true
                if i != last_matched_strada_id {
                    insert_strada          = true
                    last_matched_strada_id = i
                }
            } else {
                // // fmt.Printf("\n::else::\n")
                interval_between_data:=make([]int,2)
                first_strada         := indexes[0]
                last_strada          := indexes[1]
                var replace_on_insert bool
                if first_strada > last_strada {
                    first_strada = indexes[1]
                    last_strada  = indexes[0]
                    replace_on_insert = true
                }
                //last_delim_index := delims[(len(delims)-1)][1]
                //last_data_index  := data[(len(data)-1)][1]
                if da == 0 {
                    interval_between_data[0] = 0
                    if len(data) == 1 {
                        if last_delim_index >= last_data_index {
                            interval_between_data[1] = last_delim_index
                        } else {
                            interval_between_data[1] = last_data_index
                        }
                    } else {
                        interval_between_data[1] = last_data

                    }
                } else if da == len(data)-1 {

                    interval_between_data[0] = first_data
                    if last_delim_index >= last_data_index {
                        interval_between_data[1]  = last_delim_index
                    } else {
                         interval_between_data[1] = last_data_index
                    }

                } else {
                    interval_between_data[0] = data_part[1]
                    interval_between_data[1] = data[da+1][0]
                }
                if  commons.DigitInInterval(first_strada, interval_between_data)  == DIGIT_IN_INTERVAL && commons.DigitInInterval(last_strada, interval_between_data)  == DIGIT_IN_INTERVAL {
                    last_matched_strada_id = i
                    if replace_on_insert == false {
                        insert_strada = true
                    } else {
                         nindexes:=make([]int,2)
                         nindexes[0] = last_strada
                         nindexes[1] = first_strada
                         insert_strada = true
                    }
                }
                //var interval_between_data []int
                //interval_between_data[0] 


            }
        }
        if includes == false {
            ndata=append(ndata, data_part)
        }
        if insert_strada == true {
            ndata=append(ndata, strada[last_matched_strada_id])
        }
    }
    return //ndelims,ndata
}

