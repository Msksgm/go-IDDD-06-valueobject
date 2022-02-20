package identity

import (
	"fmt"
	"unicode"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	tenantId TenantId
	userName string
	password string
}

const STRONG_THRESHOL = 20

func NewUser(tenantId TenantId, userName string, password string) (_ *User, err error) {
	user := new(User)

	user.tenantId = tenantId

	if err := user.setUserName(userName); err != nil {
		return nil, err
	}

	bcryptedPassword, err := user.protectPassword("", password)
	if err != nil {
		return nil, err
	}
	user.password = bcryptedPassword

	return user, nil
}

func (user *User) setUserName(userName string) (err error) {
	defer ierrors.Wrap(&err, "user.setUserName(%s)", userName)
	if userName == "" {
		return fmt.Errorf("The username is required.")
	}
	if len(userName) < 3 || len(userName) > 250 {
		return fmt.Errorf("The username must be 3 to 250 characters.")
	}
	user.userName = userName
	return nil
}

func (user *User) protectPassword(currentPassword string, changedPassword string) (string, error) {
	if err := user.assertPasswordNotSame(currentPassword, changedPassword); err != nil {
		return "", err
	}

	if changedPassword == "" {
		return "", fmt.Errorf("The password must not be empty")
	}

	strength := 0

	length := len(changedPassword)

	if length > 7 {
		strength += 10
		// bonus: one point each additional
		strength += (length - 7)
	}
	digitCount, letterCount, lowerCount, upperCount, symbolCount := 0, 0, 0, 0, 0
	for _, ch := range changedPassword {
		if unicode.IsLetter(ch) {
			letterCount++
			if unicode.IsUpper(ch) {
				upperCount++
			} else {
				lowerCount++
			}
		} else if unicode.IsDigit(ch) {
			digitCount++
		} else {
			symbolCount++
		}
	}

	strength += (upperCount + lowerCount + symbolCount)

	// bonus: letters and digits
	if letterCount >= 2 && digitCount >= 2 {
		strength += (letterCount + digitCount)
	}

	if strength < STRONG_THRESHOL {
		return "", fmt.Errorf("The password must be stronger.")
	}

	if err := user.assertUsernamePasswordNotSame(changedPassword); err != nil {
		return "", err
	}

	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(changedPassword), 12)
	if err != nil {
		return "", err
	}

	return string(bcryptedPassword), nil
}

func (user *User) assertPasswordNotSame(currentPassword string, changedPassword string) (err error) {
	defer ierrors.Wrap(&err, "user.assertPasswordNotSame(%s, %s)", currentPassword, changedPassword)
	if currentPassword == changedPassword {
		return fmt.Errorf("The password is unchanged")
	}
	return nil
}

func (user *User) assertUsernamePasswordNotSame(changedPassword string) (err error) {
	if changedPassword == user.userName {
		return fmt.Errorf("The username and password must not be the same.")
	}
	return nil
}

func (user *User) Equals(other User) bool {
	return user.tenantId.id == other.tenantId.id
}
