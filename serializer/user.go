package serializer

import "auth/models"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	
}

func BuildUser(user models.User) User {
	return User{
		ID:        user.ID,
		Username:  user.Username,
		Nickname:  user.Nickname,
	}
}

func UserRsp(user models.User) Response {
	return Response{
		Code: 0,
		Msg:  "success",
		Data: BuildUser(user),
	}
}
