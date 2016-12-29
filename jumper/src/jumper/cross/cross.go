package cross

type Garreth struct {
    // kind of connector
    dbtype   int
    username string
    password string
    host     string
    dbname   string
    dbpath   string
}

func CreateConnectorTemplate(dbtype string)(*Garreth,error){
    var g Garreth
    switch {
        case dbtype == "mongo" || dbtype == "mongodb":
            g.dbtype=MONGODB
            return &g, nil
        case dbtype == "postgres" || dbtype == "postgresql"  :
            g.dbtype=POSTGRESQL
            return &g, nil
        case dbtype == "bolt" || dbtype == "boltdb":
            g.dbtype=BOLTDB
            return &g, nil
    }
    return nil, db_type_is_incorrect
}

func (g Garreth)GetDbType()(int){
    return g.dbtype
}


func (g Garreth)SetCred( username string, password string )(error){
    if username == "" {
        return db_username_is_empty
    }else if password == "" {
        return db_password_is_empty
    } else {
        g.username = username
        g.password = password
        return nil
    }
}

func (g Garreth)GetCred()(string,string){
    return g.username, g.password
}

func (g Garreth)SetPath(path string)(error){
    if path == "" {
        return db_path_is_empty
    }else {
        g.dbpath = path
        return nil
    }
}

func (g Garreth)GetPath()(string){
    return g.dbpath
}



func (g Garreth)SetDbname(dbname string)(error){
    if dbname == "" {
        return db_dbname_is_empty
    }else {
        g.dbname = dbname
        return nil
    }
}

func (g Garreth)GetDbname()(string){
    return g.dbname
}

func (g Garreth)SetHost(host string)(error){
    if host == "" {
        return db_host_is_empty
    }else {
        g.host = host
        return nil
    }
}

func (g Garreth)GetHost()(string){
    return g.host
}


/*

func(g Garreth)Open()(d Database,err error) {

    switch {
        case g.Dbtype == MONGODB:
        //    d=&MongoDb{username:username,
        //              password:password,
        //              host:host,
        //              dbname:dbname}
        //    err:=d.Connect()
        //    if err == nil {
        //        return d, nil
        //    }
        case g.Dbtype == POSTGRES:



        case g.Dbtype == BOLTDB:

    }
    return nil, cant_open_database
}
*/
