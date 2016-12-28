package cross

type Query struct {

    Type      int
    Table     string
    KeyBody   map[string]interface{}
    QueryBody map[string]interface{}

}

type Database interface {

    Connect()(error)
    RunQuery(result_slice_addr *[]map[string]interface{},err error)
    Close()()

}
