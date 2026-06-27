type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var encoded strings.Builder
	for _, str := range strs {
		
		s := fmt.Sprintf("%d#%s", len(str), str)
		encoded.WriteString(s)
	}
	return encoded.String()
}

func (s *Solution) Decode(encoded string) []string {
	var decoded []string
	for i :=0 ; i < len(encoded);{
		j := i
		for string(encoded[j]) != "#"{
			j++
		}
		v, err := strconv.Atoi(encoded[i:j])
		if err == nil {
			decoded = append(decoded, encoded[j+1:j+1+v])
			i = j + 1 + v
		}else{
			i++
		}
	}
	return decoded
}
