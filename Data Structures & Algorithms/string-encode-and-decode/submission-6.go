type Solution struct{}

func (s *Solution) Encode(strs []string)string{
	var encoded strings.Builder
	for _, str:= range strs{
		encoded.WriteString(fmt.Sprintf("%d#%s", len(str),str))
	}
	return encoded.String()
}
func (s *Solution) Decode(encoded string) []string {
	decoded := make([]string, 0)
	l := make([]byte,0)
	for i := 0; i < len(encoded);{
	
		if encoded[i] != '#'{
			l = append(l, encoded[i])
			i++
		}else{
			n, _ := strconv.Atoi(string(l))
			decoded = append(decoded, encoded[i+1:i+1+n])
			i = i + n + 1
			l = make([]byte,0)
			
		}
	}
	return decoded
	
}
