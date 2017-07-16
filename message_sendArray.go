package kavenegar

import (
	"net/url"
	"strings"
)

//SendArray ...
func (m *MessageService) SendArray(sender []string, receptor []string, message []string, params *MessageParam) ([]Message, error) {
	v := url.Values{}
	v.Set("sender", strings.Join(sender, ","))
	v.Set("receptor", strings.Join(receptor, ","))
	v.Set("message", strings.Join(message, ","))
	if params != nil {
		if !params.Date.IsZero() {
			v.Set("date", TimeToUnix(params.Date))
		}
		if params.Type != nil {
			v.Set("type", ArrayEncodeToString(params.Type))
		}
		if params.LocalID != nil {
			v.Set("localid", ArrayEncodeToString(params.LocalID))
		}
	}
	return m.CreateSend(v)
}

//CreateSendArray ...
func (m *MessageService) CreateSendArray(v url.Values) ([]Message, error) {
	u := m.client.EndPoint("sms", "sendarray")
	vc := url.Values{}
	if v.Get("sender") != "" {
		vc.Set("sender", v.Get("sender"))
	}
	if v.Get("message") != "" {
		vc.Set("message", v.Get("message"))
	}
	if v.Get("receptor") != "" {
		vc.Set("receptor", v.Get("receptor"))
	}
	if v.Get("date") != "" {
		vc.Set("date", v.Get("date"))
	}
	if v.Get("type") != "" {
		vc.Set("type", v.Get("type"))
	}
	if v.Get("Localid") != "" {
		vc.Set("Localid", v.Get("Localid"))
	}
	res := new(MessageResult)
	err := m.client.Execute(u.String(), v, res)
	return res.Entries, err
}
