package main

import (
	"errors"
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("HELLO WORD")

	result := calculation.Add(3, 5)
	result2 := calculation.Multiply(5, 3)

	fmt.Println(result)
	fmt.Println(result2)

	var name string = "Dafa"
	var age int
	age = 12
	fmt.Println(name)
	fmt.Println(age)

	score := 50
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score < 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	fmt.Println(grade)

	number := 1
	switch number {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("DEFAULT")
	}

	// ====================================================================
	for i := 1; i <= 20; i++ {
		fmt.Println("Learn For loop golang index: ", i)
	}

	// ====================================================================
	i := 1
	for i <= 10 {
		fmt.Println("Learn For loop golang-2 index:", i)
		i++
	}

	// ====================================================================
	word := "golang the best language"
	// for index, letter := range word {
	// 	fmt.Println("Index: ", index, " Letter: ", string(letter))
	// }
	// for _, letter := range word {
	// 	fmt.Println("Letter: ", string(letter))
	// }

	for index, letter := range word {
		if index%2 == 0 {
			fmt.Println("Index: ", index, " Letter: ", string(letter))
		}
	}

	// print AIUEO
	for _, letter := range word {
		letterString := string(letter)
		// if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "o" {
		// 	fmt.Println("Letter with AIOUE result: ", letterString)
		// }
		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("letter with AIUEO and result:", letterString)
		}
	}

	//======================= ARRAY ========================
	/** defined another array
		var languages [5]string
		languages[0] = "Go"
		languages[1] = "Javascript"
		languages[2] = "PHP"
		languages[3] = "Python"
		languages[4] = "C++"
	**/

	//languages := [5]string{"Go", "Javascript", "PHP", "Python", "C++"}
	languages := [...]string{"Go", "Javascript", "PHP", "Python", "C++", "Ruby"}

	for index, language := range languages {
		fmt.Println("Index:", index, "Print on looping array:", language)
	}

	fmt.Println(languages)
	fmt.Println("Length of languages:", len((languages)))

	var gamingConsoles []string
	gamingConsoles = append(gamingConsoles, "PlayStation4")    // append like .push on javascript
	gamingConsoles = append(gamingConsoles, "Nintendo Switch") // append like .push on javascript
	gamingConsoles = append(gamingConsoles, "XBox")            // append like .push on javascript

	for _, console := range gamingConsoles {
		fmt.Println("Console Gaming", console)
	}

	/**
		var myMap map[string]int
		myMap = map[string]int{}

		myMap["Javascript"] = 9
		myMap["PHP"] = 10
		myMap["Python"] = 8
		myMap["Python"] = 10

		fmt.Println(myMap["Python"])

	**/

	myMap := map[string]string{
		"Javascript": "Wow Javascript",
		"PHP":        "Wow PHP",
		"Python":     "Wow Python",
		"Laravel":    "Is the Framework of PHP",
	}
	fmt.Println(myMap)

	for key, value := range myMap {
		fmt.Println("key:", key, "value:", value)
	}
	fmt.Println("===========")
	delete(myMap, "Laravel") // Remove by key
	fmt.Println("Result after deleted: ", myMap)

	value, isAvaliable := myMap["C++"] // Check by key
	fmt.Println(isAvaliable)
	fmt.Println(value)

	students := []map[string]string{
		{"name": "Budi", "score": "A"},
		{"name": "Jordan", "score": "B"},
		{"name": "Willy", "score": "E"},
	}

	for _, student := range students {
		fmt.Println(student["name"], "=====", student["score"])
	}

	scores := []int{100, 80, 75, 92, 70, 93, 88, 67} // lesson
	scoreLength := len(scores)
	var goodScores []int
	var totalResult int

	for _, score := range scores {
		totalResult = totalResult + score
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}
	fmt.Println(float64(len(scores)))
	sumOfScore := float64(totalResult) / float64(scoreLength)

	fmt.Println("Sum Of Score", sumOfScore)
	fmt.Println("GoodScores:", goodScores)
	setence := printResult("Hello.")
	fmt.Println(setence)

	luas, keliling := calcuate(4, 5)
	fmt.Println("Luas:", luas)
	fmt.Println("Keliling:", keliling)

	sumTest := sumTest(scores)

	fmt.Println(sumTest)

	//result, err := calcuateTest(2,4,"+")
	//result, err := calcuateTest(2, 4, "-")
	//result, err := calcuateTest(2, 4, "*")
	//result, err := calcuateTest(2, 4, "/")
	result, err := calcuateTest(2, 4, "=")
	if err != nil {
		fmt.Println("Terjadi Kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

func printResult(value string) string {
	result := value + " from function"
	return result
}

// function with predefined return value
func calcuate(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	return
}

func sumTest(numbers []int) (total int) {
	for _, number := range numbers {
		total = total + number
	}
	return total
}

func calcuateTest(numberOne int, numberTwo int, operator string) (result int, errResult error) {
	switch operator {
	case "+":
		result = numberOne + numberTwo
	case "-":
		result = numberOne - numberTwo
	case "*":
		result = numberOne * numberTwo
	case "/":
		result = numberOne / numberTwo
	default:
		errResult = errors.New("Uknown Operation value")
	}
	return result, errResult
}

/**
func calcuate(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}
**/

// go build; go mod init pertama; ./pertama;
