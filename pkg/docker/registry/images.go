package registry

import (
	"fmt"
	"strings"
)

func MakeProxiedImageURL(proxyHost string, appSlug string, image string) string {
	fmt.Printf("+++++convert %s ->\n", image)
	parts := strings.Split(image, "@")
	if len(parts) == 2 {
		ttt := strings.Join([]string{proxyHost, "proxy", appSlug, parts[0]}, "/")
		fmt.Printf("+++++to 1 %s\n", ttt)
		return ttt
	}

	// TODO: host with a port breaks this
	parts = strings.Split(image, ":")
	ttt := strings.Join([]string{proxyHost, "proxy", appSlug, parts[0]}, "/")
	fmt.Printf("+++++to 2 %s\n", ttt)
	return ttt
}
