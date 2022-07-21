package custom

import (
	"path/filepath"
	"strings"
)

func ReadFileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
