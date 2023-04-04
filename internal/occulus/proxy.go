package occulus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Proxy struct {
	authToken    string // holds the authentication token
	logger       *logrus.Logger
	serverconfig *eqemuserverconfig.Config
	process      *ProcessManagement
}

func NewProxy(
	logger *logrus.Logger,
	serverconfig *eqemuserverconfig.Config,
	process *ProcessManagement,
) *Proxy {
	return &Proxy{
		logger:       logger,
		serverconfig: serverconfig,
		process:      process,
	}
}

func (o *Proxy) GetBaseUrl() string {
	return fmt.Sprintf("http://localhost:%v", o.process.Port())
}

func (o *Proxy) AuthToken() string {
	return o.authToken
}

func (o *Proxy) SetAuthToken(authToken string) {
	o.authToken = authToken
}

// LoginRequestBody is the payload structure for logging into Occulus
type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login is the main method used to logging into the Occulus admin panel
func (o *Proxy) Login() error {
	config := o.serverconfig.Get()

	postBody := LoginRequestBody{
		Username: "admin",
		Password: config.WebAdmin.Application.Admin.Password,
	}

	payload, err := json.Marshal(postBody)
	if err != nil {
		return err
	}

	client := o.getHttpClient()
	resp, err := client.Post(
		fmt.Sprintf("%v/api/v1/auth/login", o.GetBaseUrl()),
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

func (o *Proxy) GetProcessCounts() (GetProcessCountsResponse, error) {
	endpoint := fmt.Sprintf("%v/api/v1/server/process_counts", o.GetBaseUrl())
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

func (o *Proxy) ProxyRequest(c echo.Context) (*http.Response, []byte, error) {
	err := o.Login()
	if err != nil {
		return nil, nil, err
	}

	// request object
	r := c.Request().Clone(c.Request().Context())

	// rewrites
	requestUrl := strings.ReplaceAll(r.RequestURI, "/admin/occulus", "")
	r.RequestURI = ""
	r.Host = o.GetBaseUrl()

	targetUrl, err := url.Parse(fmt.Sprintf("%v%v", o.GetBaseUrl(), requestUrl))
	if err != nil {
		return nil, nil, err
	}

	r.URL = targetUrl
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %v", o.AuthToken()))

	client := o.getHttpClient()
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

func (o *Proxy) getHttpClient() http.Client {
	return http.Client{
		Timeout: 300 * time.Second,
	}
}
