package constant

const ApiUrl = "https://api.hostker.net/v2"

// DomainType 域名类型：https://doc.hostker.dev/dns/createRecord.html
type DomainType string

const (
	DomainTypeA     DomainType = "A"
	DomainTypeCname DomainType = "CNAME"
	DomainTypeAAAA  DomainType = "AAAA"
	DomainTypeMX    DomainType = "MX"
	DomainTypeTXT   DomainType = "TXT"
	DomainTypeSRV   DomainType = "SRV"
	DomainTypeNS    DomainType = "NS"
	DomainTypeCAA   DomainType = "CAA"
)
