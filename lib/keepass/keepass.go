package keepass

import (
	"encoding/json"
	"errors"
	"github.com/tobischo/gokeepasslib/v3"
	w "github.com/tobischo/gokeepasslib/v3/wrappers"
	"log"
	"os"
)

type Secret struct {
	Service  string `json:"service"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (s *Secret) ToJson() string {
	b, _ := json.Marshal(s)
	return string(b)
}

type KeePass struct {
	DatabaseFile   string
	MasterPassword string
}

func mkValue(key string, value string) gokeepasslib.ValueData {
	return gokeepasslib.ValueData{Key: key, Value: gokeepasslib.V{Content: value}}
}

func mkProtectedValue(key string, value string) gokeepasslib.ValueData {
	return gokeepasslib.ValueData{
		Key:   key,
		Value: gokeepasslib.V{Content: value, Protected: w.NewBoolWrapper(true)},
	}
}

func (k *KeePass) CreateDatabase() error {
	if _, err := os.Stat(k.DatabaseFile); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(k.DatabaseFile)
		if err != nil {
			return err
		}
		defer file.Close()

		rootGroup := gokeepasslib.NewGroup()
		rootGroup.Name = "Root"

		db := &gokeepasslib.Database{
			Header:      gokeepasslib.NewHeader(),
			Credentials: gokeepasslib.NewPasswordCredentials(k.MasterPassword),
			Content: &gokeepasslib.DBContent{
				Meta: gokeepasslib.NewMetaData(),
				Root: &gokeepasslib.RootData{
					Groups: []gokeepasslib.Group{rootGroup},
				},
			},
		}

		db.LockProtectedEntries()
		keepassEncoder := gokeepasslib.NewEncoder(file)
		if err := keepassEncoder.Encode(db); err != nil {
			panic(err)
		}

		log.Printf("Database %s created\n", k.DatabaseFile)
		return nil
	}
	log.Printf("Database %s already exists", k.DatabaseFile)
	return nil
}

func (k *KeePass) CheckPassword() (err error) {
	r, err := os.Open(k.DatabaseFile)
	if err != nil {
		return err
	}
	defer r.Close()

	//Recover from panic
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Wrong Password")
		}
	}()

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials(k.MasterPassword)
	gokeepasslib.NewDecoder(r).Decode(db)
	db.UnlockProtectedEntries()
	return
}

type withDB func(db *gokeepasslib.Database)

func (k *KeePass) withDBRead(fn withDB) (err error) {
	r, err := os.Open(k.DatabaseFile)
	if err != nil {
		return
	}
	defer r.Close()

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials(k.MasterPassword)
	gokeepasslib.NewDecoder(r).Decode(db)
	db.UnlockProtectedEntries()
	fn(db)
	return
}

func (k *KeePass) withDBWrite(fn withDB) error {
	r, err := os.Open(k.DatabaseFile)
	if err != nil {
		return err
	}
	defer r.Close()

	db := gokeepasslib.NewDatabase()
	db.Credentials = gokeepasslib.NewPasswordCredentials(k.MasterPassword)
	gokeepasslib.NewDecoder(r).Decode(db)
	db.UnlockProtectedEntries()

	fn(db)

	db.LockProtectedEntries()

	w, err := os.Create(k.DatabaseFile)
	if err != nil {
		return err
	}
	defer w.Close()

	encoder := gokeepasslib.NewEncoder(w)
	if err := encoder.Encode(db); err != nil {
		return err
	}
	return nil
}

func (k *KeePass) Read(service string) (secret Secret, err error) {
	k.withDBRead(func(db *gokeepasslib.Database) {
		entries := db.Content.Root.Groups[0].Entries
		var index = -1
		for i, e := range entries {
			if e.GetTitle() == service {
				index = i
			}
		}

		if index == -1 {
			err = errors.New("Secret not found")
		} else {
			e := entries[index]
			secret = Secret{e.GetTitle(), e.GetContent("UserName"), e.GetPassword()}
		}
	})
	return
}

func (k *KeePass) List() (names []string, err error) {
	k.withDBRead(func(db *gokeepasslib.Database) {
		entries := db.Content.Root.Groups[0].Entries
		for _, e := range entries {
			names = append(names, e.GetTitle())
		}
	})
	return
}

func (k *KeePass) Write(s Secret) (err error) {
	entry := gokeepasslib.NewEntry()
	entry.Values = append(entry.Values, mkValue("Title", s.Service))
	entry.Values = append(entry.Values, mkValue("UserName", s.Account))
	entry.Values = append(entry.Values, mkProtectedValue("Password", s.Password))

	err = k.withDBWrite(func(db *gokeepasslib.Database) {

		// Find secret by title
		var index = -1
		for i, e := range db.Content.Root.Groups[0].Entries {
			if e.GetTitle() == entry.GetTitle() {
				index = i
			}
		}

		if index == -1 {
			db.Content.Root.Groups[0].Entries = append(db.Content.Root.Groups[0].Entries, entry)
		} else {
			db.Content.Root.Groups[0].Entries[index] = entry
		}
	})
	return
}

func (k *KeePass) Delete(service string) (err error) {
	err = k.withDBWrite(func(db *gokeepasslib.Database) {

		// Find secret by title
		var index = -1
		for i, e := range db.Content.Root.Groups[0].Entries {
			if e.GetTitle() == service {
				index = i
			}
		}

		if index != -1 {
			entries := db.Content.Root.Groups[0].Entries
			db.Content.Root.Groups[0].Entries = append(entries[:index], entries[index+1:]...)
		}
	})
	return
}
