package dbsync
import (

        //"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
        "wengine_parts/repository"
        //"wengine_parts/settings"
        "wengine_parts/airparse"
)


type Repository struct {

    Url      string
    Packages []repository.RpmPackage

}



func UploadStructToDb(repofile *airparse.RepoFile) (err error){ 

   session, err := mgo.Dial("mongodb://wengine:OpenStack123@127.0.0.1/test") // settings.mongo_host

   if err != nil {

      panic(err)

   }
   defer session.Close()

   session.SetMode(mgo.Monotonic, true)

   c      :=session.DB("wengine").C("repository")
   result :=Repository {}

   err = c.Find(bson.M{"url": repofile.Url}).One(&result)

   if err!=nil {

       err = c.Insert(&Repository{Url: repofile.Url,Packages: repofile.Packages})

   } else {

       err = c.Update(bson.M{"url": repofile.Url},&Repository{Url: repofile.Url, Packages: repofile.Packages })

   }

   return nil
}
