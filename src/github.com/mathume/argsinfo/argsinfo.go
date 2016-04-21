package argsinfo

import (
	"bytes"
	"errors"
	"encoding/json"
	"strings"
)

const fieldDefinitionPrefix string = "#Fields:"

type info struct{
	fields []string
	values []string
	fieldsDefined bool
}

type Info interface{
	Read(s string) error
	Values() []string
	FieldsDefined() bool
	FieldsDefinition() []string
}

func NewInfo() Info{
	m := new(info)
	m.fields = make([]string, 0)
	m.values = make([]string, 0)
	return m
}

func (this *info)FieldsDefinition() []string{
	return this.fields
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
	s = strings.Replace(s, "\r", "", -1)
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
	if(!this.fieldsDefined){
		return errors.New("Fields definition missing")
	}
	s = strings.Replace(s, "\r", "", -1)
	lines := strings.FieldsFunc(s, lineSeparator)
	
	return this.addMapFromLines(lines)
}

func minimalLength(a []string, b []string) int{
	min := len(a)
	if(len(b)<min){
		min = len(b)
	}
	
	return min
}

func (this *info)addMapFromLines(lines []string) error{
	var e error
	for i:=0; i<len(lines); i++ {
		e = this.addValue(strings.Trim(lines[i], " "))
	}
	
	return e
}

func (this *info)addValue(line string) error {
	l := strings.Fields(line)
	min := minimalLength(l, this.fields)
	if(min == 0){
		return nil
	}
	
	value := make(map[string]string)
	for i:=0; i<min; i++ {
		value[this.fields[i]] = l[i]
	}
	
	v, err := serialize(value)
	if(err != nil){
		return err
	}
	
	this.values = append(this.values, v)
	
	return nil
}

func serialize(m map[string]string) string, error{
	b := new(bytes.Buffer)
	e := json.NewEncoder(b)
	err := e.Encode(m)
	if(err != nil){
		return nil, err
	}
	
	return strings.Trim(b.String(), "\n"), nil
}

func (this *info)Values() []string{
	return this.values
}

func (this *info)FieldsDefined() bool{
	return this.fieldsDefined
}
