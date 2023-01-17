package client

import "strconv"

type Client struct {
	http *HttpClient
}

var c *Client

func New() *Client {

	httpClient := newHttpClient()
	c = &Client{
		http: httpClient,
	}
	return c
}

func getClient() *Client {
	return c
}

func (c *Client) Auth(username, password string) (string, bool) {
	data := make(map[string]interface{})
	data["username"] = username
	data["password"] = password
	json, err := c.http.request("POST", "/token", data)
	if err != nil {
		return "", false
	}
	t, err := json.Get("Document").String()
	if t == "" || err != nil {
		return "", false
	}
	c.http.Token = t
	return t, true
}

func (c *Client) UsDeath(year int) *Resource {
	r := NewResource("UsDeath", "/v2/time_series_deaths_us"+"/"+strconv.Itoa(year))
	return r
}

func (c *Client) UsConfirmed(year int) *Resource {
	r := NewResource("UsConfirmed", "/v2/time_series_confirmed_us"+"/"+strconv.Itoa(year))
	return r
}

func (c *Client) GlobalDeath(year int) *Resource {
	r := NewResource("GlobalDeath", "/v2/time_series_deaths_global"+"/"+strconv.Itoa(year))
	return r
}

func (c *Client) GlobalConfirmed(year int) *Resource {
	r := NewResource("GlobalConfirmed", "/v2/time_series_confirmed_global"+"/"+strconv.Itoa(year))
	return r
}

func (c *Client) GlobalDailyReport(month int, year int) *Resource {
	months := map[int]string{
		1:  "jan",
		2:  "feb",
		3:  "mar",
		4:  "apr",
		5:  "may",
		6:  "jun",
		7:  "jul",
		8:  "aug",
		9:  "sep",
		10: "oct",
		11: "nov",
		12: "dec",
	}
	monStr := months[month]
	r := NewResource("GlobalDailyReport", "/v2/"+monStr+"/"+strconv.Itoa(year))
	return r
}

func (c *Client) GlobalRecovered(year int) *Resource {
	r := NewResource("GlobalRecovered", "/v2/time_series_recovered_global"+"/"+strconv.Itoa(year))
	return r
}
