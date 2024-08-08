// package auth

// import (
// 	"testing"
// 	"net/http"
// )

// func TestGetAPIKey(*testing.T) {
// 	header := http.Header()
// 	got1, got2 := GetAPIKey("Connection: keep-alive")
// 	want := "result"

//     if !reflect.DeepEqual(want, got1) {
// 		t.Fatalf("expected: %v, got: %v", want, got)
//    }
// }


package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		// {
		// 	name:          "No Authorization Header",
		// 	headers:       http.Header{},
		// 	expectedKey:   "",
		// 	expectedError: ErrNoAuthHeaderIncluded,
		// }//,
		// {
		// 	name: "Malformed Authorization Header",
		// 	headers: http.Header{
		// 		"Authorization": []string{"Bearer token"},
		// 	},
		// 	expectedKey:   "",
		// 	expectedError: errors.New("malformed authorization header"),
		// },
		{
			name: "Valid Authorization Header",
			headers: http.Header{
				"Authorization": []string{"ApiKey myapikey"},
			},
			expectedKey:   "myapikey",
			expectedError: nil,
		},
		// {
		// 	name: "Invalid Authorization Scheme",
		// 	headers: http.Header{
		// 		"Authorization": []string{"ApiKey"},
		// 	},
		// 	expectedKey:   "",
		// 	expectedError: errors.New("malformed authorization header"),
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			if key != tt.expectedKey {
				t.Errorf("expected key %v, got %v", tt.expectedKey, key)
			}
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
			if err == nil && tt.expectedError != nil {
				t.Errorf("expected error %v, got nil", tt.expectedError)
			}
		})
	}
}
