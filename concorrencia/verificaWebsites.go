package concorrencia

type VerificadorWebsite func(string) bool
type result struct {
	string
	bool
}

func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, vw(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
