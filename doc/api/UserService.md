# Responsibility
1. Provide api for client to login and sign up.
2. Provide api to set admin.

#Definition
### User
```
{
    "entity_id": string,
    "username": string,
    "icon_url": string,
    "password": string,
    "is_admin": bool,
    "create_time": uint64,
}
```
##### Url
/user
##### Method
patch
post
### Login
```
{
    "username": string,
    "password": string,
}
```
##### Url
/login
##### Method
post
### Signup
```
{
    "username": string,
    "password": string,
}
```
##### Url
/login
##### Method
post

# Model
## authorization
JSON Web Tokens
## Error Code

| 1001 | username is not available |
| --- | --- |
| 1003 | Username has not been signed up |
| 1004 | password wrong |


