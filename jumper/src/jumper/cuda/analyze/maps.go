package analyze

func GetKeyByValue(signs map[int]string, string_value string) (key int) {
    for key, value :=range signs {
        if value == string_value {
            return key
        }
    }
    return -1
}

func ValueExists(signs map[int]string,value string)(found bool ) {
    values:=GetMapValues(signs)
    for i := range values {
        if values[i]==value {
            found=true
        }
    }
    return found
}

func GetMapValues(signs map[int]string)(values []string ){
    for _, value := range signs {
        values=append(values, value)
    }
    return values
}

