// Package calculator takes as input a list of customer entries and returns a list of favourite snacks for top customers.
package calculator

import ()

type TopCustomerFavourite struct {
	Name           string `json:"name"`
	FavouriteSnack string `json:"favouriteSnack"`
	TotalSnacks    int    `json:"totalSnacks"`
}

type CustomerEntry struct {
	Name  string `json:"name"`
	Candy string `json:"candy"`
	Eaten int    `json:"eaten"`
}

type customerData map[string]int

func CalculateTopCustomerFavorites(entries []CustomerEntry) []TopCustomerFavourite {

	// Use hash map to store intermediate data
	m := make(map[string]customerData)
	for _, entry := range entries {

		// Create a customer data entry for this customer if there isn't one already
		if _, ok := m[entry.Name]; !ok {
			m[entry.Name] = customerData{}
		}

		// Create entry for candy  type if there isn't one already
		if _, ok := m[entry.Name][entry.Candy]; !ok {
			m[entry.Name][entry.Candy] = 0
		}

		// Add amount eaten to entry for this candy type
		m[entry.Name][entry.Candy] += entry.Eaten
	}

	return createTopCustomerFavouritesFromHashmap(m)
}

func createTopCustomerFavouritesFromHashmap(m map[string]customerData) []TopCustomerFavourite {
	customerFavourites := make([]TopCustomerFavourite, 0, len(m))
	for name, data := range m {
		favouriteCandyName, totalSnacks := findFavouriteCandyAndTotalSnacks(data)
		customerFavourites = append(customerFavourites, TopCustomerFavourite{
			Name:           name,
			FavouriteSnack: favouriteCandyName,
			TotalSnacks:    totalSnacks,
		})
	}

	return customerFavourites
}

func findFavouriteCandyAndTotalSnacks(data customerData) (string, int) {
	maxEaten := 0
	favouriteCandy := ""
	total := 0
	for candyName, amountEaten := range data {
		if amountEaten > maxEaten {
			favouriteCandy = candyName
		}
		total += amountEaten
	}
	return favouriteCandy, total
}
