package draft

import "strings"

const (
	kindBlockStore        = "bst"
	kindCDN               = "cdn"
	kindClient            = "cli"
	kindUnknown           = "unk"
	kindDNS               = "dns"
	kindNoSQL             = "doc"
	kindFileStore         = "fst"
	kindFunction          = "fun"
	kindGateway           = "gtw"
	kindContainersManager = "kub"
	kindLBA               = "lba"
	kindCache             = "mem"
	kindPubSub            = "msg"
	kindObjectStore       = "ost"
	kindQueue             = "que"
	kindRDB               = "rdb"
	kindService           = "ser"
	kindUser              = "usr"
	kindFirewall          = "waf"
	kindWeb               = "web"
)

// validateProvider sets a valid cloud provider (one of 'aws', 'gcp', 'azure')
func validateProvider() func(com *Component) {
	provs := map[string]bool{"aws": true, "google": true, "azure": true}

	return func(com *Component) {
		val := strings.ToLower(strings.TrimSpace(com.Provider))
		if provs[val] {
			com.Provider = val
		} else {
			com.Provider = "default"
		}
	}

}

func validateKind() func(com *Component) {
	kinds := map[string]bool{
		kindBlockStore:        true,
		kindCDN:               true,
		kindClient:            true,
		kindDNS:               true,
		kindNoSQL:             true,
		kindFileStore:         true,
		kindFunction:          true,
		kindGateway:           true,
		kindContainersManager: true,
		kindLBA:               true,
		kindCache:             true,
		kindPubSub:            true,
		kindObjectStore:       true,
		kindQueue:             true,
		kindRDB:               true,
		kindService:           true,
		kindUser:              true,
		kindFirewall:          true,
		kindWeb:               true,
	}

	return func(com *Component) {
		val := strings.ToLower(strings.TrimSpace(com.Kind))
		if kinds[val] {
			com.Kind = val
		} else {
			com.Kind = kindUnknown
		}
	}
}
