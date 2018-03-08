package polls

import (
	"github.com/alphapeter/letsvote/server/users"
	"net/http"
)

type unautorizedFetch struct {
}

func (unautorizedFetch) message() string {
	return "unauthorized"
}

func (unautorizedFetch) responseCode() int {
	return http.StatusUnauthorized
}

func hasPermissionToEdit(created userCreated, user users.User) bool {
	return user.IsAdmin || created.getUserId() == user.Id
}
