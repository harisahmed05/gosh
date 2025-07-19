package prompt

import (
	"fmt"
	"gosh/models"
	"os"
	"os/user"
	"strings"
)

func InitUser() (*models.User, error) {
	var usr models.User

	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	usr.Username = currentUser.Username

	usr.Homedir, err = os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}
	usr.Hostname, err = os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("failed to get hostname: %w", err)
	}
	usr.Pwd, err = os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	return &usr, nil
}

func Show(usr *models.User, flag bool) {
	if flag && usr != nil {
		BetterPwd(usr)
		fmt.Printf("%v@%v:%v ->", usr.Username, usr.Hostname, usr.Pwd)
	} else {
		fmt.Print("->")
	}
}

func BetterPwd(usr *models.User) { // replaces homedir with ~
	if strings.HasPrefix(usr.Pwd, usr.Homedir) {
		usr.Pwd = strings.Replace(usr.Pwd, usr.Homedir, "~", 1)
	}
}
