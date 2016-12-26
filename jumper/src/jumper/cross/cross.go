package cross

type garreth struct {

    username string
    password string
    host     string
    dbname   string
    dbpath   string

}

func (g garreth) set_cred ( username string, password string )(error){

    return nil


}

func OpenDatabase ( g garreth) ( d Database,err error ) {

    switch {
        case dbtype == "mongo":
            d=&MongoDb{username:username,
                      password:password,
                      host:host,
                      dbname:dbname}
            err:=d.Connect()
            if err == nil {
                return d
            }
        case dbtype == "postgres":



        case dbtype == "boltdb":

    }
    return nil
}





