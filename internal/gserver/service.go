package gserver

import "reflect"

type Service struct {
	name string
	typ reflect.Type
	val reflect.Value
	methods map[string] reflect.Method
}