package validate

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"sync"
	"time"
)

type From interface {
	SetMessage() map[string]string
}

// 初始化验证器
var validate *validator.Validate

// 重写gin框架的验证器
type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

func (d *defaultValidator) ValidateStruct(obj interface{}) error {

	value := reflect.ValueOf(obj)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}

	if valueType == reflect.Struct {
		d.lazyinit()

		if err := d.validate.Struct(obj); err != nil {
			errMsg := err.Error()
			if _, ok := interface{}(obj).(From); ok {
				temp := value.MethodByName("SetMessage").Call(nil)[0]
				message := temp.Interface().(map[string]string)

				errs := err.(validator.ValidationErrors)
				for _, err := range errs {
					messageKey := fmt.Sprintf("%s.%s", err.Field, err.Tag)
					if _, ok := message[messageKey]; ok {
						errMsg = message[messageKey]
						break
					}
				}
			}
			return errors.New(errMsg)
		}
	}
	return nil
}

func (d *defaultValidator) Engine() interface{} {
	d.lazyinit()
	return d.validate
}

func (d *defaultValidator) lazyinit() {
	d.once.Do(func() {
		config := &validator.Config{TagName: "binding"}
		d.validate = validator.New(config)
	})
}

func init() {
	binding.Validator = &defaultValidator{}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("date", DateValidate)
	}
}

// 自定规则-判断日期格式是否正确
func DateValidate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	loc, _ := time.LoadLocation("Local")
	_, err := time.ParseInLocation(param, field.String(), loc)

	if err != nil {
		return false
	}

	return true
}
