package easy_reflect

import (
	"reflect"
)

// Type Get type
func (r *ReflectItem) Type() reflect.Kind {
	return r.vValue.Kind()
}

// Elem Take the internal type
// 取内部类型 (取指针具体类型, 取数组具体类型...)
func (r *ReflectItem) Elem() *ReflectItem {
	return &ReflectItem{
		val:    r.val,
		vType:  r.vType.Elem(),
		vValue: r.vValue.Elem(),
	}
}
