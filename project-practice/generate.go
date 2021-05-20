package generate

//go:generate kratos proto client .
//go:generate go run entgo.io/ent/cmd/ent generate ./internal/weekly-task/data/ent/schema
