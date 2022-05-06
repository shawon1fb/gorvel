package validations

import "github.com/lucidfy/lucid/pkg/rules/must"

func UserValidateCreate() *must.SetOfRules {
	return &must.SetOfRules{
		"name": {
			&must.Required{},
			&must.Min{Value: 4},
		},
		"email": {
			&must.Email{},
		},
		"password": {
			&must.Min{Value: 5},
			&must.StrictPassword{
				WithSpecialChar: true,
				WithUpperCase:   true,
				WithLowerCase:   true,
				WithDigit:       true,
			},
		},
	}
}
