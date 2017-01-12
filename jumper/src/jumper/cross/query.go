package cross

type Query struct {
    Type        int
    Table       string
    TableList   []string
    //KeyBody   map[string]interface{}
    //QueryBody map[string]interface{}
    QueryXY
    Queries     []QueryXY
    Multi       bool
    RonB        bool // means that queries from Queries field will run over nested bucket identified by KeyBody
}

type QueryXY struct {
    KeyBody   map[string]interface{}
    QueryBody map[string]interface{}
}

func MakeQuery(query_type int, table_name string)(Query){
    q:=Query{Type:query_type,Table:table_name}
    return q
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
    query_type:=q.Type
    return IsOk(query_type, TABLE_OPS)
}

func(q *Query)CheckTableName()(error){
    table_name:=q.Table
    if table_name == "" {
        return EmptyTableName
    }
    return nil
}

func(q *Query)Validate()(match_by_key bool,match_by_value bool,err error){
    match_by_key,match_by_value,err = q.ValidateBodies()
    if err == nil {  err = q.CheckTableName()  }
    return
}
