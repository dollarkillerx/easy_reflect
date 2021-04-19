package easy_reflect

import (
	"fmt"

	"github.com/pkg/errors"
)

// NumField View the number of structure fields
// 查看结构体字段数量
func (r *ReflectItem) NumField() (int, error) {
	if !r.IsStruct() {
		return 0, errors.New("Structure must be struct")
	}

	return r.vType.NumField(), nil
}

// FieldByIndex finds the field by Index
// 通过Index找到字段
func (r *ReflectItem) FieldByIndex(i int) (*Field, error) {
	if !r.IsStruct() {
		return nil, errors.New("Structure must be struct")
	}

	return &Field{
		Value:       r.vValue.Field(i),
		StructField: r.vType.Field(i),
	}, nil
}

// FieldByIndex finds the field by Name
// 通过Name找到字段
func (r *ReflectItem) FieldByName(name string) (*Field, error) {
	if !r.IsStruct() {
		return nil, errors.New("Structure must be struct")
	}

	val := r.vValue.FieldByName(name)
	st, b := r.vType.FieldByName(name)
	if !b {
		return nil, errors.New(fmt.Sprintf("%s does not exist", name))
	}

	return &Field{
		StructField: st,
		Value:       val,
	}, nil
}
