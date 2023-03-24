package tasklist

type task struct {
	description, content string
}

func loadTasks() []task {
	allChoices := []task{
		task{description: "Raspberry Pi’s", content: "I have ’em all over my house"},
		task{description: "Nutella", content: "It's good on toast"},
		task{description: "Bitter melon", content: "It cools you down"},
		task{description: "Nice socks", content: "And by that I mean socks without holes"},
		task{description: "Eight hours of sleep", content: "I had this once"},
		task{description: "Cats", content: "Usually"},
		task{description: "Plantasia, the album", content: "My plants love it too"},
		task{description: "Pour over coffee", content: "It takes forever to make though"},
		task{description: "VR", content: "Virtual reality...what is there to say?"},
		task{description: "Noguchi Lamps", content: "Such pleasing organic forms"},
		task{description: "Linux", content: "Pretty much the best OS"},
		//task{description: "Business school", content: "Just kidding"},
		//task{description: "Pottery", content: "Wet clay is a great feeling"},
		//task{description: "Shampoo", content: "Nothing like clean hair"},
		//task{description: "Table tennis", content: "It’s surprisingly exhausting"},
		//task{description: "Milk crates", content: "Great for packing in your extra stuff"},
		//task{description: "Afternoon tea", content: "Especially the tea sandwich part"},
		//task{description: "Stickers", content: "The thicker the vinyl the better"},
		//task{description: "20° Weather", content: "Celsius, not Fahrenheit"},
		//task{description: "Warm light", content: "Like around 2700 Kelvin"},
		//task{description: "The vernal equinox", content: "The autumnal equinox is pretty good too"},
		//task{description: "Gaffer’s tape", content: "Basically sticky fabric"},
		//task{description: "Terrycloth", content: "In other words, towel fabric"},
	}
	return allChoices
}
