package solis

//import "crypto/rsa"
import "client/settings"
import "client/cross"
import "client/common"
import "client/common/types"

var database = cross.Database
var pubkey_path     = settings.PUBKEY_PATH
var privatekey_path = settings.PRIVATEKEY_PATH

func GenPrivateKey()() {


}

func SaveKeyToFile()() {


}

func LoadKeyFromFile()(){


}

func CheckHostId()(error){

    check_query:=cross.Query{Table:types.SOLIS_T,Type:CHECK_EXIST}


}
