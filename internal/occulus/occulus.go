package occulus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/serverconfig"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type OcculusProxy struct {
	authToken    string // holds the authentication token
	logger       *logrus.Logger
	serverconfig *serverconfig.EQEmuServerConfig
}

func NewOcculusProxy(
	logger *logrus.Logger,
	serverconfig *serverconfig.EQEmuServerConfig,
) *OcculusProxy {
	return &OcculusProxy{
		logger:       logger,
		serverconfig: serverconfig}
}

const (
	baseUrl = "http://localhost:3000"
)

func (o *OcculusProxy) AuthToken() string {
	return o.authToken
}

func (o *OcculusProxy) SetAuthToken(authToken string) {
	o.authToken = authToken
}

// LoginRequestBody is the payload structure for logging into Occulus
type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login is the main method used to logging into the Occulus admin panel
func (o *OcculusProxy) Login() error {
	config := o.serverconfig.Get()

	postBody := LoginRequestBody{
		Username: "admin",
		Password: config.WebAdmin.Application.Admin.Password,
	}

	payload, err := json.Marshal(postBody)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		fmt.Sprintf("%v/api/v1/auth/login", baseUrl),
		"application/json",
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	type LoginResponse struct {
		Success     string `json:"success"`
		AccessToken string `json:"access_token"`
	}

	var loginResponse LoginResponse
	err = json.Unmarshal(b, &loginResponse)
	if err != nil {
		return err
	}

	o.SetAuthToken(loginResponse.AccessToken)

	return nil
}

type GetProcessCountsResponse struct {
	Zone        int `json:"zone"`
	World       int `json:"world"`
	Ucs         int `json:"ucs"`
	Queryserv   int `json:"queryserv"`
	Loginserver int `json:"loginserver"`
}

func (o *OcculusProxy) GetProcessCounts() (GetProcessCountsResponse, error) {
	endpoint := fmt.Sprintf("%v/api/v1/server/process_counts", baseUrl)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return GetProcessCountsResponse{}, nil
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", o.AuthToken()))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return GetProcessCountsResponse{}, nil
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return GetProcessCountsResponse{}, nil
	}

	var processCountsResponse GetProcessCountsResponse
	err = json.Unmarshal(b, &processCountsResponse)
	if err != nil {
		return GetProcessCountsResponse{}, nil
	}

	return processCountsResponse, nil
}

func (o *OcculusProxy) ProxyRequest(c echo.Context) (*http.Response, []byte, error) {
	err := o.Login()
	if err != nil {
		return nil, nil, err
	}

	// request object
	r := c.Request().Clone(c.Request().Context())

	// rewrites
	requestUrl := strings.ReplaceAll(r.RequestURI, "/admin/occulus", "")
	r.RequestURI = ""
	r.Host = baseUrl

	targetUrl, err := url.Parse(fmt.Sprintf("%v%v", baseUrl, requestUrl))
	if err != nil {
		return nil, nil, err
	}

	r.URL = targetUrl
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %v", o.AuthToken()))

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(r)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	defer res.Body.Close()

	return res, body, nil
}
