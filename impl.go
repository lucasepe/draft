package draft

import (
	"encoding/csv"
	"io"
	"strings"

	"github.com/rakyll/statik/fs"
)

func setImpl(com *Component) {
	if s := strings.TrimSpace(com.Impl); len(s) > 0 {
		return
	}

	impl := getCloudImpl(com.Provider, com.Kind)
	if len(impl) > 0 {
		com.Impl = impl
	}
}

func getCloudImpl(provider, kind string) string {
	switch strings.TrimSpace(strings.ToLower(provider)) {
	case "aws":
		return awsImpl()(kind)
	case "google":
		return googleImpl()(kind)
	case "azure":
		return azureImpl()(kind)
	default:
		return defaultImpl()(kind)
	}
}

func awsImpl() func(string) string {
	dict, _ := readCsvFile("/aws.csv")

	return func(key string) string {
		return dict[key]
	}
}

func googleImpl() func(string) string {
	dict, _ := readCsvFile("/google.csv")

	return func(key string) string {
		return dict[key]
	}
}

func azureImpl() func(string) string {
	dict, _ := readCsvFile("/azure.csv")

	return func(key string) string {
		return dict[key]
	}
}

func defaultImpl() func(string) string {
	dict, _ := readCsvFile("/default.csv")

	return func(key string) string {
		return dict[key]
	}
}

func readCsvFile(filePath string) (map[string]string, error) {
	dict := map[string]string{}

	sfs, err := fs.New()
	if err != nil {
		return dict, err
	}

	f, err := sfs.Open(filePath)
	if err != nil {
		return dict, err
	}
	defer f.Close()

	csvr := csv.NewReader(f)

	for {
		row, err := csvr.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return dict, err
		}

		dict[row[0]] = row[1]
	}
}
