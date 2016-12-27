package cross

func isOk(value int, valid []int)(isok bool){
    for i:=range valid { if value==valid[i] { isok=true ; break } }
    return
}
