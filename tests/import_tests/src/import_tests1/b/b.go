package b
import "import_tests/a"

var Test = func ()() {
    a.A=a.A+1
}

func Increase ()() {
    a.A += 1
}
