package newtype

func isNull(in []byte) bool {
	return len(in) == 4 && (in[0] == 'n' && in[1] == 'u' && in[2] == 'l' && in[3] == 'l')
}

func isString(in []byte) bool {
	return len(in) >= 2 && in[0] == '"' && in[len(in)-1] == '"'
}
