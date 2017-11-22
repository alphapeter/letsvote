package users

import (
	"github.com/alphapeter/letsvote/server/config"
	"github.com/jinzhu/gorm"
)

func SetSession(user User, sessionId string) error{
	session := Session{
		Id:     sessionId,
		UserId: user.Id,
	}
	return config.DB.Save(&session).Error
}

func DeleteSession(sessionId string) {
	config.DB.Delete(&Session{}, "id = ?", sessionId)
}

func GetOrCreateUser(user User) (User, error) {
	user.Gravatar = EmailToHash(user.Email)
	dbUser := User{}
	err := config.DB.First(&dbUser, "id = ?", user.Id).Error

	if err == gorm.ErrRecordNotFound {
		err = config.DB.Save(&user).Error
		return user, err
	}

	err = config.DB.Model(dbUser).Update(user).Error

	if err != nil {
		return user, err
	}

	return dbUser, err
}
