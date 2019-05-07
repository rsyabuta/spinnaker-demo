package main

// func main() {
// 	tlsConfig := &tls.Config{}
// 	conf := api.DefaultConfig()
// 	conf.Address = "https://vault-server.in.quid.com:443"
// 	conf.HttpClient.Transport = &http.Transport{
// 		TLSClientConfig: tlsConfig,
// 	}
// 	c, err := api.NewClient(conf)
// 	if err != nil {
// 		log.Fatalf("client %v", err)
// 	}
// 	h, err := c.Sys().Health()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%+v\n", h)
// }

// func main() {
// 	client := cleanhttp.DefaultClient()
// 	transport := client.Transport.(*http.Transport)
// 	transport.TLSHandshakeTimeout = 10 * time.Second
// 	transport.TLSClientConfig = &tls.Config{
// 		MinVersion: tls.VersionTLS12,
// 	}
// 	if err := http2.ConfigureTransport(transport); err != nil {
// 		log.Fatalf("http2 %v", err)
// 	}
//
// 	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
// 		// Returning this value causes the Go net library to not close the
// 		// response body and to nil out the error. Otherwise retry clients may
// 		// try three times on every redirect because it sees an error from this
// 		// function (to prevent redirects) passing through to it.
// 		return http.ErrUseLastResponse
// 	}
//
// 	// resp, err := client.Get("https://vault-server.in.quid.com/v1/sys/health")
// 	resp, err := retryablehttp.Get("https://vault-server.in.quid.com/v1/sys/health")
// 	if err != nil {
// 		log.Fatalf("client %v", err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalf("body %v", err)
// 	}
// 	fmt.Println(string(body))
// }
