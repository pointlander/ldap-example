// Copyright 2020 The LDAP Example Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/tls"
	"errors"
	"fmt"

	"github.com/go-ldap/ldap"
)

func main() {
	err := Auth("john", "johnldap", "uid=john,ou=People,dc=nodomain", "johnldap")
	if err != nil {
		panic(err)
	}
	err = Auth("jane", "janeldap", "uid=jane,ou=People,dc=nodomain", "janeldap")
	if err != nil {
		panic(err)
	}
	err = Auth("john", "password", "uid=john,ou=People,dc=nodomain", "password")
	if err == nil {
		panic("password should be invalid")
	}
	err = Auth("jane", "password", "uid=jane,ou=People,dc=nodomain", "password")
	if err == nil {
		panic("password should be invalid")
	}
}

// Auth authorizes an user
func Auth(username, password, bindusername, bindpassword string) error {
	config := tls.Config{
		InsecureSkipVerify: true,
	}
	connection, err := ldap.DialURL("ldap://:1234")
	if err != nil {
		return err
	}

	err = connection.StartTLS(&config)
	if err != nil {
		return err
	}

	err = connection.Bind(bindusername, bindpassword)
	if err != nil {
		return err
	}

	searchRequest := ldap.NewSearchRequest(
		"dc=nodomain",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", username),
		[]string{"dn"},
		nil,
	)

	sr, err := connection.Search(searchRequest)
	if err != nil {
		return err
	}

	if len(sr.Entries) != 1 {
		return errors.New("User does not exist or too many entries returned")
	}

	user := sr.Entries[0].DN
	err = connection.Bind(user, password)
	if err != nil {
		return err
	}

	return nil
}
