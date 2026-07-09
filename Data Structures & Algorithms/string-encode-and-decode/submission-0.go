type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder

	for _, str := range strs {
		builder.WriteString(strconv.Itoa(len(str)))
		builder.WriteString("#")
		builder.WriteString(str)
	}

	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	result := []string{}

	i := 0
	for i < len(encoded) {
		j := i

		for encoded[j] != '#' {
			j++
		}

		length, _ := strconv.Atoi(encoded[i:j])

		start := j + 1
		end := start + length

		result = append(result, encoded[start:end])

		i = end
	}

	return result
}
