package polls

import (
	"net/http"
	"github.com/alphapeter/letsvote/server/users"
)

type unautorizedFetch struct {

}

func (unautorizedFetch) message() string{
	return "unauthorized"
}

func (unautorizedFetch) responseCode() int{
	return http.StatusUnauthorized
}

func hasPermissionToEdit(created userCreated, user users.User) bool {
	return created.getUserId() == user.Id
}
