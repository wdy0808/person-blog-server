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

