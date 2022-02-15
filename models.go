package easy_reflect

import "reflect"

type ReflectItem struct {
	val    interface{}
	vType  reflect.Type
	vValue reflect.Value

	structType reflect.StructField

	father *ReflectItem
}

func NewReflect(v interface{}) *ReflectItem {
	return &ReflectItem{
		val:    v,
		vType:  reflect.TypeOf(v),
		vValue: reflect.ValueOf(v),
	}
}

// base

func (r *ReflectItem) GetVal() interface{} {
	return r.val
}

func (r *ReflectItem) GetType() reflect.Type {
	return r.vType
}

func (r *ReflectItem) GetValue() reflect.Value {
	return r.vValue
}

func (r *ReflectItem) GetFather() *ReflectItem {
	return r.father
}

func (r *ReflectItem) GetStructType() reflect.StructField {
	return r.structType
}

// v1

func (r *ReflectItem) Kind() reflect.Kind {
	return r.vType.Kind()
}

func (r *ReflectItem) Elem() *ReflectItem {
	return &ReflectItem{
		father: r,
		vType:  r.vType.Elem(),
		vValue: r.vValue.Elem(),
	}
}
