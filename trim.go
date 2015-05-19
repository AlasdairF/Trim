package trim

func Space(word []byte) []byte {
	return SpaceRight(SpaceLeft(word))
}

func SpaceLeft(word []byte) []byte {
	for i, b := range word {
		if b > 32 {
			return word[i:]
		}
	}
	return []byte{}
}

func SpaceRight(word []byte) []byte {
	for i:=len(word)-1; i>=0; i-- {
		if word[i] > 32 {
			return word[0:i+1]
		}
	}
	return []byte{}
}

func SpacePunc(word []byte) []byte {
	return SpacePuncRight(SpacePuncLeft(word))
}

func SpacePuncLeft(word []byte) []byte {
	for i, b := range word {
		if b > 47 {
			return word[i:]
		}
	}
	return []byte{}
}

func SpacePuncRight(word []byte) []byte {
	for i:=len(word)-1; i>=0; i-- {
		if word[i] > 47 {
			return word[0:i+1]
		}
	}
	return []byte{}
}
