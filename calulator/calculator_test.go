package calculator

import (
	"testing"
)

type testCase struct {
	input    []CustomerEntry
	expected []TopCustomerFavourite
}

func TestCalculateTopCustomerFavorites(t *testing.T) {
	tests := []testCase{
		{
			input: []CustomerEntry{
				{
					Name:  "Annika",
					Candy: "Geisha",
					Eaten: 100,
				},
				{
					Name:  "Jonas",
					Candy: "Geisha",
					Eaten: 200,
				},
				{
					Name:  "Jonas",
					Candy: "Kexchoklad",
					Eaten: 100,
				},
				{
					Name:  "Aadya",
					Candy: "Notchoklad",
					Eaten: 2,
				},
				{
					Name:  "Jonas",
					Candy: "Notchoklad",
					Eaten: 3,
				},
				{
					Name:  "Jane",
					Candy: "Notchoklad",
					Eaten: 17,
				},
				{
					Name:  "Annika",
					Candy: "Geisha",
					Eaten: 100,
				},
				{
					Name:  "Jonas",
					Candy: "Geisha",
					Eaten: 700,
				},
				{
					Name:  "Jane",
					Candy: "Notchoklad",
					Eaten: 4,
				},
				{
					Name:  "Aadya",
					Candy: "Center",
					Eaten: 7,
				},
				{
					Name:  "Jonas",
					Candy: "Geisha",
					Eaten: 900,
				},
				{
					Name:  "Jane",
					Candy: "Notchoklad",
					Eaten: 1,
				},
				{
					Name:  "Jonas",
					Candy: "Kexchoklad",
					Eaten: 12,
				},
				{
					Name:  "Jonas",
					Candy: "Plopp",
					Eaten: 20,
				},
				{
					Name:  "Jonas",
					Candy: "Center",
					Eaten: 27,
				},
				{
					Name:  "Aadya",
					Candy: "Center",
					Eaten: 2,
				},
				{
					Name:  "Annika",
					Candy: "Center",
					Eaten: 8,
				},
			},
			expected: []TopCustomerFavourite{
				{},
			},
		},
	}

	for _, t := range tests {

	}
}
