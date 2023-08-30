package constant

const ApiUrl = "https://api.hostker.net/v2"

type DomainType string

const (
	DomainTypeA     DomainType = "A"
	DomainTypeCname DomainType = "CNAME"
	DomainTypeAAAA  DomainType = "AAAA"
	DomainTypeMX    DomainType = "MX"
	DOmainTypeTXT   DomainType = "TXT"
	DomainTypeSRV   DomainType = "SRV"
	DomainTypeNS    DomainType = "NS"
	DomainTypeCAA   DomainType = "CAA"
)
