package argsinfo

import (
	"testing"
	"encoding/json"
	"bytes"
)

func TestFieldsDefinedIsFalseIfNotRead(){
	info := NewInfo()
	if(info.FieldsDefined()){
		t.Fail()
	}
}

func TestOneDefinitionIsSet(t *testing.T) {
	input := "#Fields: a"
	info := NewInfo()
	e := info.Read(input)
	if(e != nil || !info.FieldsDefined()){
		t.Fail()
	}
}

func TestOneValueWithoutDefinition(t *testing.T){
	input := "a b"
	info := NewInfo()
	e := info.Read(input)
	if(e == nil){
		t.Fail()
	}
}

func TestOneValue(t *testing.T){
	input := `#Fields: a b
aValue bValue`
	expectedOutput := encode(map[string]string{"a":"aValue","b":"bValue"})
	info := NewInfo()
	_ = info.Read(input)
	if(len(info.Values()) != 1 || info.Values()[0] != expectedOutput){
		t.Fail()
	}
}

func TestAddValue(t *testing.T){
	input1 := "#Fields: a b"
	input2 := "aValue bValue"
	expectedOutput := encode(map[string]string{"a":"aValue","b":"bValue"})
	info := NewInfo()
	_ = info.Read(input1)
	_ = info.Read(input2)
	if(len(info.Values()) != 1 || info.Values()[0] != expectedOutput){
		t.Fail()
	}
}

func TestAddSecondValue(t *testing.T){
	input1 := `#Fields: a b
aValue bValue`
	input2 := "a1Value a2Value"
	expectedOutput := encode(map[string]string{"a":"a1Value","b":"b1Value"})
	info := NewInfo()
	_ = info.Read(input1)
	_ = info.Read(input2)
	if(len(info.Values()) != 2 || info.Values()[1] != expectedOutput){
		t.Fail()
	}
}	

func TestTwoValues(t *testing.T){
	input := `#Fields: a b
aValue bValue
a1Value a2Value`
	expectedOutput := encode(map[string]string{"a":"a1Value","b":"b1Value"})
	info := NewInfo()
	_ = info.Read(input)
	if(len(info.Values()) != 2 || info.Values()[1] != expectedOutput){
		t.Fail()
	}
}

func encode(m map[string]string)	string{
		b := new(bytes.Buffer)
		e := json.NewEncoder(b)
		if(e != nil){
			return "couldn't encode in test"
		}
		return b.String()
}