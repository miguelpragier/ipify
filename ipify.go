package ipify

import (
	"net/http"
)

func bodyDispose(r *http.Response) {
	if r != nil {
		_ = r.Body.Close()
	}
}

// PublicIP retrieves the current public ip from IPIFy https://www.ipify.org/
func PublicIP() (string, error) {
	if resp, err := http.Get(`https://api.ipify.org?format=json`); err != nil {
		return "", err
	} else {
		defer bodyDispose(resp)

		var ipify struct {
			IP string `json:"ip"`
		}

		return ipify.IP, nil
	}
}
