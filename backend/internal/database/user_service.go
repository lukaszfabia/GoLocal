package database

import (
	"backend/internal/forms"
	"backend/internal/models"
	"backend/pkg/parsers"
	"log"

	"gorm.io/gorm"
)

type UserService interface {
	// Creates user if it does not exists in db
	// Returns occured error
	GetOrCreateUser(user *models.User) (*models.User, error)

	// Get user using cond like: "email = joe.doe@example.com", use fmt.Sprintf
	GetUser(conds ...interface{}) (*models.User, error)

	SaveUser(new *forms.EditAccount, old *models.User) (*models.User, error)
	DeleteUser(user *models.User) error
	AddDevice(device *forms.Device) error

	VerifyUser(user models.User) (models.User, error)
	ChangePassword(newPassword string, user models.User) error
}

func NewUserService(db *gorm.DB) UserService {
	return &userServiceImpl{db: db}
}

type userServiceImpl struct {
	db *gorm.DB
}

// ChangePassword implements UserService.
func (u *userServiceImpl) ChangePassword(newPassword string, user models.User) error {
	user.Password = &newPassword

	if err := u.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userServiceImpl) VerifyUser(user models.User) (models.User, error) {
	user.IsVerified = true

	if err := u.db.Save(user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *userServiceImpl) GetOrCreateUser(user *models.User) (*models.User, error) {
	if err := u.db.FirstOrCreate(&user, "email = ?", user.Email).Error; err != nil {
		log.Printf("Couldn't create or find user with email %s: %v", user.Email, err)
		return nil, err
	}
	return user, nil
}

func (s *service) UserService() UserService {
	return s.userService
}

func (u *userServiceImpl) GetUser(conds ...interface{}) (*models.User, error) {
	var res models.User

	if err := u.db.Preload("Devices").First(&res, conds...).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (u *userServiceImpl) SaveUser(new *forms.EditAccount, old *models.User) (*models.User, error) {

	// Assign all form fields to the user model
	old.FirstName = new.FirstName
	old.LastName = new.LastName
	old.Email = new.Email
	old.Password = &new.Password
	old.Bio = &new.Bio

	parsedDate := parsers.ParseDate(new.Birthday)
	old.Birthday = &parsedDate

	if err := u.db.Save(&old).Error; err != nil {
		return nil, err
	}

	newUser := &models.User{}

	if err := u.db.First(newUser, "email = ?", old.Email).Error; err != nil {
		return nil, err
	}

	return newUser, nil
}

// Soft deletes given user
func (u *userServiceImpl) DeleteUser(user *models.User) error {
	if err := u.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userServiceImpl) AddDevice(device *forms.Device) error {
	token := models.DeviceToken{
		Token:     device.Token,
		OSVersion: device.OSVersion,
		Platform:  device.Platform,
	}

	var user models.User
	if err := u.db.Model(&user).First(&user, "id = ?", device.UserID).Error; err != nil {
		log.Println("Error fetching user:", err)
		return err
	}

	u.db.Model(&user).Association("Devices").Append(&token)

	return nil
}
