package easy_reflect

import (
	"errors"
	"reflect"
	"strings"
)

type ReflectStruct struct {
	ReflectItem *ReflectItem
	Index       map[string]int // 原生索引
	TagIndex    map[string]int // tag 索引
}

func (r *ReflectItem) Struct() (*ReflectStruct, error) {
	resp := &ReflectStruct{
		ReflectItem: r,
		Index:       map[string]int{},
		TagIndex:    map[string]int{},
	}
	if r.Kind() != reflect.Ptr {
		resp.ReflectItem = r.Elem()
	}
	if resp.ReflectItem.Kind() != reflect.Struct {
		return nil, errors.New("not struct or ptr struct")
	}

	for i := 0; i < resp.NumField(); i++ {
		resp.Index[resp.ReflectItem.vType.Field(i).Name] = i
	}

	return resp, nil
}

func (r *ReflectStruct) NumField() int {
	return r.ReflectItem.vType.NumField()
}

// BuildTagIndex 更具tag 构建索引 (返回构建成功非空个数)
func (r *ReflectStruct) BuildTagIndex(tag string) int {
	tagNum := 0
	for i := 0; i < r.NumField(); i++ {
		rTag := strings.TrimSpace(r.ReflectItem.vType.Field(i).Tag.Get(tag))
		if rTag == "" {
			continue
		}
		tagNum += 1
		r.TagIndex[rTag] = i
	}

	return tagNum
}

func (r *ReflectStruct) GetFieldByName(name string) (*ReflectItem, error) {
	i, ex := r.Index[name]
	if !ex {
		return nil, errors.New("not found: " + name)
	}

	vValue := r.ReflectItem.vValue.Field(i)
	vType := r.ReflectItem.vType.Field(i)

	return &ReflectItem{
		father:     r.ReflectItem,
		vValue:     vValue,
		structType: vType,
	}, nil
}

func (r *ReflectStruct) GetFieldByTag(tag string) (*ReflectItem, error) {
	i, ex := r.TagIndex[tag]
	if !ex {
		return nil, errors.New("not found: " + tag)
	}

	vValue := r.ReflectItem.vValue.Field(i)
	vType := r.ReflectItem.vType.Field(i)

	return &ReflectItem{
		father:     r.ReflectItem,
		vValue:     vValue,
		structType: vType,
	}, nil
}
