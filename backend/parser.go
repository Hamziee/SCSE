package backend

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func ReadINI(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make(map[string]string)
	currentSection := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = line[1 : len(line)-1]
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			fullKey := key
			if currentSection != "" {
				fullKey = currentSection + "|" + key
			}
			data[fullKey] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func WriteINI(path string, data map[string]string) error {
	sections := make(map[string]map[string]string)
	globals := make(map[string]string)

	for fullKey, value := range data {
		parts := strings.SplitN(fullKey, "|", 2)
		if len(parts) == 2 {
			sec := parts[0]
			key := parts[1]
			if sections[sec] == nil {
				sections[sec] = make(map[string]string)
			}
			sections[sec][key] = value
		} else {
			globals[fullKey] = value
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	var globalKeys []string
	for k := range globals {
		globalKeys = append(globalKeys, k)
	}
	sort.Strings(globalKeys)

	for _, k := range globalKeys {
		if _, err := fmt.Fprintf(writer, "%s=%s\n", k, globals[k]); err != nil {
			return err
		}
	}

	var sectionNames []string
	for k := range sections {
		sectionNames = append(sectionNames, k)
	}
	sort.Strings(sectionNames)

	for _, sec := range sectionNames {
		var keys []string
		for k := range sections[sec] {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		if _, err := fmt.Fprintf(writer, "\n[%s]\n", sec); err != nil {
			return err
		}
		for _, k := range keys {
			if _, err := fmt.Fprintf(writer, "%s=%s\n", k, sections[sec][k]); err != nil {
				return err
			}
		}
	}

	return writer.Flush()
}
