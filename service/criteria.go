package service

import (
	"errors"
	"fmt"
	"github.com/cloudintheking/go-criteria/utils"
	"reflect"
	"strconv"
	"strings"
)

type StringFilter struct {
	Regexp    *string
	Contains  *string
	Equals    *string
	NotEquals *string
	In        *string
}

type IntFilter struct {
	Equals    *int64
	NotEquals *int64
	Lte       *int64
	Lt        *int64
	Gt        *int64
	Gte       *int64
	In        *string
}

type FloatFilter struct {
	Equals    *float64
	NotEquals *float64
	Lte       *float64
	Lt        *float64
	Gt        *float64
	Gte       *float64
	In        *string
}

type TimeFilter struct {
	Lte *string
	Lt  *string
	Gt  *string
	Gte *string
}

type BoolFilter struct {
	Equals *bool
}

/**
 * @Description:  gin queryParams 转换为 criteria
 * @param m
 * @param sType
 */
func GinQuery2Criteria(queryParams map[string][]string, criteria interface{}) error {
	queryMap := getGinMap(queryParams)
	return Map2Struct(criteria, queryMap)
}

/**
 * @Description: 将gin map[string][]string 转 map[string]interface{}
 * @param req
 * @return map[string]interface{}
 */
func getGinMap(req map[string][]string) map[string]interface{} {
	var des = map[string]interface{}{}
	for k, v := range req {
		karr := strings.Split(k, ".")
		if len(karr) == 2 {
			outKey := utils.Ucfirst(karr[0])
			inKey := utils.Ucfirst(karr[1])
			if _, ok := des[outKey]; !ok {
				des[outKey] = map[string]interface{}{}
			}
			des[outKey].(map[string]interface{})[inKey] = v[0]
		} else {
			outKey := strings.ToTitle(k)
			des[outKey] = v[0]
		}
	}
	return des
}

/**
 * @Description:  map转struct简易封装(适用于value为string类型)
 * @param dst 自定义criteria
 * @param src map
 * @return err
 */
func Map2Struct(dst, src interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()
	dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	err = reflect2Struct(dstType, dstValue, src)
	return err
}

/**
 * @Description:  src反射进struct中
 * @param dstType
 * @param dstValue
 * @param src
 * @return error
 */
func reflect2Struct(dstType reflect.Type, dstValue reflect.Value, src interface{}) error {
	// dst必须结构体指针类型
	if dstType.Kind() == reflect.Ptr {
		// 取具体内容
		dstType, dstValue = dstType.Elem(), dstValue.Elem()
	}
	if dstType.Kind() != reflect.Struct {
		return errors.New("dst type should be a struct pointer")
	}
	srcType := reflect.TypeOf(src)
	// src必须为map或者map指针
	if srcType.Kind() == reflect.Ptr {
		// 取具体内容
		srcType = srcType.Elem()
	}
	if srcType.Kind() != reflect.Map {
		return errors.New("src type should be a map or a map pointer")
	}

	m := src.(map[string]interface{})
	for i := 0; i < dstType.NumField(); i++ {
		// 属性
		property := dstType.Field(i)
		propertyName := property.Name

		v, ok := m[propertyName]
		// map中不存在该key
		if !ok {
			continue
		}
		propertyValue := dstValue.Field(i)
		//只赋值公共属性
		if !propertyValue.CanSet() {
			continue
		}
		if propertyValue.Kind() == reflect.Ptr {
			if propertyValue.IsNil() {
				propertyValue.Set(reflect.New(propertyValue.Type().Elem())) //指针初始化
			}
			propertyValue = propertyValue.Elem() //获取指针所指对象
		}
		switch propertyValue.Kind() {
		case reflect.Struct:
			reflect2Struct(propertyValue.Type(), propertyValue, v)
		case reflect.Int64:
			res, _ := strconv.ParseInt(v.(string), 10, 64)
			propertyValue.SetInt(res)
		case reflect.Float64:
			res, _ := strconv.ParseFloat(v.(string), 64)
			propertyValue.SetFloat(res)
		case reflect.Bool:
			res, _ := strconv.ParseBool(v.(string))
			propertyValue.SetBool(res)
		case reflect.String:
			propertyValue.SetString(v.(string))

		}
	}
	return nil
}
