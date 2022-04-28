package main

import(
    "fmt"
    "github.com/tlopo-go/secrets/lib/keepass"
)

func main(){
    fmt.Println("Started")
    k := keepass.KeePass { "/tmp/db.kdbx", "1234" }
    k.CreateDatabase()
    k.Write(keepass.Secret{"foo","bar","1234"})
    k.Write(keepass.Secret{"foo","bar","12345"})
    s, _ := k.Read("foo")
    fmt.Printf("%#v\n",s)
}
