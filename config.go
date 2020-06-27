package draft

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Opt is a function to define a configuration property.
type Opt func(cfg *Config)

// Verbose sets verbose output.
func Verbose(enable bool) func(*Config) {
	return func(cfg *Config) {
		cfg.verbose = enable
	}
}

// BottomTop sets the graph layout.
func BottomTop(enable bool) func(*Config) {
	return func(cfg *Config) {
		cfg.bottomTop = enable
	}
}

// Provider sets the cloud implementation (one of 'aws', 'gcp', 'azure')
func Provider(s string) func(*Config) {
	provs := map[string]bool{"aws": true, "google": true, "azure": true}

	return func(cfg *Config) {
		val := strings.ToLower(strings.TrimSpace(s))
		if provs[val] {
			cfg.provider = val
		} else {
			cfg.provider = ""
		}
	}
}

// URI sets the input YAML definition file.
// Can be also an HTTP URL.
func URI(s string) func(*Config) {
	return func(cfg *Config) {
		if !strings.HasPrefix(s, "http") {
			cfg.uri, _ = filepath.Abs(s)
		} else {
			cfg.uri = s
		}
	}
}

// IconsPath sets the custom icons path.
func IconsPath(s string) func(*Config) {
	return func(cfg *Config) {
		cfg.iconsPath = strings.TrimSpace(s)
	}
}

// Config defines the 'draft' configuration.
type Config struct {
	bottomTop bool
	provider  string
	verbose   bool
	iconsPath string
	uri       string
}

// NewConfig create a configuration
// with the specified attributes.
func NewConfig(opts ...Opt) Config {
	res := Config{}

	for _, op := range opts {
		op(&res)
	}

	return res
}

// Load a YAML from the config info.
func Load(cfg Config) (Design, error) {
	const bytesLimit = 500 * 1024

	if strings.HasPrefix(cfg.uri, "http") {
		body, err := getURI(cfg.uri, bytesLimit)
		if err != nil {
			return Design{}, err
		}

		return decodeYAML(body)
	}

	body, err := getFILE(cfg.uri, bytesLimit)
	if err != nil {
		return Design{}, err
	}

	return decodeYAML(body)
}

// getURI fetch data (with limit) from an HTTP URL
func getURI(uri string, limit int64) ([]byte, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(io.LimitReader(response.Body, limit))
}

// getFILE fetch data (with limit) from an file
func getFILE(fin string, limit int64) ([]byte, error) {
	file, err := os.Open(fin)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return ioutil.ReadAll(io.LimitReader(file, limit))
}

// decodeYAML decode a YAML to return a Design struct.
func decodeYAML(dat []byte) (Design, error) {
	res := Design{}

	// Init new YAML decode
	d := yaml.NewDecoder(bytes.NewReader(dat))

	// Start YAML decoding from file
	if err := d.Decode(&res); err != nil {
		return res, err
	}

	return res, nil
}
