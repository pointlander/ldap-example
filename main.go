// Copyright 2019 The LDAP Example Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/tls"
	"fmt"

	"github.com/go-ldap/ldap"
)

func main() {
	username := "john"
	//password := "johnldap"

	bindusername := "uid=john,ou=People,dc=nodomain"
	bindpassword := "johnldap"

	config := tls.Config{
		InsecureSkipVerify: true,
	}
	connection, err := ldap.DialURL("ldap://:1234")
	if err != nil {
		panic(err)
	}

	err = connection.StartTLS(&config)
	if err != nil {
		panic(err)
	}

	err = connection.Bind(bindusername, bindpassword)
	if err != nil {
		panic(err)
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
		panic(err)
	}

	if len(sr.Entries) != 1 {
		panic("User does not exist or too many entries returned")
	}
}
