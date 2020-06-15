package draft

import (
	"strings"

	"github.com/emicklei/dot"
)

func getCloudImpl(provider, kind string) string {
	switch strings.TrimSpace(strings.ToLower(provider)) {
	case "aws":
		return awsImpl()(kind)
	case "gcp":
		return gcpImpl()(kind)
	case "azure":
		return azureImpl()(kind)
	default:
		return ""
	}
}

func awsImpl() func(string) string {
	dict := map[string]string{
		"bst": "Elastic Block\nStore (EBS)",
		"cdn": "Cloudfront",
		"dns": "Route 53",
		"doc": "DynamoDB",
		"fst": "Elastic\nFile System (EFS)",
		"fun": "Lambda",
		"gtw": "API Gateway",
		"lba": "Elastic\nLoad Balancer",
		"kub": "Elastic Container Service\nfor Kubernetes (EKS)",
		"mem": "ElastiCache",
		"msg": "SNS",
		"ost": "Simple Storage\nService (S3)",
		"rdb": "Relational Database\nService (RDS)",
		"waf": "AWS WAF",
	}

	return func(key string) string {
		return dict[key]
	}
}

func gcpImpl() func(string) string {
	dict := map[string]string{
		"bst": "Persistent Disk",
		"cdn": "Cloud CDN",
		"dns": "Cloud DNS",
		"doc": "Cloud Datastore",
		"fst": "Cloud Filestore",
		"fun": "Cloud Functions",
		"gtw": "Cloud Endpoints",
		"lba": "Cloud Load Balancing",
		"kub": "Google Kubernetes\nEngine",
		"mem": "Cloud Memorystore",
		"msg": "Cloud Pub/Sub",
		"ost": "Cloud Storage",
		"rdb": "Cloud SQL",
		"waf": "Google Armor",
	}

	return func(key string) string {
		return dict[key]
	}
}

func azureImpl() func(string) string {
	dict := map[string]string{
		"bst": "Disk Storage",
		"cdn": "Azure CDN",
		"dns": "Azure DNS",
		"doc": "Cosmos DB",
		"fst": "Azure File Storage",
		"fun": "Azure Functions",
		"gtw": "Azure API Management",
		"lba": "Load Balancer",
		"kub": "Azure Kubernetes\nService (AKS)",
		"mem": "Redis Caches",
		"msg": "Notification Hubs",
		"ost": "Blob Storage",
		"rdb": "SQL Database",
		"waf": "Azure Firewall",
	}

	return func(key string) string {
		return dict[key]
	}
}

func guessImpl(cmp *Component, cluster *dot.Graph) {
	if s := strings.TrimSpace(cmp.Impl); len(s) > 0 {
		return
	}

	impl := getCloudImpl(cmp.provider, cmp.Kind)
	if len(impl) > 0 {
		cluster.Attr("label", impl)
	}
}
