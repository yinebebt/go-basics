package Reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// fn("some string passed") //input= "some string passed"
	// fn("argument-2")

	val := getValue(x)
	//inoreder to make walk only extract value, let we DRY it by assigning to walkValue
	// walkValue := func(value reflect.Value) {
	// 	walk(value.Interface(), fn)
	// 	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ { //NumField used only for a struct
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array: // two possible cases
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn) //Interface returns v's current value as an interface{}.
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}

	case reflect.String: //this is the case where, all after the recusives comes to the end-case string which is passed to fn
		fn(val.String())
	}
}

func getValue(a any) reflect.Value {
	return reflect.ValueOf(a) //ValueOf returns a new Value initialized to the concrete value stored in the interface a.
}

//it is also possible to know the field length for all and make only one single fr loog.
//I skip since this one gives me more clearer to understand.

/*
numberOfValues := 0
var getField func(int) reflect.Value
...
numberOfValues = val.NumField()
getField = val.Field
...
for i=0;. i<numberOfValues; i++{
	...
}
*/
