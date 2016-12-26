package cross

type garreth struct {
    dbtype   int
    username string
    password string
    host     string
    dbname   string
    dbpath   string
}

func CreateConnector(dbtype string)(*garreth,error){
    var g garreth
    switch {
        case dbtype == "mongo" || dbtype == "mongodb":
            g.dbtype=MONGODB
            return &g, nil
        case dbtype == "postgres":
            g.dbtype=POSTGRES
            return &g, nil
        case dbtype == "bolt" || dbtype == "boltdb":
            g.dbtype=BOLTDB
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

func (g garreth)set_path( path string )(error){
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








func(g garreth)Open()(d Database,err error) {

    switch {
        case g.dbtype == MONGODB:
            d=&MongoDb{username:username,
                      password:password,
                      host:host,
                      dbname:dbname}
            err:=d.Connect()
            if err == nil {
                return d, nil
            }
        case g.dbtype == POSTGRES:



        case g.dbtype == BOLTDB:

    }
    return nil, cant_open_database
}





