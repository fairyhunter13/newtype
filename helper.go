package newtype

func IsNullJSON(in []byte) bool {
	return len(in) == 4 && (in[0] == 'n' && in[1] == 'u' && in[2] == 'l' && in[3] == 'l')
}

func IsStringJSON(in []byte) bool {
	return len(in) >= 2 && in[0] == '"' && in[len(in)-1] == '"'
}

func RemoveQuotesJSON(in []byte) []byte {
	if IsStringJSON(in) {
		in = in[1 : len(in)-1]
	}
	return in
}
