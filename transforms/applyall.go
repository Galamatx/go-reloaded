package transforms

func Apply(text string) string {

	for {
		oldText := text
		text = ApplyNumbers(text)
		text = ApplyCasing(text)
		if text == oldText {
			break
		}
	}
	text = ApplyArticles(text)
	text = ApplyPunctuation(text)
	text = ApplyQuotes(text)
	return text

}
