package ipify

import (
  "fmt"
  "net/http"
)

func bodyDispose(r *http.Response) {
			_ = resp.Body.Close()
}

// PublicIP retrieves the current public ip from IPIFy https://www.ipify.org/
func PublicIP()(string,error) {
	if resp, err := http.Get(`https://api.ipify.org?format=json`); err != nil {
		return "",err
	} else {
		defer bodyDispose(resp)

		var ipify struct {
			IP string `json:"ip"`
		}

		return ipify.IP
	}
}
