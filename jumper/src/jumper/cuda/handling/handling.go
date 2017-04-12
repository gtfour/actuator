package handling

import "jumper/cuda/result"
import "jumper/cuda/targets"
// //  import "jumper/cuda/filtering"

//
//  here we will run filters over targets and return result
//

func Handle(t targets.Target)(result.Result){
    //
    //
    //
    switch target_type:=t.GetType();target_type {
        case targets.TARGET_LINE:
            return result.BlankResult( result.RESULT_TYPE_LINE )
        case targets.TARGET_FILE:
            return result.BlankResult( result.RESULT_TYPE_FILE )
        case targets.TARGET_DIR:
            return result.BlankResult( result.RESULT_TYPE_DIR )
        default:
            return nil
    }
    //
    //
    //
}
