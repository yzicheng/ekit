// Code generated by go2go; DO NOT EDIT.


//line proxy.go2:1
package proxy

//line proxy.go2:1
import (
//line proxy.go2:1
 "reflect"
//line proxy.go2:1
 "testing"
//line proxy.go2:1
)

//line proxy.go2:38
type Invocation struct {
	Method reflect.Method
	Args   []reflect.Value
}

func (inv *Invocation) Invoke() []reflect.Value {
	return inv.Method.Func.Call(inv.Args)
}

//line proxy.go2:45
type Importable୦ int

//line proxy.go2:45
var _ = reflect.Append
//line proxy.go2:45
var _ = testing.AllocsPerRun
