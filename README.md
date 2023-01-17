<div align="center">

  <img alt="coronavirusapi-go logo" src="assets/go.png" height="150" />
  <h3 align="center">coronavirusapi-go</h3>
  <p align="center"><a href="https://moatsystems.com/covid19-api/">COVID-19 API</a> SDK for Golang. 🦠 </p>

[![Release](https://img.shields.io/github/release/moatsystems/coronavirusapi-go.svg)](https://github.com/moatsystems/coronavirusapi-go/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/moatsystems/coronavirusapi-go)](https://goreportcard.com/report/github.com/moatsystems/coronavirusapi-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/moatsystems/coronavirusapi-go.svg)](https://pkg.go.dev/github.com/moatsystems/coronavirusapi-go)
[![License](https://img.shields.io/github/license/moatsystems/coronavirusapi-go)](/LICENSE)

</div>

---

### Getting Started

Import Covid-19 API Golang module:
 
```golang
import "github.com/moatsystems/coronavirusapi-go"
```

### Dependence

- [simplejson](https://github.com/bitly/go-simplejson)
- [resty](https://github.com/go-resty/resty)


### Testing

Run the following command to test the package:

```sh
go test
```

### Implementation

```golang
c: = client.New()
token,ok := c.Auth("<username>", "<password>")
...

// Call an endpoint such as US Death
usd := client.UsDeath()
  
// Retrieve all US Death data
j, err := usd.GetAll()

// Get API result (see simplejson)
j.Get("Code").String()
j.Get("Message").String()
j.Get("Document").Array()

// GetById
usd.GetById(1)
```

Below are all the available endpoints:

```sh
- UsDeath  
- UsConfirmed  
- GlobalDeath  
- GlobalConfirmed  
- GlobalRecovered  
- GlobalDailyReport 
```

All the endpoints listed above have five methods:

```sh
- GetAll()
- GetById(id int)
- Create(data map[string]interface{})
- UpdateById(id int, data map[string]interface{})
- DeleteById(id int)  
```

The COVID-19 API documentation is available [here](https://docs.covid19api.dev/). If you need further assistance, don't hesitate to [contact us](https://moatsystems.com/contact/).

### Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

### Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/moatsystems/coronavirusapi-go/tags).

### License

This project is licensed under the [BSD 3-Clause License](LICENSE) - see the file for details.

### Copyright

(c) 2020 - 2023 [Moat Systems Limited](https://moatsystems.com).
