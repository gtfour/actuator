package dusk

import "fmt"

func(d *MongoDb)CheckAccess(initiator_type,initiator_id,target_type,target_id string)(error) {

    fmt.Printf("InitiatorType: %s InitiatorId: %s\nTargetType: %s TargetId: %s",initiator_type,initiator_id,target_type,target_id)
    //user   := utah.User{}
    /*dashboard_list.List = make([]dashboard.Dashboard, 0)
    if d.TokenExists(user_id,token_id) == false { return dashboard_list,TokenDoesNotExist() }
    user:=utah.User{}
    cu      := d.Session.DB(d.dbname).C(d.users_c_name)
    err     =  cu.Find(bson.M{"id": user_id}).One(&user)
    if err!= nil { return dashboard_list,err }
    cd      := d.Session.DB(d.dbname).C(d.dashboards_c_name)
    for d := range user.Dashboards {
        d_id:=user.Dashboards[d]
        dashboard:=dashboard.Dashboard{}
        err     =  cd.Find(bson.M{"id": d_id}).One(&dashboard)
        if err == nil { dashboard_list.List = append(dashboard_list.List, dashboard)  }
    }
    return dashboard_list,nil
    */
    return nil
}

func(d *MongoDb)GrantAccess(initiator_type,initiator_id,target_type,target_id string)(error) {
    return nil
}

func(d *MongoDb)RemoveAccess(initiator_type,initiator_id,target_type,target_id string)(error) {
    return nil
}



