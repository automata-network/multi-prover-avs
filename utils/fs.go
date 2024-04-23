package utils

import (
	"crypto/ecdsa"
	"os"
	"path/filepath"
	"strings"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdkEcdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	"github.com/chzyer/logex"
	"github.com/chzyer/readline"
)

func FixFilepath(path string) string {
	if strings.HasPrefix(path, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			logex.Fatal(err)
		}
		path = filepath.Join(homeDir, path[2:])
	}
	return path
}

func ReadBlsKey(path string, password string) (*bls.KeyPair, error) {
	expandPath := FixFilepath(path)
	key, err := bls.ReadPrivateKeyFromFile(expandPath, password)
	if err != nil {
		return nil, logex.NewErrorf("read bls key from [%v] fail: %v", path, err)
	}
	return key, nil
}

func PromptEcdsaKey(path string) (*ecdsa.PrivateKey, error) {
	expandKey := FixFilepath(path)
	if _, err := os.Stat(expandKey); err != nil {
		if os.IsNotExist(err) {
			return nil, logex.NewErrorf("ecdsa key not found: %v", path)
		}
	}
	password, err := ReadPassword("Enter the password for " + path)
	if err != nil {
		return nil, logex.Trace(err)
	}
	key, err := sdkEcdsa.ReadKey(expandKey, password)
	if err != nil {
		return nil, logex.NewErrorf("read ecdsa key from [%v] fail: %v", path, err)
	}
	return key, nil
}

func ReadPassword(prompt string) (string, error) {
	rl, err := readline.New(prompt)
	if err != nil {
		return "", logex.Trace(err)
	}
	defer rl.Close()

	setPasswordCfg := rl.GenPasswordConfig()
	setPasswordCfg.SetListener(func(line []rune, pos int, key rune) (newLine []rune, newPos int, ok bool) {
		rl.SetPrompt(prompt + ": " + strings.Repeat("*", len(line)))
		rl.Refresh()
		return nil, 0, false
	})
	pswd, err := rl.ReadPasswordWithConfig(setPasswordCfg)
	if err != nil {
		return "", logex.Trace(err)
	}
	return string(pswd), nil
}
