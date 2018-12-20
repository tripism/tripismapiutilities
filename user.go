package tripismapiutilities

import (
	"github.com/globalsign/mgo/bson"
	"github.com/tripism/api/types"
)

// **************************************************************************************
// NOTE: Interface changes to the InXXX interfaces should also be reflected in the file
// core.users.go
// **************************************************************************************

// UserToUserInNotification converts a user (public) type to UserInNotification type
func UserToUserInNotification(user *types.User) *types.UserInNotification {
	u := &types.UserInNotification{
		ID:            user.ID,
		Name:          user.Name,
		TokenName:     user.TokenName,
		Nickname:      user.Nickname,
		TokenNickname: user.TokenNickname,
		AvatarURL:     user.AvatarURL,
		Orgs:          user.Orgs,
	}
	return u
}

// UserToAdminUserDetails converts a User type to AdminUserDetails type
func UserToAdminUserDetails(user *types.User) *types.AdminUserDetails {
	u := &types.AdminUserDetails{
		ID:                user.ID,
		Name:              user.Name,
		Nickname:          user.Nickname,
		Username:          user.Username,
		UserPrincipalName: user.UserPrincipalName,
		Description:       user.Description,
		Location:          user.Location,
		Email:             user.Email,
		AvatarURL:         user.AvatarURL,
		AccessControl:     user.AccessControl,
		AccessFeatures:    user.AccessFeatures,
		Orgs:              user.Orgs,
		Achievements:      user.Achievements,
		Joined:            user.Joined,
		LastLoggedIn:      user.LastLoggedIn,
		PreviousLoggedIn:  user.PreviousLoggedIn,
	}
	return u
}

// UserInCardToUserInNotification converts a user in UserInCard type to UserInNotification type
func UserInCardToUserInNotification(user *types.UserInCard) *types.UserInNotification {
	u := &types.UserInNotification{
		ID:            user.ID,
		Name:          user.Name,
		TokenName:     user.TokenName,
		Nickname:      user.Nickname,
		TokenNickname: user.TokenNickname,
		AvatarURL:     user.AvatarURL,
		Orgs:          user.Orgs,
	}
	return u
}

// UserInTripBoardToUserInNotification converts a user in UserInTripBoard type to UserInNotification type
func UserInTripBoardToUserInNotification(user *types.UserInTripBoard) *types.UserInNotification {
	u := &types.UserInNotification{
		ID:            user.ID,
		Name:          user.Name,
		TokenName:     user.TokenName,
		Nickname:      user.Nickname,
		TokenNickname: user.TokenNickname,
		AvatarURL:     user.AvatarURL,
		Orgs:          user.Orgs,
	}
	return u
}

// UserInActivityToUserInNotification converts a user in UserInActivity type to UserInNotification type
func UserInActivityToUserInNotification(user *types.UserInActivity) *types.UserInNotification {
	u := &types.UserInNotification{
		ID:            user.ID,
		Name:          user.Name,
		TokenName:     user.TokenName,
		Nickname:      user.Nickname,
		TokenNickname: user.TokenNickname,
		AvatarURL:     user.AvatarURL,
		Orgs:          user.Orgs,
	}
	return u
}

// getUserOrgIDs returns user org IDs
func GetUserOrgIDs(user *types.User) []bson.ObjectId {
	if user == nil {
		return []bson.ObjectId{}
	}
	userOrgIDs := make([]bson.ObjectId, len(user.Orgs))
	for i, org := range user.Orgs {
		userOrgIDs[i] = org.ID
	}
	return userOrgIDs
}
