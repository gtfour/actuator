package cross

type Query struct {

    Type      int
    Table     string
    KeyBody   map[string]interface{}
    QueryBody map[string]interface{}

}

func(q *Query)ValidateBodies()(match_by_key bool,match_by_value bool,err error){

    if q.KeyBody != nil || q.QueryBody == nil {
        match_by_key   = true
    }
    if q.KeyBody == nil || q.QueryBody != nil {
        match_by_value = true
    }
    if match_by_key == false && match_by_value == false {
        return false,false,KeyAndValueEmpty
    }
    err=nil
    return
}

func(q *Query)IsTableQuery()(bool){
    q.Type

}

type Database interface {

    Connect()(error)
    RunQuery(*Query)(result_slice_addr *[]map[string]interface{},err error)
    Close()()

}
