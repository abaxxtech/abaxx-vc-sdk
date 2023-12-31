package util

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL_MarshalJSON(t *testing.T) {
	u := URL{url.URL{Scheme: "https", Host: "example.com"}}

	got, err := u.MarshalJSON()

	expected := []byte(`"https://example.com"`)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
}

func TestURL_UnmarshalJSON(t *testing.T) {
	data := []byte(`"dwn://abaxx.id"`)
	var u URL

	assert.NoError(t, u.UnmarshalJSON(data))

	expected := URL{url.URL{Scheme: "dwn", Host: "abaxx.id"}}
	assert.Equal(t, expected, u)
}
