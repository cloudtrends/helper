package helper

import (
	//"html"
	"html/template"
	"log"
	"reflect"
	//"os"
	"bytes"
	//"regexp"
	//"strconv"
	"strings"
	//"time"
)

var ERROR_CLASS = "hasError"

func init() {
	log.Printf("Begin init framework tempalte")
}
func DynamicRender(template_name, template_value string) (str string, is_ok bool, err_str string) {
	var doc bytes.Buffer
	if t, err := template.New(template_name).Funcs(TemplateFuncs).Parse(template_value); nil != err || t == nil {
		log.Printf("Error when dynamic render :%s", err.Error())
		str = ""
		err_str = err.Error()
		is_ok = false
		log.Printf("Template parser failling")
	} else {
		t.Execute(&doc, "")
		if 0 == doc.Len() {
			str = ""
			is_ok = false
			err_str = "May be function not defined"
			log.Printf("Template parser ok, but doc is nil")
		} else {
			if 0 == doc.Len() {
				is_ok = false
				err_str = "May be function not defined"
				str = ""
			} else {
				str = doc.String()
				is_ok = true
				log.Printf("Template parser ok:%s", str)
			}
			//if nil == nil.(str) {
			//	str = ""
			//}
		}
	}
	//t.Execute(os.Stdout, "Joseki")
	//str = ""
	// "<h1>{{.}}</h1>{{label \"alex\"}}"
	return
}

var (
	// The functions available for use in the templates.
	TemplateFuncs = map[string]interface{}{
		"label": strings.Title,
		"eq":    Equal,
		"set": func(renderArgs map[string]interface{}, key string, value interface{}) template.HTML {
			renderArgs[key] = value
			return template.HTML("")
		},
		"append": func(renderArgs map[string]interface{}, key string, value interface{}) template.HTML {
			if renderArgs[key] == nil {
				renderArgs[key] = []interface{}{value}
			} else {
				renderArgs[key] = append(renderArgs[key].([]interface{}), value)
			}
			return template.HTML("")
		},

		// Replaces newlines with <br>
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
		},

		// Skips sanitation on the parameter.  Do not use with dynamic data.
		"raw": func(text string) template.HTML {
			return template.HTML(text)
		},
	}
)

// Equal is a helper for comparing value equality, following these rules:
//  - Values with equivalent types are compared with reflect.DeepEqual
//  - int, uint, and float values are compared without regard to the type width.
//    for example, Equal(int32(5), int64(5)) == true
//  - strings and byte slices are converted to strings before comparison.
//  - else, return false.
func Equal(a, b interface{}) bool {
	if reflect.TypeOf(a) == reflect.TypeOf(b) {
		return reflect.DeepEqual(a, b)
	}
	switch a.(type) {
	case int, int8, int16, int32, int64:
		switch b.(type) {
		case int, int8, int16, int32, int64:
			return reflect.ValueOf(a).Int() == reflect.ValueOf(b).Int()
		}
	case uint, uint8, uint16, uint32, uint64:
		switch b.(type) {
		case uint, uint8, uint16, uint32, uint64:
			return reflect.ValueOf(a).Uint() == reflect.ValueOf(b).Uint()
		}
	case float32, float64:
		switch b.(type) {
		case float32, float64:
			return reflect.ValueOf(a).Float() == reflect.ValueOf(b).Float()
		}
	case string:
		switch b.(type) {
		case []byte:
			return a.(string) == string(b.([]byte))
		}
	case []byte:
		switch b.(type) {
		case string:
			return b.(string) == string(a.([]byte))
		}
	}
	return false
}
