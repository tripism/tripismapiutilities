package tripismapiutilities

import (
	"errors"

	"github.com/globalsign/mgo/bson"
)

// Sort2BSON converts sort order from []string{"-likes", "title"} to what Mongo accepts - { $sort: { likes: -1, title: 1 } }
// Stolen from http://bazaar.launchpad.net/+branch/mgo/v2/view/head:/session.go#L2130
func Sort2BSON(fields []string) (order bson.D, err error) {
	if len(fields) == 0 {
		return
	}
	for _, field := range fields {
		n := 1
		if field != "" {
			switch field[0] {
			case '+':
				field = field[1:]
			case '-':
				n = -1
				field = field[1:]
			}
		}
		if field == "" {
			err = errors.New("Sort: empty field name")
			return
		}
		order = append(order, bson.DocElem{Name: field, Value: n})
	}
	return
}
