package prompt

import (
	"fmt"
	"gosh/models"
	"os"
	"strings"
)

func InitUser() (*models.User, error) {
	var user models.User
	var err error

	user.Username, err = os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}
	user.Hostname, err = os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("failed to get hostname: %w", err)
	}
	user.Pwd, err = os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	return &user, nil
}

func Show(usr *models.User, flag bool) {
	if flag && usr != nil {
		fmt.Printf("%v@%v-%v ->", strings.TrimPrefix(usr.Username, "/home/"), usr.Hostname, usr.Pwd)
	} else {
		fmt.Print("->")
	}
}
