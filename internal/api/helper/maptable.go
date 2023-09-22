package helper

import (
	"strings"
)

func MapTableName(input string) (string, bool) {
	tableDict := map[string]string{
		"invitation": "invitations",
		"bride":      "brides",
		"groom":      "grooms",
	}
	tableName := tableDict[strings.ToLower(input)]

	if tableName == "" {
		return "", false
	}
	return tableName, true
}
