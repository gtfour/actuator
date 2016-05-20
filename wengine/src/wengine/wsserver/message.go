package wsserver
import "encoding/json"

type Message struct {
    DataType   string          `json:"datatype"`
    Data       json.RawMessage `json:"data"`
}

