// Code generated by "stringer -type=Status"; DO NOT EDIT.

package http

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RespOK-0]
	_ = x[RespErrorInvaildRequest-1]
	_ = x[RespErrorInternalServer-2]
	_ = x[RespErrorNotFound-3]
	_ = x[RespConfigDecodeFailed-4]
	_ = x[RespConfigCheckFailed-5]
	_ = x[RespSendToChanFailed-6]
}

const _Status_name = "RespOKRespErrorInvaildRequestRespErrorInternalServerRespErrorNotFoundRespConfigDecodeFailedRespConfigCheckFailedRespSendToChanFailed"

var _Status_index = [...]uint8{0, 6, 29, 52, 69, 91, 112, 132}

func (i Status) String() string {
	if i >= Status(len(_Status_index)-1) {
		return "Status(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
