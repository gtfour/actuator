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

func InputPropValidation(prop map[string]string)(int) {

    _, ok_initiator_type := prop["initiator_type"]
    _, ok_initiator_id   := prop["initiator_id"]
    _, ok_action_type    := prop["action_type"]
    _, ok_target_type    := prop["target_type"]
    _, ok_target_id      := prop["target_id"]
    //action_id, ok := m["initiator_type"]

        if ok_initiator_type == false {return ARISTO_initiatorTypeIsEmpty}
        if ok_initiator_id   == false {return ARISTO_initiatorIdIsEmpty}
        if ok_action_type    == false {return ARISTO_actionTypeIsEmpty}
        if ok_target_type    == false {return ARISTO_targetTypeIsEmpty}
        if ok_target_id      == false {return ARISTO_targetIdIsEmpty}

    return ARISTO_inputIsCorrect

}

func InputStanceValidation(prop map[string]string)(int){
    return ARISTO_inputIsCorrect
}

func(d *MongoDb)GrantAccess(prop map[string]string)(error,int) {

    check_input_state:=InputPropValidation(prop)
    if check_input_state != ARISTO_inputIsCorrect {
        return nil,check_input_state
    }
    return nil,ARISTO_AccessGranted
}

func(d *MongoDb)RemoveAccess(prop map[string]string)(error,int) {

    check_input_state:=InputPropValidation(prop)
    if check_input_state != ARISTO_inputIsCorrect {
        return nil,check_input_state
    }
    return nil,ARISTO_AccessRemoved

}

func(d *MongoDb)GetPrivileges(prop map[string]string)(error){
    // _, ok := m["route"]

    return nil

}

func(d *MongoDb)AristoMakeQuery(prop map[string]string)(error){
    // _, ok := m["route"]

    return nil

}


