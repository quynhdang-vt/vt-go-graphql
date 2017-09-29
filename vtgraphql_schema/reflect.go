package vtgraphql_schema

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

// usage
//
// IntrospectType (reflect.TypeOf( (*SchemaQueryResponse)(nil)).Elem())

//helper function to construct struct with all nested so that we can pass on to json  unmarshaller
func IntrospectType(t reflect.Type, indent string) {
	if t == nil {
		return
	}

	fmt.Printf("%s%s -- kind=%v ", indent, t.Name(), t.Kind())
	if t.Kind() == reflect.Slice {
		// it's an array so we can't really do NumFields
		fmt.Printf("%s..\n", t.Elem().Name())
		IntrospectType(t.Elem(), indent+"  ")
	} else if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			//f:=s.Field(i)
			fmt.Printf("%sfieldName=%s, fieldType=%v, fieldTag(json)=%v\n", indent, t.Field(i).Name, t.Field(i).Type, GetTagValues(t.Field(i).Tag, "json"))
			IntrospectType(t.Field(i).Type, indent+"  ")
		}
	}
}
func GetTagValues(tag reflect.StructTag, tagName string) []string {
	if len(tag) > 0 {
		return strings.Split(tag.Get(tagName), ",")
	}
	return nil
}

func InitializeASchemaQueryResponseStruct() *SchemaQueryResponse {
	s := &SchemaQueryResponse{}
	InitializeStruct("  ", "ROOT", reflect.TypeOf(*s), reflect.ValueOf(s).Elem())
	return s
}

func InitializeStruct(indent string, fieldName string, t reflect.Type, v reflect.Value) {
	//	defer func() { if r:=recover(); r!=nil {fmt.Printf("%s ***** Panicking in %s *****\n", indent, fieldName)} }()
	fmt.Printf("%s initializeStruct ENTER:: fieldName=%s, [T: type=%v, kind=%v] [V: type=%v, kind=%v, NumField=%v, settable=%v, addressable=%v]\n",
		indent,
		fieldName,
		t.Name(), t.Kind(),
		v.Type().Name(), v.Kind(), v.NumField(),
		v.CanSet(),

		v.CanAddr())

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		fmt.Printf("%s initializeStruct:: [FT: fieldName=%v type=%v, kind=%v] [F(V): type=%v, settable=%v, addressable=%v]\n",
			indent,
			ft.Name, ft.Type.Name(), ft.Type.Kind(),
			f.Type().Name(), f.CanSet(), f.CanAddr())
		switch ft.Type.Kind() {
		case reflect.Map:
			fmt.Printf("%s MAP\n", indent)
			f.Set(reflect.MakeMap(ft.Type))
		case reflect.Slice:
			fmt.Printf("%s SLICE\n", indent)

			item := reflect.MakeSlice(ft.Type, 0, 5)
			if ft.Type.Elem().Kind() == reflect.Struct {

				newItem := reflect.New(ft.Type.Elem())
				InitializeStruct(indent+"  ", " An item of array -- "+ft.Name, ft.Type.Elem(), newItem.Elem()) //reflect.ValueOf(newItem))
				/** DEBUG
				if strings.Contains(string(ft.Type.Elem().Name()), "GraphQLFieldStruct") {
					gn := newItem.Interface().(*GraphQLFieldStruct)
					fmt.Printf("Try one item to see if it's there %v\n", gn)
				}
				*/
				item = reflect.Append(item, newItem.Elem())
			}
			f.Set(item)
			//
		case reflect.Chan:
			fmt.Printf("%s CHANNEL\n", indent)

			f.Set(reflect.MakeChan(ft.Type, 0))
		case reflect.Struct:
			fmt.Printf("%s STRUCT\n", indent)

			InitializeStruct(indent+"  ", "A struct of -- "+ft.Name, ft.Type, f)
		case reflect.Ptr:
			fmt.Printf("%s PTR  >> FT (%v) type=%v, f type=%v\n", indent, ft.Name, ft.Type.Elem(), reflect.TypeOf(f).Kind())
			fv := reflect.New(ft.Type.Elem())
			fmt.Printf("%s fv -- just allocated, type = %v, kind=%v\n", indent, fv.Elem().Type().Name(), fv.Elem().Type().Kind())
			// only move forward if not scalar
			if fv.Elem().Kind()==reflect.Struct {
				InitializeStruct(indent+"  ", "A pointer of -- "+ft.Name, ft.Type.Elem(), fv.Elem()) //reflect.ValueOf(&fv))
			}
			f.Set(fv)
			fmt.Printf("%s << (%v) PTR\n", indent, ft.Name)
			//f.Set(fv.Elem())
		default:
			fmt.Printf("%s ????\n", indent)
		}
	}
	fmt.Printf("%s initializeStruct::EXIT fieldName=%s, %v\n", indent, fieldName, t.Name())
}

func WriteToFile(fileName string, sdata string) (err error) {
	log.Printf("WriteToFile %s\n", fileName)
	f, err := os.Create(fileName)
	defer f.Close()
	if err == nil {
		_, err = f.WriteString(sdata)
	}
	return err
}

/** given a pointer to a struct,
func checkNils(x interface{}) bool {
	v := reflect.ValueOf(x)
        if v.Kind() != reflect.Struct {
                 return false
         }
         for i := 0; i < v.NumField(); i++ {
                 f := v.Field(i)
                 if f.Kind() != reflect.Ptr {
                         continue
                 }
                 if f.IsNil() {
                         return false
                 }
         }
         return true
}
*/