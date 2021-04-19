package easy_reflect

import "reflect"

func (r *ReflectItem) IsPtr() bool {
	if r.Type() == reflect.Ptr {
		return true
	}

	return false
}

func (r *ReflectItem) IsSlice() bool {
	if r.Type() == reflect.Slice {
		return true
	}

	return false
}

func (r *ReflectItem) IsMap() bool {
	if r.Type() == reflect.Map {
		return true
	}

	return false
}

func (r *ReflectItem) IsStruct() bool {
	if r.Type() == reflect.Struct {
		return true
	}

	return false
}
