package kavenegar

// VerifyService ...
type VerifyService struct {
	client *Client
}

// VerifyLookupResult ...
type VerifyLookupResult struct {
	*Return `json:"return"`
	Entries Message `json:"entries"`
}

// VerifyLookupParam ...
type VerifyLookupParam struct {
	Token2  string
	Token3  string
	Token10 string
	Token20 string
	Type    VerifyLookupType
}

// NewVerifyService ...
func NewVerifyService(client *Client) *VerifyService {
	m := &VerifyService{client: client}
	return m
}

type VerifyLookupType int

const (
	Type_VerifyLookup_Sms VerifyLookupType = iota
	Type_VerifyLookup_Call
)

var VerifyLookupTypeMap = map[VerifyLookupType]string{
	Type_VerifyLookup_Sms:  "0",
	Type_VerifyLookup_Call: "1",
}

func (t VerifyLookupType) String() string {
	return VerifyLookupTypeMap[t]
}
