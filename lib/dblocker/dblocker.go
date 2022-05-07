package dblocker

import (
	"github.com/denisbrodbeck/machineid"
	"github.com/tlopo-go/secrets/lib/crypt"
	"log"
	"os"
)

func Lock() {
	// just remove the file
}

func Unlock(pass, filename string) {
	c := getCrypt()

	enc, err := c.Encrypt([]byte(pass))
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Before writing the file, check if the password is correct
	err = os.WriteFile(filename, enc, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPassword(filename string) (pass string) {
	enc, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	c := getCrypt()

	dec, err := c.Decrypt(enc)
	if err != nil {
		log.Fatal(err)
	}
	pass = string(dec)
	return
}

func getCrypt() (c crypt.Crypt) {
	id, err := machineid.ProtectedID("secrets")
	if err != nil {
		log.Fatal(err)
	}

	key := []byte(id)[:32]
	iv := []byte(id)[:16]
	c = crypt.NewCrypt(key, iv)
	return
}
