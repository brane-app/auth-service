module github.com/imonke/auth-service

go 1.13

require (
	github.com/gastrodon/groudon v0.0.0-20200703235457-2189c8855d11
	github.com/google/uuid v1.1.1
	github.com/imonke/monkebase v0.0.0-20200702171637-920be2f7cf25
	github.com/imonke/monketype v0.0.0-20200702003911-d46b79f825eb
)

replace github.com/imonke/monkebase => ../../monkebase

replace github.com/imonke/monketype => ../../monketype
