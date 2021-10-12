package assignreflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestAssignmentByReflectExample1 assigned struct, slice, int, float, string and bool
func TestAssignmentByReflectExample1(t *testing.T) {
	m := map[string]interface{}{
		"StructField": struct {
			IntFieldInStruct       int
			InterfaceFieldInStruct interface{}
		}{IntFieldInStruct: 100500, InterfaceFieldInStruct: nil},
		"SliceIntField": []int{1, 3, 5, 7},
		"IntField":      111,
		"FloatField":    222.222,
		"StringField":   "test1",
		"BoolField":     true,
	}

	inModel := struct {
		StructField struct {
			IntFieldInStruct       int
			InterfaceFieldInStruct interface{}
		}
		SliceIntField []int
		IntField      int
		FloatField    float64
		StringField   string
		BoolField     bool
	}{
		StructField: struct {
			IntFieldInStruct       int
			InterfaceFieldInStruct interface{}
		}{IntFieldInStruct: 0, InterfaceFieldInStruct: nil},
		SliceIntField: []int{},
		IntField:      0,
		FloatField:    0.00,
		StringField:   "",
		BoolField:     false,
	}

	inModelAssigned := struct {
		StructField struct {
			IntFieldInStruct       int
			InterfaceFieldInStruct interface{}
		}
		SliceIntField []int
		IntField      int
		FloatField    float64
		StringField   string
		BoolField     bool
	}{
		StructField: struct {
			IntFieldInStruct       int
			InterfaceFieldInStruct interface{}
		}{IntFieldInStruct: 100500, InterfaceFieldInStruct: nil},
		SliceIntField: []int{1, 3, 5, 7},
		IntField:      111,
		FloatField:    222.222,
		StringField:   "test1",
		BoolField:     true,
	}

	err := AssignmentByReflect(&inModel, m)
	if err != nil {
		t.Fatalf("Error on assigment: %s", err)
	}

	assert.EqualValues(t, inModel, inModelAssigned, "Incorrect result, expected %v got %v", inModelAssigned, inModel)
}

// TestAssignmentByReflectExample2 assigned array, int, float and string
func TestAssignmentByReflectExample2(t *testing.T) {
	m := map[string]interface{}{
		"ArrField":    [3]int{4, 3, 1},
		"IntField":    222,
		"FloatField":  333.333,
		"StringField": "test2",
	}

	inModel := struct {
		ArrField    [3]int
		IntField    int
		FloatField  float64
		StringField string
	}{
		ArrField:    [3]int{},
		IntField:    0,
		FloatField:  0.00,
		StringField: "",
	}

	inModelAssigned := struct {
		ArrField    [3]int
		IntField    int
		FloatField  float64
		StringField string
	}{
		ArrField:    [3]int{4, 3, 1},
		IntField:    222,
		FloatField:  333.333,
		StringField: "test2",
	}

	err := AssignmentByReflect(&inModel, m)
	if err != nil {
		t.Fatalf("Error on assigment: %s", err)
	}

	assert.EqualValues(t, inModel, inModelAssigned, "Incorrect result, expected %v got %v", inModelAssigned, inModel)
}

// TestAssignmentByReflectExample3 assigned pointer, int and string
func TestAssignmentByReflectExample3(t *testing.T) {
	i := 666

	m := map[string]interface{}{
		"PointerField": &i,
		"IntField":     333,
		"StringField":  "test3",
	}

	inModel := struct {
		PointerField *int
		IntField     int
		StringField  string
	}{
		PointerField: nil,
		IntField:     0,
		StringField:  "",
	}

	inModelAssigned := struct {
		PointerField *int
		IntField     int
		StringField  string
	}{
		PointerField: &i,
		IntField:     333,
		StringField:  "test3",
	}

	err := AssignmentByReflect(&inModel, m)
	if err != nil {
		t.Fatalf("Error on assigment: %s", err)
	}

	assert.EqualValues(t, inModel, inModelAssigned, "Incorrect result, expected %v got %v", inModelAssigned, inModel)
}

// TestNewIn test initial new in struct
func TestNewIn(t *testing.T) {
	in := NewIn()
	assert.NotNil(t, in, "New struct in is nil")
}
