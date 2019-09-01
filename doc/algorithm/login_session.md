# Background
1. Create session mechanism to keep login state.
2. Implement admin authorization.

# Related Algorithm
JSON Web Tokens
> https://jwt.io/introduction/

# Details
## Structure Definition
### Header
```shell
{
    "alg": "HS256",
    "typ": "JWT"
}
```
### Payload
```shell
{
    "expire_name": "{nanoseconds}",
    "user_id": "",
    "is_admin": "{true}",
    "is_root": "{true only wdy}"    
}
```
### Secret key
"person-blog-server-auth"