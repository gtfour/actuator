package a

var A = 1

func GetA()(int) {
    return A
}

type Test struct {
    A int
    a int
}

func CreateTest()(t Test) {
    t.A = 1
    t.a = 2
    return t
}
