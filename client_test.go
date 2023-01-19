package client

import (
	"testing"

	simplejson "github.com/bitly/go-simplejson"
	"gopkg.in/stretchr/testify.v1/assert"
)

func TestApi(t *testing.T) {
	c := New()
	token, ok := c.Auth("public_user", "public_pass")
	t.Log("token: ", token)
	if !ok {
		t.Error("Authentication failed!")
	}

	globalRecovered2023 := c.GlobalRecovered(2023)
	globalRecovered2023.Name = "Time Series - Recovered Global 2023"
	endpoints := []*Resource{globalRecovered2023, c.GlobalConfirmed(2023), c.GlobalDeath(2023), c.UsConfirmed(2023), c.UsDeath(2023)}
	var res *simplejson.Json
	var err error
	var id string
	for _, endpoint := range endpoints {
		t.Log("begin to test " + endpoint.Name + "'s GetAll api")
		res, err = endpoint.GetAll()
		assert.Empty(t, err, endpoint.Name+" GetAll Api test failed")

		id = res.Get("Document").GetIndex(0).Get("id").MustString()

		t.Log("begin to test " + endpoint.Name + "'s GetByID api, ID: " + id)
		_, err = endpoint.GetById(id)
		assert.Empty(t, err, endpoint.Name+" GetByID Api test failed")
	}

}
