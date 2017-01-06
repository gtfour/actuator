package cross

type Database interface {
    Connect()(error)
    RunQuery(*Query)(result_slice_addr *[]map[string]interface{},err error)
    Close()()
}
