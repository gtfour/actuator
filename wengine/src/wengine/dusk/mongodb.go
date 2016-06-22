package dusk

import "gopkg.in/mgo.v2"

type MongoDb struct {
    Session  *mgo.Session
    Users    []string
    Groups   []string
    username string
    password string
    host     string
    dbname   string
    CollectionsNames
}
type CollectionsNames struct {
    users_c_name            string
    actions_c_name          string
    dashboards_c_name       string
    dashboard_groups_c_name string
    tokens_c_name           string
}

func (d *MongoDb)Close()() {
    d.Session.Close()
}

func (d *MongoDb)Connect() ( err error ) {
    d.Session,err = mgo.Dial("mongodb://"+d.username+":"+d.password+"@"+d.host+"/"+d.dbname)
    d.Session.SetMode(mgo.Monotonic, true)
    d.users_c_name            ="dashboard_users"
    d.tokens_c_name           = "user_tokens"
    d.dashboards_c_name       = "dashboards"
    d.dashboard_groups_c_name = "dashboard_groups"
    return err
}
