package dusk

import "gopkg.in/mgo.v2/bson"
import "wengine/core/utah"
import "wengine/core/common"

func (d *MongoDb)CreateToken(userid string)(token_id string,err error){
    _,err         = d.GetUserById(userid)
    if (err!=nil) {return "",err}
    token_id,err = common.GenId()
    if (err!=nil) {return "",err}
    token       := utah.Token{UserId:userid,Id:token_id}
    c           := d.Session.DB(d.dbname).C(d.tokens_c_name)
    err         = c.Insert(token)
    return token_id,err

}


func(d *MongoDb)TokenExists(user_id string,token_id string ) (bool) {
    token  := utah.Token{}
    c      := d.Session.DB(d.dbname).C(d.tokens_c_name)
    err    :=  c.Find(bson.M{"userid": user_id, "id":token_id}).One(&token)
    if err!=nil { return false } else { return true }
}

func(d *MongoDb)RemoveToken(token_id string ,user_id string)(error) {

    c      := d.Session.DB(d.dbname).C(d.tokens_c_name)
    err    := c.Remove(bson.M{"id": token_id,"userid":user_id })
    return err

}

func(d *MongoDb)GetAnUserToken(user_id string) (string,error) {
    token  := utah.Token{}
    c      := d.Session.DB(d.dbname).C(d.tokens_c_name)
    err    :=  c.Find(bson.M{"userid": user_id}).One(&token)
    if err!=nil { return "",err } else { return token.Id,nil  }
}

