package basicauth

import (
    "testing"
    "strings"
    "encoding/base64"
)

func TestBasicAuth(t *testing.T) {
    b := base64.StdEncoding.EncodeToString([]byte("username:password"))
    result, err := BasicAuth("username", "password")
    if err != nil {
        t.Errorf("expected BasicAuth(\"username\", \"password\") to return a nil error, got " + err.Error())
        return
    }

    if b != result {
        t.Errorf("expected BasicAuth(\"username\", \"password\") to return " + b + " got " + result)
    }
}

func TestBasicAuthEmptyUser(t *testing.T) {
    if _, err := BasicAuth("", "password"); err == nil {
        t.Errorf("expected an empty username to return an error, got nil")
    }
}

func TestBasicAuthEmptyPass(t *testing.T) {
    if _, err := BasicAuth("username", ""); err == nil {
        t.Errorf("expected an empty password to return an error, got nil")
    }
}

func TestBasicAuthEmptyUserPass(t *testing.T) {
    if _, err := BasicAuth("", ""); err == nil {
        t.Errorf("expected an empty username and password to return an error, got nil")
    }

    _, err := BasicAuth("", "")
    if !strings.Contains(err.Error(), "username") && !strings.Contains(err.Error(), "password")  {
        t.Errorf("expected an empty username and password to return an error, got nil")
    }
}

func TestBasicAuthUnsafe(t *testing.T) {
    b := base64.StdEncoding.EncodeToString([]byte("username:password"))
    result := BasicAuthUnsafe("username", "password")

    if b != result {
        t.Errorf("expected BasicAuthUnsafe(\"username\", \"password\") to return " + b + " got " + result)
    }
}
