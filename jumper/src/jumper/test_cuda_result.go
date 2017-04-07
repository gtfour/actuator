package main

import "fmt"
import "jumper/cuda/result"

func main(){
    rline := result.BlankResult( result.RESULT_TYPE_LINE )
    fmt.Printf("\nLine result: %v  %v\n", rline.GetType(), result.RESULT_TYPE_LINE)
}
