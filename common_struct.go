package helper

import (
	//"encoding/json"
	"log"
	//"reflect"
	//"strconv"
	//"reflect"
	"strings"
)

type JsonTopicContent struct {
	// any thing save into topic content is a json
	// how about the version of JsonTopicContent ?
	// The ContentType is a flag for identify how to use json .
	// ContentType
	// "" empty or "plain_text" is for txt
	// "html" for html string
	// json_xxx is for json  class name , and xxx is stand for which json class
	//
	FieldList   FieldList
	Content     string
	ContentType string
}
type TopicField struct {
	Name     string
	ShowName string
	Value    string
	Type     string // string, hidden
}
type FieldList struct {
	Fields []TopicField
}

const (
	JNamePrefix     = "JName_"
	JShowNamePrefix = "JShowName_"
)

//obj interface{},
func LoadFormValueAndParserJsonTopicContent(jc JsonTopicContent, m map[string][]string) JsonTopicContent {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("Panic! %v\n", e)
		}
	}()

	// val := reflect.ValueOf(obj)
	// if val.Kind() == reflect.Ptr {
	//  val = val.Elem()
	// }

	// Loop over map, try to match the key to a field
	var name_array []string
	//var show_name_array []string
	for k, _ := range m {
		log.Println("origin key:" + k)
		if strings.HasPrefix(k, JNamePrefix) {
			log.Println("find k:" + k)
			name_array = append(name_array, strings.Replace(k, JNamePrefix, "", -1))
		} else if strings.HasPrefix(k, "show_") {
			//show_name_array = append(show_name_array, strings.Replace(k, "show_", "", -1))
		}

		// if f := val.FieldByName(k); f.IsValid() {
		//  // Is it assignable?
		//  if f.CanSet() {

		//      // Assign the map's value to this field, converting to the right data type.
		//      switch f.Type().Kind() {
		//      // Only a few kinds, just to show the basic idea...
		//      case reflect.Int:
		//          if i, e := strconv.ParseInt(v, 0, 0); e == nil {
		//              f.SetInt(i)
		//          } else {
		//              fmt.Printf("Could not set int value of %s: %s\n", k, e)
		//          }
		//      case reflect.Float64:
		//          if fl, e := strconv.ParseFloat(v, 0); e == nil {
		//              f.SetFloat(fl)
		//          } else {
		//              fmt.Printf("Could not set float64 value of %s: %s\n", k, e)
		//          }
		//      case reflect.String:
		//          f.SetString(v)

		//      default:
		//          fmt.Printf("Unsupported format %v for field %s\n", f.Type().Kind(), k)
		//      }
		//  } else {
		//      fmt.Printf("Key '%s' cannot be set\n", k)
		//  }
		// } else {
		//  // Key does not map to a field in obj
		//  //fmt.Printf("Key '%s' does not have a corresponding field in obj %+v\n", k, obj)
		// }
	}

	//
	for _, n := range name_array {
		log.Println("list name:" + n)
		//log.Println("list name value:" + m[n])
		//
		tf := TopicField{}
		tf.Name = n
		if 1 == len(m[n]) {
			tf.Value = m[n][0]

		}
		if 1 == len(m[JShowNamePrefix+n]) {
			tf.ShowName = m[JShowNamePrefix+n][0]
		}
		jc.FieldList.Fields = append(jc.FieldList.Fields, tf)
		//log.Println("reflect type:", reflect.TypeOf(m[n]))
	}

	// for _, n := range show_name_array {
	//  log.Println("list show name:" + n)
	//  log.Println("list show name value:" + m[n])
	//  //
	//  tf := TopicField{}
	//  tf.Name = n
	//  tf.Value = m[n]
	//  tf.ShowName = m[JShowNamePrefix+n]
	//  jc.FieldList.Fields = append(jc.FieldList.Fields, tf)
	//  //

	// }
	return jc
}

func main2() {
	// m2 := map[string][]string{
	//  //"name_Age":    "Age",

	//  JShowNamePrefix + "Age": "年龄",
	//  JNamePrefix + "Age":     "Age",
	//  "Age":                   "36",
	//  //
	//  JNamePrefix + "Name":   "Name",
	//  "Name":                 "Johnny",
	//  JNamePrefix + "Salary": "Salary",
	//  "Salary":               "1400.33",
	//  "Ignored":              "True",
	// }
	// //p := new(TopicField)
	// //fl := FieldList{}
	// jc := JsonTopicContent{}
	// //

	// jc = LoadFormValueAndParserJsonTopicContent(jc, m2)
	// //fmt.Printf("After LoadModel: Person=%+v\n", p)
	// for _, field_one := range jc.FieldList.Fields {
	//  log.Println("Final Result:")
	//  log.Println(field_one.Name + "=" + field_one.Value + "=" + field_one.Type)
	// }

	// b, err := json.Marshal(jc)
	// if err != nil {
	//  log.Println("Error when marshal :")
	//  log.Println(err)
	// }
	// log.Println("json result:" + string(b))

}
