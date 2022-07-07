package jwt

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	sampleToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	expected := map[string]interface{}{
		"iat":  float64(1516239022),
		"name": "John Doe",
		"sub":  "1234567890",
	}

	resultStr := Parse(sampleToken)
	result := make(map[string]interface{})

	json.Unmarshal([]byte(resultStr), &result)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, result)
	}
}
