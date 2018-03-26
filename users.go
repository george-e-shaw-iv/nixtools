// nixtools is a toolkit for linux system operations
package nixtools

import (
	"os/exec"
	"bytes"
	"errors"
	"strconv"
)

// Function GetUserID returns the ID associated with
// a username if it exists or -1 and an error if the
// given username does not exist.
func GetUserID(username string) (int, error) {
	cmd := exec.Command("id", "-u", username)

	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return -1, errors.New(stderr.String())
	}

	id, err := strconv.Atoi(stdout.String())
	if err != nil {
		return -1, err
	}

	return id, nil
}

// Function UserExists returns true if the given
// username exists within the system or false if
// the given username does not exist.
func UserExists(username string) (bool) {
	cmd := exec.Command("id", "-u", username)

	if err := cmd.Run(); err != nil {
		return false
	}

	return true
}

// Create user will attempt to create a new user
// in the system with the given username.
func CreateUser(username string) {

}

// Create user will attempt to delete a given user
// in the system with the given username if the
// given username exists.
func DeleteUser(username string) {

}