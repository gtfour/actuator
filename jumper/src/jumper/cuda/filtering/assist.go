package filtering

func Shifter(interval [][]int)(ninterval [][]int) {
    // what this fucking function does ?? ??
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

