package dusk

//import "gopkg.in/mgo.v2/bson"
import "wengine/activa"

func(d *MongoDb)WriteMotion(m *activa.Motion)(err error) {

    //token_id,err = common.GenId()
    //raw,err:=bson.Marshal(d)

    if (err!=nil) {return err}
    c           := d.Session.DB(d.dbname).C(d.motions_c_name)
    err         = c.Insert(m)

    return err

}

func (d *MongoDb)GetMotion(keys ...activa.Key)(motions []activa.Motion, err error) {
    //c      := d.Session.DB(d.dbname).C(d.users_c_name)
    //err    =  c.Find(bson.M(user_prop)).One(&user)
    if err!=nil {
        return motions,err
    }
    return motions,nil
}
