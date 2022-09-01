package db

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent --idtype string --feature sql/upsert --feature sql/modifier ./schema
