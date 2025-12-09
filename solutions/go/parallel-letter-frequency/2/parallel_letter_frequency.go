package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	channel := make(chan FreqMap, len(texts))
	frequencies := FreqMap{}
	for _, text := range texts {
		go Freq(text, channel)
	}

	for i := 0; i < len(texts); i++ {
		c := <-channel

		for k, v := range c {
			frequencies[k] += v
		}

	}
	return frequencies
}

func Freq(text string, c chan FreqMap) {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	c <- frequencies
}
