package wmwerrors

import "fmt"

type WMWErrorCode int

const (
	ok = iota
	noAuth
	badReq
	internal = 999
	_max     = 1000
)

var codeMessageMap = map[WMWErrorCode]string{
	ok:       "ok",
	noAuth:   "noAuth",
	badReq:   "badReq",
	internal: "internal",
	_max:     "max",
}

type WMWError interface {
	Error() string
	Code() WMWErrorCode
}

type wmwError struct {
	msg  string
	code WMWErrorCode
}

func (w *wmwError) Error() string {
	if w == nil {
		return codeMessageMap[ok]
	}
	return w.msg
}

func (w *wmwError) Code() WMWErrorCode {
	if w == nil {
		return ok
	}
	return w.code
}

func NoAuth(a ...any) WMWError { return newErr(noAuth, a...) }

func BadReq(a ...any) WMWError { return newErr(badReq, a...) }

func Internal(a ...any) WMWError { return newErr(internal, a...) }

func Nil() WMWError { var res *wmwError; return res }

func newErr(c WMWErrorCode, a ...any) *wmwError {
	var m string
	if len(a) == 0 {
		m = codeMessageMap[c]
	} else {
		m = fmt.Sprint(a)
	}
	return &wmwError{msg: m, code: c}
}
