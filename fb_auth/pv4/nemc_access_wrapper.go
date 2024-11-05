package fbauth

import (
	I18n "Eulogist/fb_auth/i18n"
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

type AccessWrapper struct {
	ServerCode     string
	ServerPassword string
	Token          string
	Client         *Client
	Username       string
	Password       string
}

func NewAccessWrapper(Client *Client, ServerCode, ServerPassword, Token, username, password string) *AccessWrapper {
	return &AccessWrapper{
		Client:         Client,
		ServerCode:     ServerCode,
		ServerPassword: ServerPassword,
		Token:          Token,
		Username:       username,
		Password:       password,
	}
}

func (aw *AccessWrapper) GetAccess(ctx context.Context, publicKey []byte) (authResponse AuthResponse, err error) {
	/*
		if aw.ServerCode != "68497817" && aw.ServerCode != "48285363" {
			return AuthResponse{}, fmt.Errorf("GetAccess: Server code not match, and it only could be 68497817 or 48285363")
		}
	*/
	// pre check
	pubKeyData := base64.StdEncoding.EncodeToString(publicKey)
	authResponse, err = aw.Client.Auth(ctx, aw.ServerCode, aw.ServerPassword, pubKeyData, aw.Token, aw.Username, aw.Password)
	if err != nil {
		return AuthResponse{}, err
	}
	if len(authResponse.FBToken) != 0 {
		homedir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(I18n.T(I18n.Warning_UserHomeDir))
			homedir = "."
		}
		fbconfigdir := filepath.Join(homedir, ".config", "fastbuilder")
		os.MkdirAll(fbconfigdir, 0755)
		ptoken := filepath.Join(fbconfigdir, "fbtoken")
		// 0600: -rw-------
		token_file, err := os.OpenFile(ptoken, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return AuthResponse{}, err
		}
		_, err = token_file.WriteString(authResponse.FBToken)
		if err != nil {
			return AuthResponse{}, err
		}
		token_file.Close()
	}
	return
}