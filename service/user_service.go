package service

import (
	"auth/config"
	"auth/models"
	"auth/util"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRegisterService struct {
	Nickname string `form:"nickname" json:"nickname"`
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"max=40"`
	Superior uint   `form:"superior" json:"superior"`
	RegCode  string `form:"reg_code" json:"reg_code" binding:"required"`
}

func (service *UserRegisterService) Register() error {
	if service.RegCode != config.REGISTRATION_CODE {
		return errors.New("注册码不正确")
	}

	users := models.GetAllUsers()
	for _, u := range users {
		if u.Username == service.Email {
			return errors.New("账号已存在")
		}
	}

	if len(service.Password) == 0 {
		service.Password = config.DEFAULT_PASSWORD
	}
	var digest, err = util.Digest(service.Password)
	if err != nil {
		return err
	}

	// Assign a new ID (simple auto-increment for demonstration)
	var newID uint = 1
	if len(users) > 0 {
		newID = users[len(users)-1].ID + 1
	}

	user := models.User{
		ID:       newID,
		Username: service.Email,
		Nickname: service.Nickname,
		Password: digest,
	}

	models.AddUser(user)

	if service.Superior != 0 {
		superior := models.Superior{
			UserId:     user.ID,
			SuperiorId: service.Superior,
		}
		models.AddSuperior(superior)
	}

	return nil
}

type UserPasswordService struct {
	Password    string `form:"password" json:"password" binding:"required,min=8,max=40"`
	NewPassword string `form:"newPassword" json:"newPassword" binding:"required,min=8,max=40"`
}

// 修改密码
func (service *UserPasswordService) ChangePassword(c *gin.Context, user models.User) error {
	if !util.CheckPassword(service.Password, user.Password) {
		return errors.New("原密码错误")
	}

	var digest, err = util.Digest(service.NewPassword)
	if err != nil {
		return err
	}

	users := models.GetAllUsers()
	for i, u := range users {
		if u.ID == user.ID {
			users[i].Password = digest
			break
		}
	}
	// This is a simplified approach. In a real application, you'd have a way to update a single user.
	// For now, we'll just re-add all users, which will overwrite the file.
	// A more efficient way would be to have an UpdateUser function in json_store.go
	// For now, we'll just re-add all users, which will overwrite the file.
	// A more efficient way would be to have an UpdateUser function in json_store.go
	// To avoid duplicate entries, clear the existing users and re-add them.
	// This is a temporary workaround for the lack of a proper update function.
	models.ClearUsers()
	for _, u := range users {
		models.AddUser(u)
	}

	return nil
}

func ResetPassword(c *gin.Context) error {
	var digest, err = util.Digest(config.DEFAULT_PASSWORD)
	if err != nil {
		return err
	}
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return errors.New("无效的用户ID")
	}

	users := models.GetAllUsers()
	found := false
	for i, u := range users {
		if u.ID == uint(id) {
			users[i].Password = digest
			found = true
			break
		}
	}

	if !found {
		return errors.New("用户不存在")
	}

	// To avoid duplicate entries, clear the existing users and re-add them.
	// This is a temporary workaround for the lack of a proper update function.
	models.ClearUsers()
	for _, u := range users {
		models.AddUser(u)
	}

	return nil
}

func ChangeSuperior(c *gin.Context) error {
	idStr, superIdStr := c.Param("id"), c.Param("superiorId")

	userId, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return err
	}
	superiorId, err := strconv.ParseUint(superIdStr, 10, 32)
	if err != nil {
		return err
	}

	// Remove existing superior entry for this user
	allSuperiors := models.GetAllSuperiors()
	newSuperiors := []models.Superior{}
	for _, s := range allSuperiors {
		if s.UserId != uint(userId) {
			newSuperiors = append(newSuperiors, s)
		}
	}

	// Add the new superior entry
	newSuperiors = append(newSuperiors, models.Superior{
		UserId:     uint(userId),
		SuperiorId: uint(superiorId),
	})

	// To avoid duplicate entries, clear the existing superiors and re-add them.
	// This is a temporary workaround for the lack of a proper update function.
	models.ClearSuperiors()
	for _, s := range newSuperiors {
		models.AddSuperior(s)
	}

	return nil
}

func DeleteUser(c *gin.Context) error {
	uid := c.GetUint("uid")
	var currentUser models.User
	for _, u := range models.GetAllUsers() {
		if u.ID == uid {
			currentUser = u
			break
		}
	}

	if currentUser.Role != "admin" {
		return errors.New("无权删除")
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return errors.New("无效的用户ID")
	}

	users := models.GetAllUsers()
	newUsers := []models.User{}
	found := false
	for _, u := range users {
		if u.ID != uint(id) {
			newUsers = append(newUsers, u)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("用户不存在")
	}

	// To avoid duplicate entries, clear the existing users and re-add them.
	// This is a temporary workaround for the lack of a proper delete function.
	models.ClearUsers()
	for _, u := range newUsers {
		models.AddUser(u)
	}

	return nil
}

type Staff struct {
	ID         uint   `json:"id"`
	Nickname   string `json:"nickname"`
	Username   string `json:"username"`
	SuperiorId uint   `json:"superiorId"`
	Superior   string `json:"superior"`
	Self       bool   `json:"self"`
}

func GetUserList(uid uint) []Staff {
	var staffList []Staff
	users := models.GetAllUsers()
	superiors := models.GetAllSuperiors()

	for _, user := range users {
		var superiorId uint
		var superiorNickname string

		for _, sup := range superiors {
			if sup.UserId == user.ID {
				superiorId = sup.SuperiorId
				for _, sUser := range users {
					if sUser.ID == superiorId {
						superiorNickname = sUser.Nickname
						break
					}
				}
				break
			}
		}

		staff := Staff{
			ID:         user.ID,
			Nickname:   user.Nickname,
			Username:   user.Username,
			SuperiorId: superiorId,
			Superior:   superiorNickname,
			Self:       false,
		}

		if uid == user.ID {
			staff.Self = true
		}
		staffList = append(staffList, staff)
	}
	return staffList
}

func GetUserListAll() []Staff {
	var staffList []Staff
	users := models.GetAllUsers()
	superiors := models.GetAllSuperiors()

	for _, user := range users {
		var superiorId uint
		var superiorNickname string

		for _, sup := range superiors {
			if sup.UserId == user.ID {
				superiorId = sup.SuperiorId
				for _, sUser := range users {
					if sUser.ID == superiorId {
						superiorNickname = sUser.Nickname
						break
					}
				}
				break
			}
		}

		staff := Staff{
			ID:         user.ID,
			Nickname:   user.Nickname,
			Username:   user.Username,
			SuperiorId: superiorId,
			Superior:   superiorNickname,
			Self:       false,
		}
		staffList = append(staffList, staff)
	}
	return staffList
}

func UserMe(uid uint) models.User {
	for _, u := range models.GetAllUsers() {
		if u.ID == uid {
			return u
		}
	}
	return models.User{}
}

func GetSubList(uid uint) []uint {
	subList := make([]uint, 0)
	allSuperiors := models.GetAllSuperiors()
	for _, superior := range allSuperiors {
		if superior.SuperiorId == uid {
			subList = append(subList, superior.UserId)
		}
	}
	return subList
}
