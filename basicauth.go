package basicauth

import (
    "encoding/base64"
    "github.com/pkg/errors"
)

const VERSION = "1.0.0"

// BasicAuth takes a username / password as a string
// and returns a base64 encoded string
func BasicAuth(username, password string) (string, error) {
    if err := validate(username, password); err != nil {
        return "", errors.Wrap(err, "cannot build basic auth string")
    }
    return basicAuth(username, password), nil
}

// BasicAuthUnsafe takes a username / password as a string
// and returns a base64 encoded string, however, string length
// will not be validated
func BasicAuthUnsafe(username, password string) string {
    return basicAuth(username, password)
}

func basicAuth(u, p string) string {
	a := []byte(u + ":" + p)
	return base64.StdEncoding.EncodeToString(a)
}

func validate(u, p string) error {
    var uerr error
    var perr error

    if len(u) == 0 {
        uerr = errors.New("username is length zero")
    }

    if len(p) == 0 {
        perr = errors.New("password is length zero")
    }

    if uerr != nil && perr != nil {
        return errors.Wrap(uerr, perr.Error())
    }

    if uerr != nil {
        return uerr
    }

    if perr != nil {
        return perr
    }

    return nil
}
