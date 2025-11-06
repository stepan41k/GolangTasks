package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

// Alphabet представляет алфавит для анализа.
type Alphabet struct {
	Name    string
	Letters []rune
	Map     map[rune]int
}

// NewAlphabet создает и инициализирует новый алфавит.
func NewAlphabet(letters []rune, name string) *Alphabet {
	m := make(map[rune]int)
	for i, r := range letters {
		m[r] = i
	}
	return &Alphabet{
		Name:    name,
		Letters: letters,
		Map:     m,
	}
}

// AnalysisResult хранит полные результаты анализа текста.
type AnalysisResult struct {
	TextName                 string
	Alphabet                 *Alphabet
	TotalChars               int
	TotalBigrams             int
	CharCounts               []int
	BigramCounts             [][]int
	CharProbabilities        []float64
	BigramProbabilities      [][]float64
	ConditionalProbabilities [][]float64
	EntropyH_A               float64
	MarkovEntropyH_A_A       float64
}

var (
	RussianAlphabet = NewAlphabet([]rune("абвгдежзийклмнопрстуфхцчшщъыьэюя"), "Русский")
	LatinAlphabet   = NewAlphabet([]rune("abcdefghijklmnopqrstuvwxyz"), "Латиница")
)

// normalizeText очищает и нормализует текст в соответствии с заданным алфавитом.
func normalizeText(text string, alphabet *Alphabet) string {
	var result strings.Builder
	lowerText := strings.ToLower(text)
	for _, r := range lowerText {
		if _, ok := alphabet.Map[r]; ok {
			result.WriteRune(r)
		}
	}
	fmt.Println(result.String())
	return result.String()
}

// analyzeText выполняет полный частотный анализ текста.
func analyzeText(text, textName string, alphabet *Alphabet) *AnalysisResult {
	normalizedText := normalizeText(text, alphabet)
	n := len(alphabet.Letters)
	runes := []rune(normalizedText)
	textLen := len(runes)

	res := &AnalysisResult{
		TextName:                 textName,
		Alphabet:                 alphabet,
		TotalChars:               textLen,
		CharCounts:               make([]int, n),
		BigramCounts:             make([][]int, n),
		CharProbabilities:        make([]float64, n),
		BigramProbabilities:      make([][]float64, n),
		ConditionalProbabilities: make([][]float64, n),
	}

	for i := 0; i < n; i++ {
		res.BigramCounts[i] = make([]int, n)
		res.BigramProbabilities[i] = make([]float64, n)
		res.ConditionalProbabilities[i] = make([]float64, n)
	}

	// 1. Подсчет частот одиночных символов
	for _, r := range runes {
		if idx, ok := alphabet.Map[r]; ok {
			res.CharCounts[idx]++
		}
	}

	// 2. Подсчет частот биграмм
	// КОРРЕКТНЫЙ ЦИКЛ: используется `i < textLen-1`
	if textLen > 1 {
		for i := 0; i < textLen-1; i++ {
			idx1, ok1 := alphabet.Map[runes[i]]
			idx2, ok2 := alphabet.Map[runes[i+1]]
			if ok1 && ok2 {
				res.BigramCounts[idx1][idx2]++
				res.TotalBigrams++
			}
		}
	}

	// 3. Расчет вероятностей
	if res.TotalChars > 0 {
		for i := 0; i < n; i++ {
			res.CharProbabilities[i] = float64(res.CharCounts[i]) / float64(res.TotalChars)
		}
	}

	if res.TotalBigrams > 0 {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				res.BigramProbabilities[i][j] = float64(res.BigramCounts[i][j]) / float64(res.TotalBigrams)
				if res.CharCounts[i] > 0 {
					res.ConditionalProbabilities[i][j] = float64(res.BigramCounts[i][j]) / float64(res.CharCounts[i])
				}
			}
		}
	}

	// 4. Расчет энтропии
	res.EntropyH_A = calculateEntropyH_A(res.CharProbabilities)
	res.MarkovEntropyH_A_A = calculateMarkovEntropyH_A_A(res.CharProbabilities, res.ConditionalProbabilities)

	return res
}

// calculateEntropyH_A вычисляет энтропию источника H(A).
func calculateEntropyH_A(probs []float64) float64 {
	var entropy float64
	for _, p := range probs {
		if p > 0 {
			entropy -= p * math.Log2(p)
		}
	}
	return entropy
}

// calculateMarkovEntropyH_A_A вычисляет марковскую энтропию H(A|A).
func calculateMarkovEntropyH_A_A(charProbs []float64, condProbs [][]float64) float64 {
	var entropy float64
	n := len(charProbs)
	for i := 0; i < n; i++ {
		if charProbs[i] > 0 {
			var innerSum float64
			for j := 0; j < n; j++ {
				if condProbs[i][j] > 0 {
					innerSum -= condProbs[i][j] * math.Log2(condProbs[i][j])
				}
			}
			entropy += charProbs[i] * innerSum
		}
	}
	return entropy
}

// calculateConditionalEntropyH_A_B вычисляет условную энтропию H(A|B).
func calculateConditionalEntropyH_A_B(textA, textB string, alphabet *Alphabet) float64 {
	normA := normalizeText(textA, alphabet)
	normB := normalizeText(textB, alphabet)
	n := len(alphabet.Letters)

	minLen := len(normA)
	if len(normB) < minLen {
		minLen = len(normB)
	}

	if minLen == 0 {
		return 0
	}

	bCounts := make([]int, n)
	jointCounts := make([][]int, n)
	for i := 0; i < n; i++ {
		jointCounts[i] = make([]int, n)
	}

	runesA := []rune(normA)
	runesB := []rune(normB)

	// КОРРЕКТНЫЙ ЦИКЛ: используется `i < minLen`
	for i := 0; i < minLen; i++ {
		idxA, okA := alphabet.Map[runesA[i]]
		idxB, okB := alphabet.Map[runesB[i]]
		if okA && okB {
			jointCounts[idxA][idxB]++
			bCounts[idxB]++
		}
	}

	var entropy float64
	for j := 0; j < n; j++ {
		if bCounts[j] > 0 {
			probB := float64(bCounts[j]) / float64(minLen)
			var innerSum float64
			for i := 0; i < n; i++ {
				if jointCounts[i][j] > 0 {
					condProb := float64(jointCounts[i][j]) / float64(bCounts[j])
					innerSum -= condProb * math.Log2(condProb)
				}
			}
			entropy += probB * innerSum
		}
	}

	return entropy
}

// calculateJointEntropyH_A_B вычисляет совместную энтропию H(A,B).
func calculateJointEntropyH_A_B(textA, textB string, alphabet *Alphabet) float64 {
	normA := normalizeText(textA, alphabet)
	normB := normalizeText(textB, alphabet)
	n := len(alphabet.Letters)

	minLen := len(normA)
	if len(normB) < minLen {
		minLen = len(normB)
	}

	fmt.Println(minLen)

	if minLen == 0 {
		return 0
	}

	jointCounts := make([][]int, n)
	for i := 0; i < n; i++ {
		jointCounts[i] = make([]int, n)
	}

	runesA := []rune(normA)
	runesB := []rune(normB)

	// КОРРЕКТНЫЙ ЦИКЛ: используется `i < minLen`
	for i := 0; i < minLen; i++ {
		idxA, okA := alphabet.Map[runesA[i]]
		idxB, okB := alphabet.Map[runesB[i]]
		if okA && okB {
			jointCounts[idxA][idxB]++
		}
	}

	var entropy float64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if jointCounts[i][j] > 0 {
				p := float64(jointCounts[i][j]) / float64(minLen)
				entropy -= p * math.Log2(p)
			}
		}
	}

	return entropy
}

// printResults выводит результаты анализа в консоль.
func printResults(res *AnalysisResult) {
	fmt.Printf("--- Результаты анализа для текста: %s ---\n", res.TextName)
	fmt.Printf("Алфавит: %s (%d символов)\n", res.Alphabet.Name, len(res.Alphabet.Letters))
	fmt.Printf("Всего символов (после нормализации): %d\n", res.TotalChars)
	fmt.Printf("Всего биграмм: %d\n\n", res.TotalBigrams)

	fmt.Println("Гистограмма и вероятности одиночных символов:")
	type charStat struct {
		r     rune
		count int
		prob  float64
	}
	var stats []charStat
	for i, r := range res.Alphabet.Letters {
		stats = append(stats, charStat{r, res.CharCounts[i], res.CharProbabilities[i]})
	}
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].count > stats[j].count
	})

	maxCount := 0
	if len(stats) > 0 && stats[0].count > 0 {
		maxCount = stats[0].count
	}

	for _, stat := range stats {
		if stat.count > 0 {
			bar := ""
			if maxCount > 0 {
				barLength := int(float64(stat.count) / float64(maxCount) * 50.0)
				bar = strings.Repeat("█", barLength)
			}
			fmt.Printf("'%c': %5d (P=%.4f) | %s\n", stat.r, stat.count, stat.prob, bar)
		}
	}
	fmt.Println("\n(Таблицы вероятностей биграмм и условных вероятностей опущены для краткости вывода)")

	fmt.Println("\nЧисленные значения энтропии:")
	fmt.Printf("  - Энтропия источника H(A)                  = %.4f\n", res.EntropyH_A)
	fmt.Printf("  - Марковская энтропия 1-го порядка H(A|A) = %.4f\n", res.MarkovEntropyH_A_A)
}

func main() {
	// --- Анализ русского текста A ---
	russianTextABytes, err := os.ReadFile("./cmd/russian_text.txt")
	if err != nil {
		log.Fatalf("Ошибка чтения файла russian_text.txt: %v", err)
	}
	russianTextA := string(russianTextABytes)
	rusResultA := analyzeText(russianTextA, "./cmd/russian_text.txt", RussianAlphabet)
	printResults(rusResultA)

	// --- Анализ латинского текста ---
	fmt.Printf("\n=========================================================\n")
	latinTextBytes, err := os.ReadFile("./cmd/latin_text.txt")
	if err != nil {
		log.Fatalf("Ошибка чтения файла latin_text.txt: %v", err)
	}
	latinText := string(latinTextBytes)
	latResult := analyzeText(latinText, "./cmd/latin_text.txt", LatinAlphabet)
	printResults(latResult)

	// --- Совместный анализ двух русских текстов ---
	fmt.Printf("\n=========================================================\n")
	russianTextBBytes, err := os.ReadFile("./cmd/russian_text_B.txt")
	if err != nil {
		// Программа завершится, если второго файла нет, но не упадет.
		log.Printf("Файл 'russian_text_B.txt' не найден, совместный анализ пропущен: %v", err)
		return
	}
	russianTextB := string(russianTextBBytes)

	fmt.Printf("--- Совместный анализ для двух русских текстов ---\n")
	fmt.Printf("  Текст A: %s\n", rusResultA.TextName)
	fmt.Printf("  Текст B: %s\n", "./cmd/russian_text_B.txt")

	// Условная энтропия H(A|B)
	entropyH_A_B := calculateConditionalEntropyH_A_B(russianTextA, russianTextB, RussianAlphabet)
	fmt.Printf("  - Условная энтропия H(A|B) = %.4f\n", entropyH_A_B)

	// Совместная энтропия H(A,B)
	entropyH_AB := calculateJointEntropyH_A_B(russianTextA, russianTextB, RussianAlphabet)
	fmt.Printf("  - Совместная энтропия H(A,B) = %.4f\n", entropyH_AB)
}