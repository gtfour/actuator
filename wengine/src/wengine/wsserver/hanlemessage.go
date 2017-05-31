package wsserver

import "fmt"
import "encoding/json"
import "wengine/activa"
import "jumper/common/marconi"
import "jumper/cuda/result"


func (c *Client)handleMessage(msg *Message)(err error){
    //
    //
    //
    switch data_type := msg.DataType; data_type {
        //
        // according to data_type we convert Data field to appropriate message type 
        //
        case "data_update":
            //
            var msg_du marconi.DataUpdate
            data          := msg.Data
            err_unmarshal := json.Unmarshal( data, &msg_du )
            //
            //
            //
            if err_unmarshal == nil && msg_du.SourcePath != "/tmp/test/motion.test" {
                //
                // c.server.SendAll(&msg_chat)
                //
                fmt.Printf("\n<Message Data Update: %v\n",msg_du)
                var response      Message
                var response_data marconi.Response
                response_data.Status  =  marconi.STATUS_OK
                response.DataType     =  "server_response"
                response_data_raw,err := response_data.GetRaw()
                fmt.Printf("\nStatus message len %v\n",len(response_data_raw))
                if err == nil {
                    fmt.Printf("\n<<Sending response>>\n")
                    response.Data = response_data_raw
                    c.Write(&response)
                }
                //
                //
                //
            }
            //
            //
            //
            if msg_du.SourcePath == "/tmp/test/motion.test" {
                motion := activa.CreateMotion()
                database.WriteMotion(&motion)
                var response      Message
                response.DataType     =  "motion"
                response_data_raw,err := motion.GetRaw()
                if err == nil {
                    fmt.Printf("\n:: Sending motion ::\n")
                    response.Data = response_data_raw
                    c.Write(&response)
                }
            }
            //
            //
            //
        case "message_ws_state":
            //
        case "new_dynima":
            //
        case "dynima_response":
            fmt.Printf("\n<<Recieving dynima response>>")
            var resultsRaw result.ResultsRaw
            data          := msg.Data
            fmt.Printf("\nresultSetByte:\n%v\n", data)
            err_unmarshal := json.Unmarshal( data, &resultsRaw )
            fmt.Printf("\nunmarshal convert : %v", err_unmarshal)
            if err_unmarshal == nil {
                for i:= range resultsRaw {
                    r := resultsRaw[i]
                    fmt.Printf("\n-- Result with Type: %v --\n", r.Type)
                    fmt.Printf("%v\n",r.Result)
                    fmt.Printf("\n-- -- --\n")
                }

            }
            //
    }
    return nil
}

