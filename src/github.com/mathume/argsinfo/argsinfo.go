package argsinfo

import "encoding/json"

type info struct{
	fields []string
	maps [](map[string]string)
}

type Info interface{
	Read(s string) error
	Values() []string
	FieldsDefined() bool
}

func NewInfo() Info{
	m := new(info)
	m.fields = make([]string, 1, 100)
	m.maps = make([](map[string]string), 1, 100)
	return m
}

func (this *info)Read(s string) error{
	return new(json.InvalidUTF8Error)
}

func (this *info)Values() []string{
	return nil
}

func (this *info)FieldsDefined() bool{
	return false
}