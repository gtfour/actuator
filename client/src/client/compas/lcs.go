package compas

func MakeMix(text [][]string)(mix []string){
    mix=make([]string,0)
    for l:= range text {
        line:=text[l]
        for w:= range line{
            word:=line[w]
            mix = append(mix, word)
        }
    }
    return mix
}
