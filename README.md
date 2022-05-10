```
.
├── LICENSE
├── README.md
├── api
│   ├── common
│   │   └── dresponse.go
│   ├── insomnia.json
│   ├── router.go
│   └── v1
│       └── content
│           ├── controller.go
│           ├── request
│           │   ├── create_content.go
│           │   └── update_content.go
│           └── response
│               ├── create_new_content.go
│               ├── get_content_by_id.go
│               └── get_content_by_tag.go
├── app
│   ├── main.go
│   └── modules
│       └── modules.go
├── business
│   ├── content
│   │   ├── item.go
│   │   ├── service.go
│   │   ├── service_test.go
│   │   └── spec
│   │       └── upsert_item.go
│   └── error.go
├── config
│   ├── config.go
│   └── config.yaml
├── go.mod
├── go.sum
├── modules
│   └── repository
│       └── content
│           ├── couchdb_repo.go
│           ├── factory.go
│           ├── mongo_repo.go
│           └── mysql_repo.go
├── poseidon.png
├── run.sh
└── util
    ├── dbdriver.go
    └── idgen.go
```
