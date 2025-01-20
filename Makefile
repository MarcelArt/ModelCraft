swag:
	@swag init --parseInternal --parseDependency

templ:
	@templ generate

run: swag
	@go run main.go

watch: swag
	@air

scaffolder:
	@go run main.go scaffold ${model}
	make swag

migrate-up:
	@go run main.go migrate up

migrate-down:
	@go run main.go migrate down