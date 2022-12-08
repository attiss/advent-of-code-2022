package main

type Group struct {
	rucksack1Items []rune
	rucksack2Items []rune
	rucksack3Items []rune
}

func (g Group) GetBadge() rune {
	for _, rucksack1Item := range g.rucksack1Items {
		var found bool

		for _, rucksack2Item := range g.rucksack2Items {
			if rucksack1Item == rucksack2Item {
				found = true
				break
			}
		}

		if !found {
			continue
		}

		for _, rucksack3Item := range g.rucksack3Items {
			if rucksack3Item == rucksack1Item {
				return rucksack1Item
			}
		}
	}

	return -1
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
