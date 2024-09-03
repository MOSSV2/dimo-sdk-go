package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/howeyc/gopass"
)

const (
	RepoStr      = "repo"
	PasswordStr  = "password"
	EndpointStr  = "bind"
	ExposeStr    = "expose"
	RemoteURLStr = "remote"
)

func InputPassWord() (string, error) {
	var password string
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	go func() {
		defer cancel()
		fmt.Printf("Please input your password (at least 8): ")
		pd, err := gopass.GetPasswdMasked()
		if err != nil {
			return
		}
		password = string(pd)
	}()

	<-ctx.Done()

	if len(password) < 8 {
		return password, fmt.Errorf("password length should be at least 8")
	}
	return password, nil
}
