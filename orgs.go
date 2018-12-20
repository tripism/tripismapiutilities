package tripismapiutilities

import (
	"github.com/globalsign/mgo/bson"
	"github.com/tripism/api/types"
)

// Orgs2IDs gets organizations IDs from []Orgs
func Orgs2IDs(orgs []*types.Org) []bson.ObjectId {
	if orgs == nil {
		return []bson.ObjectId{}
	}
	orgIDs := make([]bson.ObjectId, len(orgs))
	for i, org := range orgs {
		orgIDs[i] = org.ID
	}
	return orgIDs
}
