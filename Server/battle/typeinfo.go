package main

import "fmt"

type TypeInfo struct {
	Names map[int8]string
}

func NewTypeInfo() *TypeInfo {
	info := &TypeInfo { Names: make(map[int8]string) }
	info.init()
	return info
}

func (t *TypeInfo) init() {
	
}

func (t *TypeInfo) GetTypeName(_id int8) string {
	value, found := t.Names[_id]
	if !found {
		fmt.Printf("Can't find TypeName with ID: %d\n", _id)
	}
	
	return value
}