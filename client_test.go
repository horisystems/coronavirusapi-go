package client

import (
	"fmt"
	"testing"
)

func TestApi(t *testing.T) {
	c := New()
	token, ok := c.Auth("public_user", "public_pass")
	t.Log("token: ", token)
	if !ok {
		t.Error("Authentication failed!")
	}

	// j, err := c.GlobalDailyReport(8, 2022).GetById(0)
	// fmt.Printf("Global Daily Report GetById: %+v", j)
	// if err != nil {
	// 	t.Error("Global Daily Report GetById failed!")
	// }

	j, err := c.GlobalRecovered(2023).GetAll()
	fmt.Printf("Global Recovered Data GetAll: %+v", j)
	if err != nil {
		t.Error("Global Recovered Data GetAll failed!")
	}

}
