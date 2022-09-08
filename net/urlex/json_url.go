package urlex

import (
	"encoding/json"
	"net/url"
)

type JSONURL struct {
	url.URL
}

func NewJSONURL(u *url.URL) *JSONURL {
	if u == nil {
		return nil
	}

	return &JSONURL{URL: *u}
}

func (ju *JSONURL) URLCopy() *url.URL {
	return (&ju.URL).ResolveReference(&url.URL{})
}

func (ju *JSONURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(ju.URL.String())
}

func (ju *JSONURL) UnmarshalJSON(b []byte) error {
	s := ""
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if s == "" {
		(*ju) = JSONURL{URL: url.URL{}}
		return nil
	}

	parsed, err := url.Parse(s)
	if err != nil {
		return err
	}

	(*ju) = JSONURL{URL: *parsed}

	return nil
}
