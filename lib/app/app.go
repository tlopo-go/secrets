package app

import (
	"errors"
	"fmt"
	"github.com/tlopo-go/secrets/lib/dblocker"
	k "github.com/tlopo-go/secrets/lib/keepass"
	"golang.org/x/term"
	"log"
	"os"
	"syscall"
)

var cfg config

type config struct {
	databaseHome   string
	databasePath   string
	masterPassword string
}

var unlockedFile string

func init() {
	cfg.databaseHome = os.Getenv("SECRETS_HOME")
	if len(cfg.databaseHome) == 0 {
		cfg.databaseHome = fmt.Sprintf("%s/.secrets", os.Getenv("HOME"))
	}

	err := os.MkdirAll(cfg.databaseHome, 0755)
	if err != nil {
		log.Fatal(err)
	}

	cfg.databasePath = fmt.Sprintf("%s/db.kdbx", cfg.databaseHome)

	unlockedFile = fmt.Sprintf("%s/.unlocked", cfg.databaseHome)

	if !IsDBLocked() {
		cfg.masterPassword = dblocker.GetPassword(unlockedFile)
	}
}

func GetDatabasePath() string {
	return cfg.databasePath
}

func GetMasterPassword() string {
	return cfg.masterPassword
}

func GetDatabaseHome() string {
	return cfg.databaseHome
}

func ValidateUnlocked() {
	if !IsDBInitialized() {
		log.Fatal("Secrets database not initialized, please run init command")
	} else if IsDBLocked() {
		log.Fatal("Secrets database is locked, please run unlock command")
	}
}

func IsDBLocked() bool {
	if _, err := os.Stat(unlockedFile); errors.Is(err, os.ErrNotExist) {
		return true
	}
	return false
}

func IsDBInitialized() bool {
	if _, err := os.Stat(cfg.databasePath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func UnlockDB() {
	msg := "To unlock database, please enter password: "
	cfg.masterPassword = promptPasswordWithMessage(msg)
	kp := k.KeePass{cfg.databasePath, cfg.masterPassword}
	err := kp.CheckPassword()
	if err != nil {
		log.Fatal(err)
	}
	dblocker.Unlock(cfg.masterPassword, unlockedFile)
}

func LockDB() {
	if !IsDBLocked() {
		dblocker.Lock(unlockedFile)
	} else {
		log.Println("Database is already locked")
	}
}

func InitDB() {
	if !IsDBInitialized() {
		fmt.Println("Creating a new secrets database\n")
		cfg.masterPassword = promptNewPassword()
		kp := k.KeePass{cfg.databasePath, cfg.masterPassword}
		kp.CreateDatabase()
		dblocker.Unlock(cfg.masterPassword, unlockedFile)
	} else {
		log.Println("Database already initialized")
	}
}

func promptNewPassword() (password string) {
	password = promptPasswordWithMessage("Please enter password for the new database: ")
	p2 := promptPasswordWithMessage("Please confirm password: ")
	if password != p2 {
		log.Fatal("Passwords entered do not match")
	}
	return
}

func promptPasswordWithMessage(message string) (password string) {
	fmt.Print(message)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		log.Fatal(err)
	}
	password = string(bytePassword)
	return
}
