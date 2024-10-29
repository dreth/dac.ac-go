package constants

import (
	"fmt"
	"os"
)

// CDNUrl Variable
var CDNUrl = GetCDNUrl()

// GetCDNUrl picks randomly between the available CDN URLs for every request
func GetCDNUrl() string {
	if os.Getenv("ENVIRONMENT") == "production" {
		if cdn := os.Getenv("CDN_URL"); cdn != "" {
			return cdn
		}
		return "https://cdn.dac.sg"
	}
	return ""
}

// Appends a random CDN url to the resource
func ResourceWithCDN(resource string) string {
	return fmt.Sprintf("%s%s", CDNUrl, resource)
}
