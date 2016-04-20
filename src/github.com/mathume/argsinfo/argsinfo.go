package argsinfo

import (
//	"encoding/json"
	"strings"
//	"unicode"
)

const fieldDefinitionPrefix string = "#Fields:"

type info struct{
	fields []string
	maps [](map[string]string)
	fieldsDefined bool
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
	var e error
	if(strings.HasPrefix(strings.Trim(s, " "), fieldDefinitionPrefix)){
		e = this.setFieldDefinition(s)
	}else{
		e = this.addMapFromString(s)
	}
	
	return e
}

func (this *info)setFieldDefinition(s string) error{
	lines := strings.FieldsFunc(s, lineSeparator)
	this.fields = strings.Fields(lines[0])[1:]
	this.fieldsDefined = true
	if(len(lines) > 1){
		this.addMapFromLines(lines[1:])
	}
	return nil
}

func lineSeparator(c rune) bool {
	return (c == '\n')
}

func (this *info)addMapFromString(s string) error{
	return nil
}

func (this *info)addMapFromLines(lines []string) error{
	return nil
}

func (this *info)Values() []string{
	return nil
}

func (this *info)FieldsDefined() bool{
	return this.fieldsDefined
}