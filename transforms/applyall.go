package transforms

func Apply(text string) string {
	text = ApplyArticles(text)
	text = ApplyCasing(text)
	text = ApplyNumbers(text)
	text = ApplyPunctuation(text)
	text = ApplyQuotes(text)
	return text
}
