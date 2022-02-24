package user

import (
	"fmt"
	"unicode"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/tenantid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	tenantId tenantid.TenantId
	userName string
	password string
}

const STRONG_THRESHOL = 20

func NewUser(tenantId tenantid.TenantId, userName string, password string) (_ *User, err error) {
	user := new(User)

	user.tenantId = tenantId

	if err := user.setUserName(userName); err != nil {
		return nil, err
	}

	if err := user.protectPassword("", password); err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) setUserName(userName string) (err error) {
	defer ierrors.Wrap(&err, "user.setUserName(%s)", userName)
	if err := ierrors.NewArgumentNotEmptyError(userName, "First name is required.").GetError(); err != nil {
		return err
	}
	if err := ierrors.NewArgumentLengthError(userName, 3, 250, "The username must be 3 to 250 characters.").GetError(); err != nil {
		return err
	}
	user.userName = userName
	return nil
}

func (user *User) protectPassword(currentPassword string, changedPassword string) error {
	if err := user.assertPasswordNotSame(currentPassword, changedPassword); err != nil {
		return err
	}

	if err := user.assertPasswordNotWeak(changedPassword); err != nil {
		return err
	}

	if err := user.assertUsernamePasswordNotSame(changedPassword); err != nil {
		return err
	}

	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(changedPassword), 12)
	if err != nil {
		return err
	}

	user.password = string(bcryptedPassword)
	return nil
}

func (user *User) assertPasswordNotSame(currentPassword string, changedPassword string) (err error) {
	defer ierrors.Wrap(&err, "user.assertPasswordNotSame(%s, %s)", currentPassword, changedPassword)
	if currentPassword == changedPassword {
		return fmt.Errorf("The password is unchanged")
	}
	return nil
}

func (user *User) assertPasswordNotWeak(changedPassword string) (err error) {
	defer ierrors.Wrap(&err, "user.assertPasswordNotWeak(%s)", changedPassword)

	if changedPassword == "" {
		return fmt.Errorf("The password must not be empty")
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
		return fmt.Errorf("The password must be stronger.")
	}

	return nil
}

func (user *User) assertUsernamePasswordNotSame(changedPassword string) (err error) {
	defer ierrors.Wrap(&err, "user.assertUsernamePasswordNotSame(%s)", changedPassword)
	if changedPassword == user.userName {
		return fmt.Errorf("The username and password must not be the same.")
	}
	return nil
}

func (user *User) Equals(other User) bool {
	return user.tenantId == other.tenantId
}
