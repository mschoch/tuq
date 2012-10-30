package main

func AppendStringSliceIfMissing(slice []string, s []string) []string {
	for _, v := range s {
		slice = AppendStringIfMissing(slice, v)
	}
	return slice
}

func AppendStringIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}
