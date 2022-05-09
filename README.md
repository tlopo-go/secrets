# secrets
Command line based password manager


# Why?
It's a bad practice to leave credentials in plain-text files in our laptops, and for local scripts Hashicorp vault is over kill.

# Usage: 

```
Command line secret manager

Usage:
  secrets [flags]
  secrets [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Retrieves a secret
  help        Help about any command
  init        Initialize the secrets database
  lock        Locks the secrets database
  set         Sets a secret
  unlock      Unlocks the secrets database

Flags:
  -h, --help   help for secrets

Use "secrets [command] --help" for more information about a command.
```

Initialize Database: 
```
$ secrets init
Creating a new secrets database

Please enter password for the new database: 
Please confirm password: 
2022/05/09 13:01:02 Database /Users/tiago/.secrets/db.kdbx created
```
Create Secret:
```
$ secrets set --service foo --account  user@noemail.com --password pass123
```
Read secret:
```
$ secrets get --service foo 
{"service":"foo","account":"user@noemail.com","password":"pass123"}
```



