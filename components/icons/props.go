package icons

import (
	"github.com/a-h/templ"
)

func IconAttrs(attrlist ...templ.KeyValue[string, any]) templ.OrderedAttributes {
	fillFound := false
	for i := range attrlist {
		if attrlist[i].Key == "fill" {
			fillFound = true
			break
		}
	}

	if !fillFound {
		attrlist = append(attrlist, templ.KeyValue[string, any]{Key: "fill", Value: "currentColor"})
	}
	attrs := templ.OrderedAttributes(attrlist)
	return attrs
}
