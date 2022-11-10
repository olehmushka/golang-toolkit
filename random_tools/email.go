package randomtools

import (
	"fmt"

	"github.com/olehmushka/golang-toolkit/wrapped_error"
)

func RandEmail() (string, error) {
	username, err := RandString(100)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not random username for email")
	}

	domainname, err := RandString(5)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not random domainname for email")
	}

	extension, err := RandString(3)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not random extension for email")
	}

	return fmt.Sprintf("%s@%s.%s", username, domainname, extension), nil
}

func GetRandomEmails(len int) ([]string, error) {
	emails := make([]string, len)

	for i := 0; i < len; i++ {
		email, err := RandEmail()
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not random email for emails generation")
		}
		emails[i] = email
	}

	return emails, nil
}
