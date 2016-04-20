package argsinfo

import (
	"testing"
	"encoding/json"
	"bytes"
	"strings"
)

func TestTestingEncodingFunction(t *testing.T){
	encoded := encode(map[string]string{"a":"b"})
	expected := "{\"a\":\"b\"}"
	if encoded != expected{
		t.Error("actual:", encoded, "expected:", expected)
	}
}

func TestFieldsDefinedIsFalseIfNotRead(t *testing.T){
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

func TestOneDefinitionIsSetCheckContent(t *testing.T) {
	input := "#Fields: a"
	info := NewInfo()
	_ = info.Read(input)
	if(len(info.FieldsDefinition()) <= 0 || info.FieldsDefinition()[0] != "a"){
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

func TestOneValueCorrectLength(t *testing.T){
	input := `#Fields: a b
aValue bValue`
	info := NewInfo()
	_ = info.Read(input)
	if(len(info.Values()) != 1){
		t.Fail()
	}
}

func TestOneValueCorrectValue(t *testing.T){
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
		t.Error(len(info.Values()),"values:", info.Values(), "expected:", expectedOutput)
	}
}

func TestAddSecondValue(t *testing.T){
	input1 := `#Fields: a b
aValue bValue`
	input2 := "a1Value b1Value"
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
a1Value b1Value`
	expectedOutput := encode(map[string]string{"a":"a1Value","b":"b1Value"})
	info := NewInfo()
	_ = info.Read(input)
	if(len(info.Values()) != 2 || info.Values()[1] != expectedOutput){
		t.Error("len", len(info.Values()), "value:", info.Values()[1], "expected:", expectedOutput)
	}
}

func TestLessFieldsThanValues(t *testing.T){
	input := `#Fields: a b
a1Value b1Value cValue`
	expectedOutput := encode(map[string]string{"a":"a1Value","b":"b1Value"})
	info := NewInfo()
	_ = info.Read(input)
	if(len(info.Values()) != 1 || info.Values()[0] != expectedOutput){
		t.Error(info.Values())
	}
}

func TestLessValuesThanFields(t *testing.T){
	input := `#Fields: a b
a1Value`
	expectedOutput := encode(map[string]string{"a":"a1Value"})
	info := NewInfo()
	_ = info.Read(input)
	if(len(info.Values()) != 1 || info.Values()[0] != expectedOutput){
		t.Error(info.Values())
	}
}

func encode(m map[string]string)	string{
		b := new(bytes.Buffer)
		e := json.NewEncoder(b)
		err := e.Encode(m);
		if(err != nil){
			return "couldn't encode in test"
		}
		return strings.Trim(b.String(), "\n")
}