// Word Ladder
// Google interview hard question

package main

import "fmt"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	type pair struct {
		word  string
		steps int
	}

	queue := []pair{{word: beginWord, steps: 1}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.word == endWord {
			return current.steps
		}

		// try all transformations
		word := []rune(current.word)
		for i := 0; i < len(word); i++ {
			original := word[i]
			for ch := 'a'; ch <= 'z'; ch++ {
				if ch == original {
					continue
				}
				word[i] = ch
				newWord := string(word)
				if wordSet[newWord] {
					queue = append(queue, pair{newWord, current.steps + 1})
					delete(wordSet, newWord) // mark visited
				}
			}
			word[i] = original
		}
	}
	return 0
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	fmt.Println("Shortest ladder length : ", ladderLength(beginWord, endWord, wordList))
}
