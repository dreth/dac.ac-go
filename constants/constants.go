package constants

import (
	"fmt"
	"os"
)

// CDNUrl is the CDN url (cdn.dac.ac)
// Only set in production environment
var CDNUrl = GetCDNUrl()

func GetCDNUrl() string {
	if os.Getenv("ENVIRONMENT") == "production" {
		return "https://cdn.dac.ac"
	} else {
		return ""
	}
}

// Appends the CDN url to the resource
func ResourceWithCDN(resource string) string {
	return fmt.Sprintf("%s%s", CDNUrl, resource)
}
