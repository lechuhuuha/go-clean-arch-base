package validate

import (
	"reflect"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"

	"github.com/lechuhuuha/oneshield/pkg/sperrors"
	"github.com/lechuhuuha/oneshield/pkg/utils"
	"github.com/lechuhuuha/oneshield/transports/http/handlers/response"
)

var mValidator *validator.Validate
var trans ut.Translator

func init() {
	mValidator = validator.New()
	enLocale := en.New()
	uni := ut.New(enLocale, enLocale)
	trans, _ = uni.GetTranslator("enLocale")
	customValidation()
	translation()
}

func customValidation() {
	var err error

	if err = mValidator.RegisterValidation("custom_url", func(fl validator.FieldLevel) bool {
		return utils.IsValidURL(fl.Field().String())
	}); err != nil {
		panic(err)
	}
}

func translation() {
	var err error
	if err = mValidator.RegisterTranslation("custom_url", trans, func(ut ut.Translator) error {
		return ut.Add(
			"custom_url",
			"{0} can not parse url",
			true,
		)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("custom_url", fe.Field(), fe.Param())
		return t
	}); err != nil {
		panic(err)
	}
}

func ValidationErrorsToErrorsResponse(errors validator.ValidationErrors) []*response.ErrorResponse {
	var result = make([]*response.ErrorResponse, 0, len(errors))

	for _, fieldError := range errors {
		result = append(result, response.FormatError(
			"ERROR",
			"Bad request error.",
			fieldError.Translate(trans),
			fieldError.Field(),
			sperrors.CodeInvalidData,
			fieldError.Value(),
		))
	}

	return result
}

func Validate(v interface{}) error {
	return mValidator.Struct(v)
}

// setDefaultValidation set default validation
func setDefaultValidation(fl validator.FieldLevel) bool {
	field := fl.Field()

	if !field.CanSet() {
		return false
	}

	var value reflect.Value
	// only support primitive types
	switch field.Kind() {
	case reflect.String:
		value = reflect.ValueOf(fl.Param())
	case reflect.Int:
		value = reflect.ValueOf(cast.ToInt(fl.Param()))
	case reflect.Int8:
		value = reflect.ValueOf(cast.ToInt8(fl.Param()))
	case reflect.Int16:
		value = reflect.ValueOf(cast.ToInt16(fl.Param()))
	case reflect.Int32:
		value = reflect.ValueOf(cast.ToInt32(fl.Param()))
	case reflect.Int64:
		value = reflect.ValueOf(cast.ToInt64(fl.Param()))
	case reflect.Uint:
		value = reflect.ValueOf(cast.ToUint(fl.Param()))
	case reflect.Uint8:
		value = reflect.ValueOf(cast.ToUint8(fl.Param()))
	case reflect.Uint16:
		value = reflect.ValueOf(cast.ToUint16(fl.Param()))
	case reflect.Uint32:
		value = reflect.ValueOf(cast.ToUint32(fl.Param()))
	case reflect.Uint64:
		value = reflect.ValueOf(cast.ToUint64(fl.Param()))
	case reflect.Float32:
		value = reflect.ValueOf(cast.ToFloat32(fl.Param()))
	case reflect.Float64:
		value = reflect.ValueOf(cast.ToFloat64(fl.Param()))
	case reflect.Bool:
		value = reflect.ValueOf(cast.ToBool(fl.Param()))
	default:
		return false
	}

	field.Set(value)
	return true
}
