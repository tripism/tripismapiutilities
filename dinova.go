package tripismapiutilities

import (
	"strings"

	"github.com/tripism/api/types"
)

func CalcDinovaAPIRestaurantSHA512(restaurant *types.DinovaAPIRestaurant) {
	if restaurant == nil {
		return
	}
	restaurant.SHA512 = CalculateSHA512(strings.Join([]string{
		restaurant.Address1,
		restaurant.Address2,
		restaurant.ChainID,
		restaurant.ChainLocations,
		restaurant.City,
		restaurant.Country,
		restaurant.Cuisines,
		restaurant.Image,
		restaurant.Latitude,
		restaurant.Longitude,
		restaurant.Phone,
		restaurant.PrimaryCuisineType,
		restaurant.ReservationsLink,
		restaurant.RestaurantName,
		restaurant.State,
		// restaurant.Styles, // not included into live Dinova API
		restaurant.WebsiteLink,
		restaurant.Zip,
	}, "-"))

	return
}
