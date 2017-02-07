package filtering

import . "jumper/cuda"

func Shifter(interval [][]int)(ninterval [][]int) {
    //
    // what this fucking function does ?? ??
    // as i see it  skips nested intervals (intervals which is smaller than another existing interval) 
    //
    var skipped []int
    for i:= range interval {
        if IsDigitIn(i,skipped) == false {
            parent_int_part:=interval[i]
            for z:= range interval {
                if z == i {continue} //do not compare interval with itself
                int_part:=interval[z]
                if len(int_part)!=2{continue}

                first:=int_part[0]
                last :=int_part[1]

                if DigitInInterval(first,parent_int_part) == DIGIT_IN_INTERVAL && DigitInInterval(last,parent_int_part) == DIGIT_IN_INTERVAL {
                    skipped=append(skipped, z)
                }

            }

        }
    }
    for x:= range interval {
        if IsDigitIn(x,skipped) == false {
            ninterval=append(ninterval, interval[x])
        }
    }
    return
}

func AlumaPaster (delims [][]int, data [][]int, strada [][]int) (ndelims [][]int, ndata [][]int) {
    // strada should be inserted in data array
    // delims with indexes included in strada will be ignored
    // data  with indexes included  in strada will be ignored
    //fmt.Printf("  delims: %v\n  data: %v\n strada: %v\n",delims,data,strada)
    // 
    // kind of another strange and magic function :) 
    // i guess i don't know what this function does again
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
    for i := range strada {
        ndelims := [][]int{}
        indexes := strada[i]
        if len(indexes)!=2 { continue }  //{ break ; return delims, data }
        first := indexes[0]
        last  := indexes[1]
        //
        if first > last {
            first = indexes[1]
            last  = indexes[0]
        }
        //
        for de := range  delims {
            delim        := delims[de]
            if len(delim)!=2 { continue }
            first_delim  := delim[0]
            last_delim   := delim[1]
            first_state       := DigitInInterval(first, delim)
            last_state        := DigitInInterval(last, delim)
            first_delim_state := DigitInInterval(first_delim, indexes)
            last_delim_state  := DigitInInterval(last_delim,  indexes)
            //fmt.Printf("\nfirst %v | firststate %v | laststate %v | strada %v | delim %v | firstdelimstate %v | lastdelimstate %v \n ",first,first_state,last_state, strada,delim,first_delim_state, last_delim_state)
            if first_state == DIGIT_IN_INTERVAL && last_state == DIGIT_IN_INTERVAL {
               // split current delim to two new delims without strada indexes
               fmt.Printf("\nStrada on delim interval\n")
               new_delim_first := make([]int,2)
               new_delim_last  := make([]int,2)
               diff_first := first - first_delim
               diff_last  := last_delim  - last
               if diff_first>0 { new_delim_first[0] = first_delim ; new_delim_first[1] = first - 1 ; ndelims=append(ndelims, new_delim_first) }
               if diff_last >0 { new_delim_last[0]  = last +1     ; new_delim_last[1]  = last_delim; ndelims=append(ndelims, new_delim_last)  }
               //if diff_first == 0 && diff_last == 0 {   }
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
            // ??? Pay attention
            } else if first_delim_state == DIGIT_IN_INTERVAL && last_delim_state == DIGIT_IN_INTERVAL {

            } else {
                ndelims=append(ndelims, delim)
            }
        }
        delims = ndelims
    }
    ndelims = delims

    last_matched_strada_id:=-1
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
            first_state          := DigitInInterval(first_data, indexes)
            last_state           := DigitInInterval(last_data, indexes)
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
                if  DigitInInterval(first_strada, interval_between_data)  == DIGIT_IN_INTERVAL && DigitInInterval(last_strada, interval_between_data)  == DIGIT_IN_INTERVAL {
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

