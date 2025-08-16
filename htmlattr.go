package rxui

import (
	"fmt"

	"github.com/a-h/templ"
)

func Class(name string) templ.KeyValue[string, any] {
	return templ.KeyValue[string, any]{
		Key:   "class",
		Value: name,
	}
}

func ID(name string) templ.KeyValue[string, any] {
	return templ.KeyValue[string, any]{
		Key:   "id",
		Value: name,
	}
}

func Style(name string) templ.KeyValue[string, any] {
	return templ.KeyValue[string, any]{
		Key:   "style",
		Value: name,
	}
}

func Fill(fill string) templ.KeyValue[string, any] {
	return templ.KeyValue[string, any]{
		Key:   "fill",
		Value: fill,
	}
}

func IsChecked() templ.KeyValue[string, any] {
	return templ.KeyValue[string, any]{
		Key:   "checked",
		Value: "checked",
	}
}

func Disabled() templ.KeyValue[string, any] {
	return templ.KeyValue[string, any]{
		Key:   "disabled",
		Value: true,
	}
}

func KV(key string, value any) templ.KeyValue[string, any] {
	return templ.KeyValue[string, any]{
		Key:   key,
		Value: value,
	}
}

func For(id string) templ.KeyValue[string, any] {
	return templ.KeyValue[string, any]{
		Key:   "for",
		Value: id,
	}
}

func Name(name string) templ.KeyValue[string, any] {
	return templ.KeyValue[string, any]{
		Key:   "name",
		Value: name,
	}
}

func MergeClass(builtinClass string, attrlist ...templ.KeyValue[string, any]) templ.OrderedAttributes {

	for i := range attrlist {
		if attrlist[i].Key == "class" {
			attrlist[i].Value = fmt.Sprintf("%s %s", builtinClass, attrlist[i].Value.(string))
			attrs := templ.OrderedAttributes(attrlist)
			return attrs
		}
	}

	attrs := append(attrlist, templ.KeyValue[string, any]{
		Key:   "class",
		Value: builtinClass,
	})

	return attrs
}

func UseDefault(def templ.KeyValue[string, any], attrlist ...templ.KeyValue[string, any]) templ.OrderedAttributes {
	found := false
	for i := range attrlist {
		if attrlist[i].Key == def.Key {
			found = true
			return templ.OrderedAttributes(attrlist)
		}
	}

	if !found {
		attrlist = append(attrlist, def)
	}
	return templ.OrderedAttributes(attrlist)
}
