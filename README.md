```
.
│   config.toml
│   go.mod
│   go.sum
│   main.go
│
├───api
│   │   router.go
│   │
│   ├───common
│   │   ├───obj
│   │   │       obj_common.go
│   │   │
│   │   └───response
│   │           response_ommon.go
│   │
│   ├───middleware
│   │       middleware.go
│   │
│   └───v1
│       ├───auth
│       │       controller.go
│       │
│       ├───movie
│       │       controller.go
│       │
│       └───user
│               controller.go
│
├───app
│   └───modules
│           modules.go
│
├───business
│   ├───movie
│   │   │   service.go
│   │   │   service_test.go
│   │   │
│   │   ├───dto
│   │   │       movie_dto.go
│   │   │
│   │   ├───entity
│   │   │       movie_entity.go
│   │   │
│   │   └───response
│   │           movie_response.go
│   │
│   └───user
│       │   auth_service.go
│       │   jwt_service.go
│       │   service.go
│       │   service_test.go
│       │
│       ├───dto
│       │       login.dto.go
│       │       register.dto.go
│       │       user.dto.go
│       │
│       ├───entity
│       │       user_entity.go
│       │
│       └───response
│               user_response.go
│
├───config
│       config.go
│
├───repository
│   ├───movie
│   │       factory.go
│   │       mysql.go
│   │       postgres.go
│   │
│   └───user
│           factory.go
│           mongo_repo.go
│           mysql.go
│           postgres.go
│
└───util
        driver.go
```
