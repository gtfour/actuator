package ws
import "github.com/gin-gonic/gin"
import "github.com/gorilla/websocket"
import "net/http"
import "fmt"

var wsupgrader = websocket.Upgrader {
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func wshandler  (w http.ResponseWriter, r *http.Request ) {

       conn,err := wsupgrader.Upgrade(w, r, nil)
        if err !=nil {
            fmt.Println("Failed to set websocket upgrade: %+v", err )
            return
        }
        for {
            t, msg, err:= conn.ReadMessage()
            if err != nil {
                break
            }
            conn.WriteMessage(t, []byte(string(msg)+"\nhello\n"))

        }
}



func WS(data  gin.H)( func(c *gin.Context) ) {

    return func(c *gin.Context)  {
        wshandler(c.Writer, c.Request)
    }

}
