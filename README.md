# ReSTful Application (Respresentational State Transfer)

```
CRUD => {Create, Read, Update, Destroy}

HTTP Methods => {GET, POST, PUT, DELETE, OPTIONS, PATCH, ...}
```

## Employee Management Server (JSON API)

```
CRUD      |       Actions         |   HTTP Method     |          URI        |         Request Body            |     Response Body
----------------------------------------------------------------------------------------------------------------------------------------------
Read      |       Index           |      GET          |   /employees        |               -                 |       [{...}, ...]
Read      |       Show            |      GET          |   /employees/{id}   |               -                 |       {...}
Create    |       Create          |      POST         |   /employees        |             {...}               |       {id: , ...}
Update    |       Update          |      PUT          |   /employees/{id}   |             {...}               |       {...}
Update    |       Update          |      PATCH        |   /employees/{id}   |            {some attrs}         |       {...} / -
Destroy   |       Destroy         |      DELETE       |   /employees/{id}   |               -                 |       {...} / -
```

---

## Architecture

MVC => {Model, View, Controller}

Clean-code Architecture => {Entity, Repository, Transport [HTTP, gRPC, ...], Service}
