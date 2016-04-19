package main

import "fmt"
import "sync"

type MutexArray struct {

    sync.Mutex
    array []int


}

func (m *MutexArray) Set(value []int) {
    m.array = value
}

func (m *MutexArray) Get()([]int) {
    return m.array
}

func (m *MutexArray) AddItem(value int) {
    m.Lock()
    m.array = append(m.array, value)
    m.Unlock()
}

func (m *MutexArray) RemoveItem(index int) {
    m.Lock()
    m.array = append(m.array[:0], m.array[0+index:]...)
    m.Unlock()
}


func main() {

    var my_int_array = MutexArray{array:[]int {1,23,35,41,5,6}}
    my_int_array.Lock()
    for i:= range my_int_array.array {
        if i == 2 { my_int_array.RemoveItem(i) }
        fmt.Printf("\n%d %d\n",i,my_int_array.array[i])
    }

}
