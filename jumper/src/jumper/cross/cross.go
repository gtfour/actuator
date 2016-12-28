package cross

type garreth struct {
    Dbtype   int
    username string
    password string
    host     string
    dbname   string
    dbpath   string
}

func CreateConnectorTemplate(dbtype string)(*garreth,error){
    var g garreth
    switch {
        case dbtype == "mongo" || dbtype == "mongodb":
            g.Dbtype=MONGODB
            return &g, nil
        case dbtype == "postgres":
            g.Dbtype=POSTGRES
            return &g, nil
        case dbtype == "bolt" || dbtype == "boltdb":
            g.Dbtype=BOLTDB
            return &g, nil
    }
    return nil, db_type_is_incorrect
}

func (g garreth)set_cred( username string, password string )(error){
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

func (g garreth)set_path(path string)(error){
    if path == "" {
        return db_path_is_empty
    }else {
        g.dbpath = path
        return nil
    }
}

func (g garreth)set_dbname(dbname string)(error){
    if dbname == "" {
        return db_dbname_is_empty
    }else {
        g.dbname = dbname
        return nil
    }
}


func(g garreth)Open()(d Database,err error) {

    switch {
        case g.Dbtype == MONGODB:
        /*    d=&MongoDb{username:username,
                      password:password,
                      host:host,
                      dbname:dbname}
            err:=d.Connect()
            if err == nil {
                return d, nil
            }*/
        case g.Dbtype == POSTGRES:



        case g.Dbtype == BOLTDB:

    }
    return nil, cant_open_database
}
