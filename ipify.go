package ipify

import (
	"encoding/json"
	"net/http"
	"time"
)

const timeoutSecondsMin = 1

func bodyDispose(r *http.Response) {
	if r != nil {
		_ = r.Body.Close()
	}
}

// PublicIP retrieves the current public ip from IPIFy https://www.ipify.org/
func PublicIP(timeoutSeconds int) (string, error) {
	if timeoutSeconds < timeoutSecondsMin {
		timeoutSeconds = timeoutSecondsMin
	}

	cli := &http.Client{Timeout: time.Duration(timeoutSeconds) * time.Second}

	resp, err := cli.Get(`https://api.ipify.org?format=json`)

	if err != nil {
		return "", err
	}

	defer bodyDispose(resp)

	var ipify struct {
		IP string `json:"ip"`
	}

	if err0 := json.NewDecoder(resp.Body).Decode(&ipify); err0 != nil {
		return "", err0
	}

	return ipify.IP, nil
}

// PublicIPString retrieves the current public ip from IPIFy https://www.ipify.org/
// The differences to PublicIP are:
// - PublicIPString ignores all errors and return an empty string if fails
// - PublicIPString defines arbitratry timeout value of 5 seconds
func PublicIPString() string {
	const timeoutSeconds = 5

	ip, _ := PublicIP(timeoutSeconds)

	return ip
}
