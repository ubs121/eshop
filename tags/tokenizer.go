package tags

import (
	"gopkg.in/mgo.v2/bson"
)

/// обектын tags талбарыг шинэчилнэ
func UpdateTags(obj bson.M) {
	tags := []string{}
	for k, _ := range obj {
		if obj[k] == nil {
			continue
		}
	}

	obj["tags"] = tags
}
