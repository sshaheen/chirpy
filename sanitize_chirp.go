package main

import "strings"

func sanitizeChirp(s string) string {
	badWords := make(map[string]struct{})

	badWords["kerfuffle"] = struct{}{}
	badWords["sharbert"] = struct{}{}
	badWords["fornax"] = struct{}{}

	words := strings.Split(s, " ")
	for idx, word := range words {
		caseIns := strings.ToLower(word)
		if _, ok := badWords[caseIns]; ok {
			words[idx] = "****"
		}
	}

	return strings.Join(words, " ")
}
