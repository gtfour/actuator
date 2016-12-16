package custom

import "jumper/cuda"


//var CustomFilters = CreateCustomFilterList()

cuda.GetCustonFilters = CreateCustomFilterList



func CreateCustomFilterList ()(cuda.FilterList) {

    fl := make(cuda.FilterList,0)
    fl.Append(shitty_filter)
    return fl

}


// You can add here custom parser's

