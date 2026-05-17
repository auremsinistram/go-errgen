package tools

import (
	"bytes"
	"encoding/json"
	"hash/fnv"
	"os"
	"path/filepath"
	"strings"

	"github.com/auremsinistram/go-errors"
)

func CheckDir(path string) error {
	dir := filepath.Dir(path)

	if dir == "" || dir == "." {
		return nil
	}

	info, err := os.Stat(dir)
	if err == nil {
		if info.IsDir() {
			return nil
		}

		return errors.Errorf("path %q exists but is not a directory", dir)
	}

	if !os.IsNotExist(err) {
		return errors.Wrap(err, "tools - CheckDir - #1")
	}

	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.Wrap(err, "tools - CheckDir - #2")
	}

	return nil
}

func PascalCase(input string) string {
	replacer := strings.NewReplacer("-", " ", "_", " ")
	normal := replacer.Replace(input)
	words := strings.Fields(normal)

	for i := range words {
		if len(words[i]) > 0 {
			words[i] = strings.ToUpper(words[i][:1]) + strings.ToLower(words[i][1:])
		}
	}

	return strings.Join(words, "")
}

func Code(input string) int {
	hash := fnv.New32a()

	hash.Write([]byte(input))

	return int(hash.Sum32())
}

func CompactJSON(data []byte) ([]byte, error) {
	var buffer bytes.Buffer

	err := json.Compact(&buffer, data)
	if err != nil {
		return nil, errors.Wrap(err, "tools - CompactJSON - #1")
	}

	return buffer.Bytes(), nil
}
