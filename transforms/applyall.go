package transforms

func Apply(text string) string {
	text = ApplyNumbers(text)
	text = ApplyCasing(text)
	text = ApplyArticles(text)
	text = ApplyPunctuation(text)
	text = ApplyQuotes(text)
	return text
}
