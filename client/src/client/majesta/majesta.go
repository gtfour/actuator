package majesta

import "fmt"
import "reflect"
import "jumper/actuator"

type CompNotes struct {
    Path  string
    State int8
    //
    SourceType string // file or directory or command ( source: actuator or blackout  )
    SourceName string
    SourcePath string // /filename or /command_name
    //UpdateType string // Update,Append,Remove,RemoveFile
    //UpdateData string //
    DataHash   string
    //ServerTime string
    //ServerId   string
    //
    List  []CompNote
}

type CompNote struct {
    Field    string
    Before   string
    After    string
}


func CompareProp(old_prop,new_prop *actuator.Prop, path string)(cnotes CompNotes) {

    valueOld:=reflect.ValueOf(old_prop).Elem()
    valueNew:=reflect.ValueOf(new_prop).Elem()
    if (new_prop.IsDir == true ) {
        cnotes.SourceType = "dir"
    } else if ( new_prop.IsRegular == true ) {
        cnotes.SourceType = "file"
    }

    field:=reflect.TypeOf(old_prop).Elem()

    old_field_count := valueOld.NumField()

    for i := 0; i <= old_field_count-1; i++  {

        if string(field.Field(i).Tag)!="ignore" &&  fmt.Sprint(valueOld.Field(i).Interface())!=fmt.Sprint(valueNew.Field(i).Interface()) {
             cnote:=CompNote{Field:field.Field(i).Name,Before:fmt.Sprint(valueOld.Field(i).Interface()),After:fmt.Sprint(valueNew.Field(i).Interface())}
             cnotes.List=append(cnotes.List, cnote)
        }
    }
    cnotes.Path = path
    return cnotes
}

func Initial(prop *actuator.Prop, path string) (cnotes CompNotes) {

    value:=reflect.ValueOf(prop).Elem()
    fields_count := value.NumField()
    field:=reflect.TypeOf(prop).Elem()

    for i := 0; i <= fields_count-1; i++  {
        cnote:=CompNote{ Field:field.Field(i).Name,
                                   Before:"",
                                   After:fmt.Sprint(value.Field(i).Interface()) }
        cnotes.List=append(cnotes.List, cnote)
    }
    cnotes.Path = path
    return cnotes
}

func (cn *CompNotes) FieldExists ( field string )(exists bool) {

    for cnote_id := range cn.List {
        cnote:=cn.List[cnote_id]
        if cnote.Field == field { exists = true ; break  }
    }
    return exists
}



