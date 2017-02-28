package arrays

func Extend(base [][]string , part [][]string)([][]string, error){
    for i:= range part {
        piece:=part[i]
        base = append(base, piece)
    }
    return base, nil
}
