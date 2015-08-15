package dbsync
import (

        //"log"

        "gopkg.in/mgo.v2"

        //"gopkg.in/mgo.v2/bson"

        //"wengine_parts/repository"
        //"wengine_parts/settings"
        "wengine_parts/airparse"
)

type Package struct {

    Name string
    Type string


}


func UploadStructToDb(repofile airparse.RepoFile) (err error){ 

   session, err := mgo.Dial("mongodb://wengine:OpenStack123@127.0.0.1/test") // settings.mongo_host

   if err != nil {
      panic(err)
   }
   defer session.Close()

   session.SetMode(mgo.Monotonic, true)
   c:=session.DB("wengine").C("package")
   err=c.Insert(&Package{Name:"python",Type:"rpm"},&Package{Name:"python-dev",Type:"deb"})

   // settings.mongo_host
   // 127.0.0.1

   return nil
}
