package pkg

func fn(x uint32) {
	if x >= 0 { // MATCH /unsigned values are always >= 0/
	}
	if x < 0 { // MATCH /unsigned values are never < 0/
	}
}
