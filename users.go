// Package nixtools is a toolkit for various linux
// system operations.
package nixtools

import (
	"os/exec"
	"bytes"
	"errors"
	"strconv"
	"io"
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
// in the system with the given username. If password
// parameter is empty then the --disabled-password
// flag will be set and the user will not be able to
// authenticate via password
func CreateUser(username, password string) (error) {
	var stderr bytes.Buffer
	var args []string

	if len(password) > 0 {
		args = []string{
			"--gecos",
			"",
			username,
		}
	} else {
		args = []string{
			"--disabled-password",
			"--gecos",
			"",
			username,
		}
	}

	cmd := exec.Command("adduser", args...)
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.New(stderr.String())
	}

	if len(password) > 0 {
		stderr.Reset()
		reader, writer := io.Pipe()

		cmdOut := exec.Command("echo", username+":"+password)
		cmdOut.Stdout = writer

		cmdIn := exec.Command("chpasswd")
		cmdIn.Stdin = reader
		cmdIn.Stderr = &stderr

		cmdOut.Start()
		cmdIn.Start()

		cmdOut.Wait()
		writer.Close()

		cmdIn.Wait()
		reader.Close()

		if len(stderr.Bytes()) > 0 {
			return errors.New(stderr.String())
		}
	}

	return nil
}

// Create user will attempt to delete a given user
// in the system with the given username if the
// given username exists. deleteOwnedFiles parameter
// will remove all files owned by the username and
// removeHome parameter will delete the home directory
// of the user.
func DeleteUser(username string, deleteOwnedFiles, removeHome bool) (error) {
	var stderr bytes.Buffer
	var args []string

	if deleteOwnedFiles {
		args = []string{
			"--quiet",
			"--remove-all-files",
			username,
		}
	} else if removeHome {
		args = []string{
			"--quiet",
			"--remove-home",
			username,
		}
	} else {
		args = []string{
			"--quiet",
			username,
		}
	}

	// Delete the user and try to remove all files associated
	cmd := exec.Command("deluser", args...)
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return errors.New(stderr.String())
	}

	// Forcefully remove the users home directory as sometimes
	// the --remove-home or --remove-all-files flags don't work
	if deleteOwnedFiles || removeHome {
		cmd = exec.Command("rm", "-rf", "/home/"+username)
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			return errors.New(stderr.String())
		}
	}

	return nil
}