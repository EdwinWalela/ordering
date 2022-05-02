package handlers

import (
	"edwinwalela/ordering/models"
	r "edwinwalela/ordering/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Repo r.Repository
}

type Response struct {
	UserCode        string `json:"user_code"`
	DeviceCode      string `json:"device_code"`
	VerificationUrI string `json:"verification_uri"`
}

type Profile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccessTokenRes struct {
	Token string `json:"access_token"`
}

func unmarshalResponse(res *http.Response) (Response, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	r := Response{}

	err = json.Unmarshal(body, &r)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}
	return r, nil
}

func unmarshalAuthToken(res *http.Response) (AccessTokenRes, error) {
	tokenRes := AccessTokenRes{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return AccessTokenRes{}, err
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &tokenRes)
	if err != nil {
		fmt.Println(err)
		return AccessTokenRes{}, err
	}
	return tokenRes, nil
}

func unmarshalProfile(res *http.Response) (Profile, error) {
	profile := Profile{}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Profile{}, err
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &profile)
	if err != nil {
		fmt.Println(err)
		return Profile{}, err
	}
	return profile, nil
}

func getCustomerProfile(token string) (Profile, error) {
	base_url := "https://api.github.com/user"
	req, err := http.NewRequest(http.MethodGet, base_url, nil)
	if err != nil {
		fmt.Println(err)
		return Profile{}, err
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Profile{}, err
	}
	profile, err := unmarshalProfile(res)
	if err != nil {
		fmt.Println(err)
		return Profile{}, err
	}
	return profile, nil
}

func (h *Handlers) VerifyOauth(c *gin.Context) {
	device_code, _ := c.GetQuery("device_code")
	const client_id = "Iv1.4a7c71ee0c0c6c1b"
	base_url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&device_code=%s&grant_type=urn:ietf:params:oauth:grant-type:device_code", client_id, device_code)
	req, err := http.NewRequest(http.MethodPost, base_url, nil)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	authTokenRes, err := unmarshalAuthToken(res)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if authTokenRes.Token == "" {
		c.JSON(200, gin.H{
			"msg": "Failed to authenticate",
		})
		return
	}

	profile, err := getCustomerProfile(authTokenRes.Token)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if authTokenRes.Token == "" {
		c.JSON(200, gin.H{
			"msg": "Failed to retrieve profile",
		})
		return
	}

	customer := models.Customer{
		Name: profile.Name,
	}

	id, err := h.Repo.CreateCustomer(customer)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "Account created",
		"id":  id,
	})
}

func (h *Handlers) RegisterCustomer(c *gin.Context) {
	const base_url = "https://github.com/login/device/code?client_id=Iv1.4a7c71ee0c0c6c1b&scope=email"
	req, err := http.NewRequest(http.MethodPost, base_url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)

	marshaledRes, err := unmarshalResponse(res)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"redirect":    marshaledRes.VerificationUrI,
		"code":        marshaledRes.UserCode,
		"device-code": marshaledRes.DeviceCode,
	})
}

func (h *Handlers) CreateCustomer(c *gin.Context) {
	var customer models.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create customer",
			"error":   err.Error(),
		})
		return
	}

	id, err := h.Repo.CreateCustomer(customer)

	if err != nil {
		fmt.Printf("Failed to insert customer to DB: %v", err)
		c.JSON(500, gin.H{
			"message": "Failed to create customer",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message":  "Customer created",
		"customer": id,
	})
}

func (h *Handlers) GetCustomers(c *gin.Context) {
	customers, err := h.Repo.GetCustomers()
	if err != nil {
		log.Printf("Failed to query DB: %v", err)
		c.JSON(500, gin.H{
			"message": "Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"customers": customers,
	})
}

func (h *Handlers) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create order",
			"error":   err.Error(),
		})
		return
	}
	orderId, err := h.Repo.CreateOrder(order)

	if err != nil {
		log.Printf("Failed to create order: %v", err)
		c.JSON(500, gin.H{
			"message": "Failed to create order",
			"error":   err.Error(),
		})
		return
	}

	customer := h.Repo.GetCustomerById(order.CustomerId)

	message := models.Message{
		Recipient: customer.Name,
		Item:      order.Item,
	}

	err = h.Repo.SmsService.SendMessage(message)
	fmt.Println(err)
	c.JSON(201, gin.H{
		"message": "Order created",
		"id":      orderId,
	})
}

func (h *Handlers) GetOrders(c *gin.Context) {
	orders, err := h.Repo.GetOrders()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create order",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"orders": orders,
	})
}
