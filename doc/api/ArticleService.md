# Responsibility
1. Provide api for client to CURD article detail.
2. Provide api for client to get CRUD group detail.

# Definition
### Article
```
{
    "entity_id": string,
    "group_id": string,
    "author_id": string,
    "title": string,
    "abstract": string,
    "content": string,
    "create_time": uint64,
    "last_update_time": uint64,
}
```
##### Url
/article
##### Method
get
patch
delete
post
##### Query
search-by-group
search-by-user
get-all
### Group
```
{
    "entity_id": string,
    "name": string,
}
```
##### Url
/article-group
##### Method
get
patch
delete
post
##### Query
get-all


