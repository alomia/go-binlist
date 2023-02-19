package card

import (
	"testing"
)

type cardTest struct {
	bin		string
	scheme	string
}

func TestGetInfo(t *testing.T) {
	tests := []struct {
		name   			string
		bin     		string
		scheme  		string
		cardType		string
		countryName		string
	}{
		{
			name:			"Valid_card",
			bin:			"45717360",
			scheme:			"visa",
			cardType:		"debit",
			countryName:	"Denmark",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cardInfo, err := GetInfo(tt.bin)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if cardInfo.Scheme != tt.scheme {
				t.Errorf("unexpected scheme result: got %q, want %q", cardInfo.Scheme, tt.scheme)
			}

			if cardInfo.Type != tt.cardType {
				t.Errorf("unexpected type result: got %q, want %q", cardInfo.Type, tt.cardType)
			}

			if cardInfo.Country.Name != tt.countryName {
				t.Errorf("unexpected type result: got %q, want %q", cardInfo.Country.Name, tt.countryName)
			}
		})
	}
}

func TestIsValidCardNumber(t *testing.T) {
    testCases := []struct {
        name    string
        bin     string
        want    bool
    }{
        {
            name:    "Valid card number",
            bin:     "1234567890123456",
            want:    true,
        },
		{
            name:    "Invalid card number (too short)",
            bin:     "12345",
            want:    false,
        },
		{
            name:    "Invalid card number (too long)",
            bin:     "12345678901234567890",
            want:    false,
        },
        {
            name:    "Invalid card number (contains letters)",
            bin:     "12345678a0123456",
            want:    false,
        },
    }

    for _, tt := range testCases {
        t.Run(tt.name, func(t *testing.T) {
            got := isValidCardNumber(tt.bin)
            if got != tt.want {
                t.Errorf("isValidCardNumber(%q) = %v, want %v", tt.bin, got, tt.want)
            }
        })
    }
}

