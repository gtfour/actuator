package maps

var TRUE_SLICE  int =  9001
var FALSE_SLICE int =  9000
var BUNT_SLICE  int =  9002
var EMPTY_SLICE int =  9004

func CheckBoolSlice(slice []bool)(slice_type int){
    true_values  := make([]bool,0)
    false_values := make([]bool,0)
    for i:= range slice {
        value := slice[i]
        if value {
            true_values=append(true_values,value)
        } else {
            false_values=append(false_values,value)
        }
    }
    if len(true_values)>0 && len(false_values)>0 {
        return BUNT_SLICE
    } else if len(true_values)>0 && len(false_values)==0 {
        return TRUE_SLICE
    } else if len(true_values)==0 && len(false_values)>0 {
        return FALSE_SLICE
    } else {
        return EMPTY_SLICE
    }
}

func CompareMap(query map[string]interface{}, dest map[string]interface{})(bool) {

    matching := make([]bool,0)
    for key,value := range query {
        if dest_value,ok := dest[key]; ok == true {
            if dest_value == value {
                matching = append(matching, true)
            } else {
                matching = append(matching, false)
            }
        }
    }
    if CheckBoolSlice(matching) == TRUE_SLICE {
        return true
    } else {
        return false
    }
}

func UpdateMap(new_map map[string]interface{}, source_map map[string]interface{})(map[string]interface{}, error) {

    if new_map == nil { return nil, InputMapIsEmpty }
    if source_map == nil { source_map=make(map[string]interface{},0) }
    for key,value := range new_map {
        source_map[key]=value
    }
    return source_map, nil
}
