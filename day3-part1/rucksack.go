package main

type Rucksack struct {
	compartment1Items []rune
	compartment2Items []rune
}

func (r Rucksack) CommonItems() []rune {
	var commonItems []rune

	for _, compartment1item := range r.compartment1Items {
		var found bool
		for _, compartment2item := range r.compartment2Items {
			if compartment1item == compartment2item {
				found = true
				break
			}
		}

		if found {
			var alreadyFound bool
			for _, commonItem := range commonItems {
				if commonItem == compartment1item {
					alreadyFound = true
					break
				}
			}
			if !alreadyFound {
				commonItems = append(commonItems, compartment1item)
			}
		}
	}

	return commonItems
}

func ItemPriority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r-'a') + 1
	}
	if r >= 'A' && r <= 'Z' {
		return int(r-'A') + 27
	}
	return -1
}
