# nixtools
A golang toolkit for linux system operations.

The GoDoc reference for `george-e-shaw-iv/nixtools` is located [here](https://godoc.org/github.com/george-e-shaw-iv/nixtools).

## Table of Contents

1. [User Functions](#user-functions)
2. [SSH Functions](#ssh-functions)


## User Functions
```go
// Function GetUserID returns the ID associated with
// a username if it exists or -1 and an error if the
// given username does not exist.
func GetUserID(username string) (int, error) {}
```

```go
// Function UserExists returns true if the given
// username exists within the system or false if
// the given username does not exist.
func UserExists(username string) (bool) {}
```

```go
// Create user will attempt to create a new user
// in the system with the given username. If password
// parameter is empty then the --disabled-password
// flag will be set and the user will not be able to
// authenticate via password
func CreateUser(username, password string) (error) {}
```

```go
// Create user will attempt to delete a given user
// in the system with the given username if the
// given username exists. deleteOwnedFiles parameter
// will remove all files owned by the username and
// removeHome parameter will delete the home directory
// of the user.
func DeleteUser(username string, deleteOwnedFiles, removeHome bool) (error) {}
```

```go
// Create user will attempt to delete a given user
// in the system with the given username if the
// given username exists. deleteOwnedFiles parameter
// will remove all files owned by the username and
// removeHome parameter will delete the home directory
// of the user.
func DeleteUser(username string, deleteOwnedFiles, removeHome bool) (error) {}
```

## SSH Functions
```go
// Function InitSSH creates the necessary folders,
// files, and generates a default key-pair for the
// given username. If parameter rootHasAccess is set
// to true then the public key of the root (sudo) user
// will be copied into the authorized_keys file of
// the user.
func InitSSH(username string, rootHasAccess bool) (error) {}
```

```go
// Function AddAuthorizedKey adds a new public key to
// a given username's authorized_keys file.
func AddAuthorizedKey(username, key string) error {}
```

```go
// Function DeleteAuthorizedKey removes a public key
// that is already in the authorized_keys file of
// said username.
func DeleteAuthorizedKey(username, key string) error {}
```

```go
// Function GetAuthorizedKeys will return a slice
// of strings that contains all of the public keys
// within a given username's authorized_keys file.
// If the parameter removeRootKey is set to true the
// public key of the current root user of the system,
// if found within the file, will not be placed within
// the slice of strings.
func GetAuthorizedKeys(username string, removeRootKey bool) ([]string, error) {}
```