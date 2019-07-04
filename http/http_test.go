package http

import "testing"

func TestConvertToQueryParams(t *testing.T) {
	str := ConvertToQueryParams(map[string]string{
		"a": "aa",
		"b": "bb",
	})
	t.Log(str)
}
