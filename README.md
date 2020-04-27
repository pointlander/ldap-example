# LDAP Example

## About
Demo of ldap auth with ldap container.

## Demo
Start the LDAP server:
```bash
docker run -d -p 1234:389 pointlander/ldap
```

Build the demo:
```bash
go build
```

Run the demo:
```bash
./ldap-example
```
