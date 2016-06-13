package sisb
import "fmt"

func CreateDynima(id string)(error) {

    fmt.Printf("Storage error %v\nNew dynima id %s\n",STORAGE_INSTANCE.Error,id)

    return nil


}

