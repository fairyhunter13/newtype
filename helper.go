package newtype

func isNullJSON(in []byte) bool {
	return len(in) == 4 && (in[0] == 'n' && in[1] == 'u' && in[2] == 'l' && in[3] == 'l')
}

func isStringJSON(in []byte) bool {
	return len(in) >= 2 && in[0] == '"' && in[len(in)-1] == '"'
}

func removeQuotesJSON(in []byte) []byte {
	if isStringJSON(in) {
		in = in[1 : len(in)-1]
	}
	return in
}
