package catserver

var bannerContent = map[string]string{
	"title":         "Kitty Litter",
	"bannerImg":     "https://riotfest.org/wp-content/uploads/2017/10/AcT9YIL.jpg",
	"copyrightYear": "2006",
}

var ownerContent = []string{
	"Pavlov",
	"Schrodinger",
	"Foucault",
	"Vel",
	"Calvin",
}

type litter []string

var catsByOwner = map[string]litter{
	"schrodinger": litter{"Leben", "Tot"},
	"pavlov":      litter{"Belle", "Dribbles", "Nibbles"},
	"foucault":    litter{"M. Fang"},
	"vel":         litter{"Opal"},
	"calvin":      litter{"Hobbes"},
}

// exports.outfits = {
//   taco: 100,
//   princess: 75.1,
//   dog: 89.1,
//   gremlin: 73.1,
//   lampshade: 28
// };
