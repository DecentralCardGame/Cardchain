package keeper

func UintItemInList(item uint64, list []uint64) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func StringItemInList(item string, list []string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func IndexOfId(cardID uint64, cards []uint64) int {
	if cards == nil {
		return -1
	}
	for i, b := range cards {
		if b == cardID {
			return i
		}
	}
	return -1
}
