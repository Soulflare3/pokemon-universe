package main

type BasicInfo struct {
	name, info string
}

func NewBasicInfo() *BasicInfo {
	return &BasicInfo{ }
}