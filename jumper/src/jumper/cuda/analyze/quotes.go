package analyze

import "strings"

func GetQuotesIndexes (lineAsArray []string) ( indexes []int) {

    var quotes                    =  [3]string {`"`, "'", "`"}

    for char:= range lineAsArray {
        for q:=range quotes {
            quote:=quotes[q]
            if lineAsArray[char] == quote {
                indexes = append(indexes, char)
            }
        }
    }
    return
}


func GroupByQuotes (lineAsArray []string) (quotes_pairs [][]int) {

    quotes_indexes := GetQuotesIndexes(lineAsArray)

    pending := make(map[string]int)
    for char:= range quotes_indexes {
        quote:=quotes_indexes[char]
        quoteInPending:=false
        quote_value:=lineAsArray[quote]
        for key, _ := range pending { if quote_value == key { quoteInPending = true } }
        if quoteInPending == false {
            pending[quote_value] = quote
        } else {
            var new_pair = []int  { pending[quote_value],quote } // 1 added for ignoring quotes  themselves
            quotes_pairs = append(quotes_pairs, new_pair)
            delete(pending, quote_value)
        }
    }

    return
}

func QuotesSpreading ( entry string) ( word_set [3]string, complete [3]bool  ) {

    var quotes                    =  [3]string {`"`, "'", "`"}

    for i:=range quotes {
        quote:=quotes[i]
        if strings.Count(entry, quote)%2 == 0 {
            complete[i] = true
        } else {
            complete[i] = false
        }
        word_set[i] = strings.Replace(entry, quote, "", -888)
    }
    return word_set, complete
}
