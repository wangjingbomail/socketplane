// generated by stringer -type=ValueType; DO NOT EDIT

package hcl

import "fmt"

const _ValueType_name = "ValueTypeUnknownValueTypeFloatValueTypeIntValueTypeStringValueTypeBoolValueTypeNilValueTypeListValueTypeObject"

var _ValueType_index = [...]uint8{0, 16, 30, 42, 57, 70, 82, 95, 110}

func (i ValueType) String() string {
	if i+1 >= ValueType(len(_ValueType_index)) {
		return fmt.Sprintf("ValueType(%d)", i)
	}
	return _ValueType_name[_ValueType_index[i]:_ValueType_index[i+1]]
}
