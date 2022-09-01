package cat

import (
	"catinator-backend/pkg/db/ent"
	"strings"
)

func stringSortToEnt(s string, fields ...string) ent.OrderFunc {
	if strings.ToLower(s) == "desc" {
		return ent.Desc(fields...)
	}
	return ent.Asc(fields...)
}
