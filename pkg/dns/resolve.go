package dns

import (
	"fmt"
	"net"
)

// Resolve service ip:port from dns srv record
func Resolve(service, proto, name string) (string, error) {
	cname, addrs, err := net.LookupSRV(service, proto, name)
	if err != nil {
		return "", err
	}

	if len(addrs) == 0 {
		return "", fmt.Errorf("SRV Lookup for %q service not found", service)
	}

	addr := fmt.Sprintf("%s:%d", cname, addrs[0].Port)

	return addr, nil
}
