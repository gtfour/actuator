package mongodb_edge

import "gopkg.in/mgo.v2"
import "jumper/cross"

type MongoDb struct {
    session  *mgo.Session
    //Users    []string
    //Groups   []string
    username string
    password string
    host     string
    dbname   string
    //CollectionsNames
}
/*
type CollectionsNames struct {
    users_c_name            string
    actions_c_name          string
    dashboards_c_name       string
    dashboard_groups_c_name string
    tokens_c_name           string
    motions_c_name          string
}
*/


func (d *MongoDb)Close()() {
    d.session.Close()
}

func (d *MongoDb)Connect() ( err error ) {
    d.session,err = mgo.Dial("mongodb://"+d.username+":"+d.password+"@"+d.host+"/"+d.dbname)
    d.session.SetMode(mgo.Monotonic, true)
    //d.users_c_name            ="dashboard_users"
    //d.tokens_c_name           = "user_tokens"
    //d.dashboards_c_name       = "dashboards"
    //d.dashboard_groups_c_name = "dashboard_groups"
    //d.motions_c_name          = "motions"
    return err
}

func GetDatabase(g *cross.Garreth)(s *MongoDb,err error){
    path:=g.GetPath()
    s.path=path
    return s, nil
}


