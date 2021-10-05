package helpers

import (
	"fmt"
	"regexp"
	"strconv"
)

//Not Zero
func NotZero(val int, fieldName string) error {
	if val == 0 {
		return fmt.Errorf(fieldName + " is required, cannot be empty")
	}
	return nil
}

//Not Empty
func MustNotEmpty(val, fieldName string) error {
	if val == "" {
		return fmt.Errorf(fieldName + " is required, cannot be empty")
	}
	return nil
}

//Not Empty, 5-20 characters
func UsernameRule(val string) error {
	minLength := 5
	maxLength := 20

	if val == "" {
		return fmt.Errorf("username is required, cannot be empty")
	}

	if len(val) < minLength || len(val) > maxLength {
		return fmt.Errorf("username need %d-%d characters", minLength, maxLength)
	}
	return nil
}

//Not Empty, 5-45 characters
func PasswordRule(val string) error {
	minLength := 5
	maxLength := 45

	if val == "" {
		return fmt.Errorf("password is required, cannot be empty")
	}

	if len(val) < minLength || len(val) > maxLength {
		return fmt.Errorf("password need %d-%d characters", minLength, maxLength)
	}
	return nil
}

//Not Empty, 5-50 characters
func EmailRule(val string) error {

	minLength := 5
	maxLength := 50
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if val == "" {
		return fmt.Errorf("email is required, cannot be empty")
	}

	if len(val) < minLength || len(val) > maxLength {
		return fmt.Errorf("email need %d-%d characters", minLength, maxLength)
	}

	if !emailRegex.MatchString(val) {
		return fmt.Errorf("%v: invalid email address", val)
	}
	return nil
}

//Not Empty, Must Number, 5-20 characters
func PhoneRule(val string) error {
	minLength := 5
	maxLength := 45

	if _, err := strconv.Atoi(val); err != nil {
		return fmt.Errorf("phone must be numeric")
	}

	if val == "" {
		return fmt.Errorf("phone number is required, cannot be empty")
	}

	if len(val) < minLength || len(val) > maxLength {
		return fmt.Errorf("phone number need %d-%d characters", minLength, maxLength)
	}
	return nil
}

//Not Empty, 1-50 characters
func FullnameRule(val string) error {
	minLength := 3
	maxLength := 50

	if val == "" {
		return fmt.Errorf("full name is required, cannot be empty")
	}

	if len(val) < minLength || len(val) > maxLength {
		return fmt.Errorf("full name need %d-%d characters", minLength, maxLength)
	}

	return nil
}

func NameRule(val string) error {
	minLength := 3
	maxLength := 50

	if val == "" {
		return fmt.Errorf("name is required, cannot be empty")
	}

	if len(val) < minLength || len(val) > maxLength {
		return fmt.Errorf("name need %d-%d characters", minLength, maxLength)
	}

	return nil
}

//Not Empty , 3-5 characters
func BankAccountNameRule(val string) error {
	minLength := 3
	maxLength := 50

	if val == "" {
		return fmt.Errorf("account name is required, cannot be empty")
	}

	if len(val) < minLength || len(val) > maxLength {
		return fmt.Errorf("account name need %d-%d characters", minLength, maxLength)
	}
	return nil
}

//Not Empty, Must Number, 8-30 characters
func BankAccountNumberRule(val string) error {
	minLength := 8
	maxLength := 30

	if _, err := strconv.Atoi(val); err != nil {
		return fmt.Errorf("account number must be numeric")
	}

	if val == "" {
		return fmt.Errorf("account number is required, cannot be empty")
	}

	if len(val) < minLength || len(val) > maxLength {
		return fmt.Errorf("account number need %d-%d characters", minLength, maxLength)
	}
	return nil
}

//Not Empty, Must Number, minimum 5 - maximum 1000000
func AmountDepoRule(val int) error {
	minDepo := 5
	maxDepo := 1000000

	if val == 0 {
		return fmt.Errorf("deposit amount is required, cannot be empty")
	}

	if val < minDepo || val > maxDepo {
		return fmt.Errorf("your deposit need %d-%d rupiah", minDepo, maxDepo)
	}
	return nil
}

func WithdrawRule(userCredit, requestWithdraw int) error {

	if userCredit == 0 {
		return fmt.Errorf("withdraw amount is required, cannot be empty")
	}

	if userCredit < requestWithdraw {
		return fmt.Errorf("your balance is not enough to withdraw")
	}

	return nil
}
