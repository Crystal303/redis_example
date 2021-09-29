module redis_example

go 1.17

require (
	github.com/gomodule/redigo v1.8.5
	github.com/stretchr/testify v1.7.0
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace github.com/gomodule/redigo v1.8.5 => github.com/gomodule/redigo v2.0.0+incompatible
