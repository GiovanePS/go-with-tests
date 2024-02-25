package concorrencia

type VerificadorWebsite func(string) bool

func VerificaWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = vw(url)
	}

	return results
}
