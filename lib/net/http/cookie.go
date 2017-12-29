package http

import (
	"net/http"

	"github.com/dop251/goja"
	"github.com/zengming00/go-server-js/lib"
)

type _cookie struct {
	runtime *goja.Runtime
	cookie  *http.Cookie
}

func (This *_cookie) stringFunc(call goja.FunctionCall) goja.Value {
	str := This.cookie.String()
	return This.runtime.ToValue(str)
}

func NewCookie(runtime *goja.Runtime, cookie *http.Cookie) *goja.Object {
	This := &_cookie{
		runtime: runtime,
		cookie:  cookie,
	}
	o := runtime.NewObject()
	o.Set("string", This.stringFunc)
	o.Set("domain", cookie.Domain)
	o.Set("expires", lib.NewTime(runtime, &cookie.Expires))
	o.Set("httpOnly", cookie.HttpOnly)
	o.Set("maxAge", cookie.MaxAge)
	o.Set("name", cookie.Name)
	o.Set("path", cookie.Path)
	o.Set("raw", cookie.Raw)
	o.Set("rawExpires", cookie.RawExpires)
	o.Set("secure", cookie.Secure)
	o.Set("unparsed", cookie.Unparsed)
	o.Set("value", cookie.Value)
	return o
}