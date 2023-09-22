package services

import (
	"fmt"
	"math"
	"strings"
	"time"

	weddinginvitation "github.com/gitkuldp/wedding-invitation-api"
	"github.com/gitkuldp/wedding-invitation-api/internal/api/helper"
	"github.com/gitkuldp/wedding-invitation-api/internal/models"
	"github.com/gitkuldp/wedding-invitation-api/internal/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	service
}

// FindUser implements UserService
func (us *userService) FindUser(ctx echo.Context, id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := us.Find(id, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetMe implements UserService
func (us *userService) GetMe(ctx echo.Context) (*models.User, error) {
	user := ctx.Get("user").(models.User)
	return &user, nil
}

// RefreshToken implements UserService
func (us userService) RefreshToken(model models.RefreshTokenView) (*models.Token, error) {
	userID, err := helper.ParseToken(us.db, model.RefreshToken)
	if err != nil {
		return nil, weddinginvitation.InvalidOperation{Message: err.Error()}
	}

	var user models.User
	if err := us.db.Where("id = ? AND NOT is_active", userID).First(&user).Error; err != nil {
		return nil, weddinginvitation.RecordNotFound{Model: "user"}
	}

	roleId := user.GetUserRole(us.db)

	var token *models.Token

	//rotate refresh token
	token, err = helper.GenerateJwtTokens(user.ID, roleId)
	if err != nil {
		return nil, weddinginvitation.InvalidOperation{Message: err.Error()}
	}

	user.TokenLastRevokedAt = time.Unix(token.Iat, 0)

	if err := us.Save(&user); err != nil {
		return nil, err
	}

	return token, nil
}

// UserLogin implements UserService
func (us *userService) UserLogin(view models.UserLoginView) (*models.Token, error) {
	//check if user's email is registered or not
	var user models.User
	if err := us.db.Where("email = ? AND NOT is_active", strings.ToLower(view.Email)).First(&user).Error; err != nil {
		return nil, weddinginvitation.RecordNotFound{Model: "user"}
	}

	//check if the account is locked or not
	if time.Now().Before(user.LockedUntil) {
		return nil, weddinginvitation.InvalidOperation{Message: fmt.Sprintf("your account is locked for '%s'", utils.ConvertMinutes(user.LockedDuration))}
	}

	//check password
	if err := CheckPassword(view, user, us.db); err != nil {
		return nil, err
	}

	//get role id
	roleId := user.GetUserRole(us.db)

	token, err := helper.GenerateJwtTokens(user.ID, roleId)
	if err != nil {
		return nil, weddinginvitation.InvalidOperation{Message: err.Error()}
	}

	//keep track of the time when the refresh token was last provided
	user.TokenLastRevokedAt = time.Unix(token.Iat, 0)

	//on successful attempt intialize to 0
	user.FailedAttempt = 0
	user.LockedDuration = 0

	if err := us.Save(&user); err != nil {
		return nil, err
	}

	return token, nil
}

// ListUsers implements UserService
func (us *userService) ListUsers(models.User) ([]*models.User, error) {
	var users []*models.User
	err := us.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// RegisterUser implements UserService
func (us *userService) RegisterUser(name models.UserView) (*models.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(name.Password), 10)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Phone:    name.Phone,
		ID:       uuid.New(),
		Email:    name.Email,
		Password: string(hashPass),
	}

	err = us.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type UserService interface {
	RegisterUser(models.UserView) (*models.User, error)
	ListUsers(models.User) ([]*models.User, error)
	UserLogin(view models.UserLoginView) (*models.Token, error)
	RefreshToken(model models.RefreshTokenView) (*models.Token, error)
	FindUser(ctx echo.Context, id uuid.UUID) (*models.User, error)
	GetMe(ctx echo.Context) (*models.User, error)
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		service{
			db: db,
		},
	}
}

func CheckPassword(view models.UserLoginView, user models.User, db *gorm.DB) error {
	// Check if the locked duration has passed and reset the condition
	zeroTime := time.Time{}
	if user.LockedUntil.Format(utils.DateFormatYMD) != zeroTime.Format(utils.DateFormatYMD) {
		if user.LockedUntil.Add(2 * time.Minute).Before(time.Now()) {
			user.LockedUntil = time.Time{} // Reset locked until time
			user.FailedAttempt = 0         // Reset failed attempt count
			user.LockedDuration = 0        // Reset locked duration
			if err := db.Save(&user).Error; err != nil {
				return err
			}
		}
	}

	//check password
	if ok := utils.CheckPasswordHash(view.Password, user.Password); !ok {
		//increment the failed attempt if password is not correct
		user.FailedAttempt = user.FailedAttempt + 1

		//if failed attempt exceeds given max attempt, block the account for some time
		if user.FailedAttempt > utils.MaxLoginAttempt {
			var lockedTimeInMinutes int
			//get default locked duration
			if user.LockedDuration == 0 {
				lockedTimeInMinutes = utils.LockedTimeInMinutes
			} else {
				//exponential increment
				lockedTimeInMinutes = int(math.Exp(float64(user.LockedDuration)))
			}
			//set the account blocked time
			user.LockedUntil = time.Now().Add(time.Duration(lockedTimeInMinutes) * time.Minute)
			user.FailedAttempt = 0 //to be able to check the unsuccessful attempts again
			user.LockedDuration = lockedTimeInMinutes
			if err := db.Save(&user).Error; err != nil {
				return err
			}
			return weddinginvitation.InvalidOperation{Message: fmt.Sprintf("your account has been locked for '%s'", utils.ConvertMinutes(lockedTimeInMinutes))} //show minutes data
		}

		if err := db.Save(&user).Error; err != nil {
			return err
		}

		return weddinginvitation.InvalidOperation{Message: "invalid username or password"}
	}

	return nil
}
