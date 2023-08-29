package api

func (ker *Ker) getDomainList() {
	const action = "DomainDNS.getDomainList"

	if response, err := ker.fetch(map[string]interface{}{"action": action}); err != nil {
		return
	}
}
