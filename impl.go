package draft

import (
	"strings"
)

func guessImplByProvider(prov string) func(*Component) {
	return func(com *Component) {
		if s := strings.TrimSpace(com.Impl); len(s) > 0 {
			return
		}

		impl := getCloudImpl(prov, com.Kind)
		if len(impl) > 0 {
			com.Impl = impl
		}
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
		return ""
	}
}

func awsImpl() func(string) string {
	dict := map[string]string{
		kindBlockStore:        "Elastic Block\nStore (EBS)",
		kindCDN:               "Cloudfront",
		kindDNS:               "Route 53",
		kindNoSQL:             "DynamoDB",
		kindFileStore:         "Elastic\nFile System (EFS)",
		kindFunction:          "Lambda",
		kindGateway:           "API Gateway",
		kindLBA:               "Elastic\nLoad Balancer",
		kindContainersManager: "Elastic Container Service\nfor Kubernetes (EKS)",
		kindCache:             "ElastiCache",
		kindPubSub:            "SNS",
		kindObjectStore:       "Simple Storage\nService (S3)",
		kindRDB:               "Relational Database\nService (RDS)",
		kindFirewall:          "AWS WAF",
	}

	return func(key string) string {
		return dict[key]
	}
}

func googleImpl() func(string) string {
	dict := map[string]string{
		kindBlockStore:        "Persistent Disk",
		kindCDN:               "Cloud CDN",
		kindDNS:               "Cloud DNS",
		kindNoSQL:             "Cloud Datastore",
		kindFileStore:         "Cloud Filestore",
		kindFunction:          "Cloud Functions",
		kindGateway:           "Cloud Endpoints",
		kindLBA:               "Cloud Load Balancing",
		kindContainersManager: "Google Kubernetes\nEngine",
		kindCache:             "Cloud Memorystore",
		kindPubSub:            "Cloud Pub/Sub",
		kindObjectStore:       "Cloud Storage",
		kindRDB:               "Cloud SQL",
		kindFirewall:          "Google Armor",
	}

	return func(key string) string {
		return dict[key]
	}
}

func azureImpl() func(string) string {
	dict := map[string]string{
		kindBlockStore:        "Disk Storage",
		kindCDN:               "Azure CDN",
		kindDNS:               "Azure DNS",
		kindNoSQL:             "Cosmos DB",
		kindFileStore:         "Azure File Storage",
		kindFunction:          "Azure Functions",
		kindGateway:           "Azure API Management",
		kindLBA:               "Load Balancer",
		kindContainersManager: "Azure Kubernetes\nService (AKS)",
		kindCache:             "Redis Caches",
		kindPubSub:            "Notification Hubs",
		kindObjectStore:       "Blob Storage",
		kindRDB:               "SQL Database",
		kindFirewall:          "Azure Firewall",
	}

	return func(key string) string {
		return dict[key]
	}
}
