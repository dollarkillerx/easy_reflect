package easy_reflect

import "reflect"

type ReflectItem struct {
	val    interface{}
	vType  reflect.Type
	vValue reflect.Value
}

func NewReflect(v interface{}) *ReflectItem {
	return &ReflectItem{
		val:    v,
		vType:  reflect.TypeOf(v),
		vValue: reflect.ValueOf(v),
	}
}

func (r *ReflectItem) GetVal() interface{} {
	return r.val
}

func (r *ReflectItem) GetValType() reflect.Type {
	return r.vType
}

func (r *ReflectItem) GetValValue() reflect.Value {
	return r.vValue
}

type Field struct {
	reflect.StructField
	reflect.Value
}

func (r *Field) GetValStructField() reflect.StructField {
	return r.StructField
}

func (r *Field) GetValValue() reflect.Value {
	return r.Value
}
