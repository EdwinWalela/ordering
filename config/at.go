package config

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"edwinwalela/ordering/models"
)

const (
	AT_BASE_URL = "https://api.sandbox.africastalking.com/version1/messaging"
)

type ATSMS struct {
	C *Config
}

func (a *ATSMS) SendMessage(msg models.Message) error {
	values := url.Values{
		"username": {a.C.ATuser},
		"to":       {"0700000000"},
		"message":  {fmt.Sprintf("Thank you %s for your order of %s. Our rider will contact you soon.", msg.Recipient, msg.Item)},
	}
	reader := strings.NewReader(values.Encode())

	req, err := http.NewRequest(http.MethodPost, AT_BASE_URL, reader)
	if err != nil {
		fmt.Println(err)
		return err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Length", strconv.Itoa(reader.Len()))
	req.Header.Set("apikey", a.C.ATKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	fmt.Println(res)
	fmt.Println(err)
	return err
}
