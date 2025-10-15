/*
Copyright © 2025 Shreya Mandi
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*
	TODO

[/] addition, sub, multiplication and division of 4,5, and 6 digits
[/] currency and metric system quick conversion problems
[/] timer feature
[/] recording progress meter
[ ] progress aggregation logic
[ ] cli framework
*/

type Problem struct {
	desc          string
	sumType       string
	numbers       []float32
	correctAnswer float32
	userAnswer    float32
	status        bool
	timeTaken     time.Duration
}

func (problem *Problem) generateNumbers(d string) {
	switch problem.sumType {
	case "inr_to_usd":
		problem.numbers = append(problem.numbers, float32(rand.Intn(1000)+100))
		return
	case "usd_to_inr":
		problem.numbers = append(problem.numbers, float32(rand.Intn(1000)+100))
		return
	case "kms_to_miles":
		problem.numbers = append(problem.numbers, float32(rand.Intn(100)+10))
		return
	case "miles_to_kms":
		problem.numbers = append(problem.numbers, float32(rand.Intn(100)+10))
		return
	case "cms_to_inches":
		problem.numbers = append(problem.numbers, float32(rand.Intn(100)+10))
		return
	case "inches_to_cms":
		problem.numbers = append(problem.numbers, float32(rand.Intn(100)+10))
		return
	default:
		diff_level_map := map[string]int{
			"LOW":    10000,
			"MEDIUM": 100000,
			"HIGH":   1000000,
		}
		greater := rand.Intn(diff_level_map[d]) + diff_level_map[d]/10

		smaller := rand.Intn(greater) + diff_level_map[d]/10
		problem.numbers = append(problem.numbers, float32(smaller))
		problem.numbers = append(problem.numbers, float32(greater))

	}
	return
}

func (problem *Problem) calculateAnswer() {
	switch problem.sumType {

	case "add":
		problem.correctAnswer = problem.numbers[1] + problem.numbers[0]
		break
	case "sub":
		problem.correctAnswer = problem.numbers[1] - problem.numbers[0]
		break
	case "mul":
		problem.correctAnswer = problem.numbers[1] * problem.numbers[0]
		break
	case "div":
		problem.correctAnswer = problem.numbers[1] / problem.numbers[0]
		break
	case "inr_to_usd":
		problem.correctAnswer = problem.numbers[0] / 83
		problem.desc += " 1 USD = 83 rupees"
		break
	case "usd_to_inr":
		problem.correctAnswer = problem.numbers[0] * 83
		problem.desc += " 1 USD = 83 rupees"
		break
	case "kms_to_miles":
		problem.correctAnswer = problem.numbers[0] * 0.6
		problem.desc += " 1 km = 0.6 miles"
		break
	case "miles_to_kms":
		problem.correctAnswer = problem.numbers[0] * 1.6
		problem.desc += " 1 mile = 1.6 kms"
		break
	case "cms_to_inches":
		problem.correctAnswer = problem.numbers[0] * 0.4
		problem.desc += " 1 cm = 0.4 in"
		break
	case "inches_to_cms":
		problem.correctAnswer = problem.numbers[0] * 2.5
		problem.desc += " 1 in = 2.5 cms"
		break
	default:
		break
	}
}

func (problem *Problem) checkAnswer(userAnswer float32) {
	problem.userAnswer = userAnswer
	if problem.userAnswer == problem.correctAnswer {
		problem.status = true
	} else {
		problem.status = false
	}
}

func (problem *Problem) print(i int) {
	fmt.Println("")
	fmt.Println("Problem Data", i+1)
	var symbol string
	switch problem.sumType {
	case "sub":
		symbol = "-"
		break
	case "mul":
		symbol = "*"
		break
	case "div":
		symbol = "/"
		break
	default:
		break
	}

	if problem.sumType == "add" || problem.sumType == "sub" || problem.sumType == "mul" || problem.sumType == "div" {
		fmt.Println(problem.numbers[1], symbol, problem.numbers[0])
	} else {
		fmt.Println(problem.desc)
		fmt.Println(problem.numbers[0])
	}
}

func (problem *Problem) store(i int, date string) {
	dirName := "mentalist_progress"
	err := os.Mkdir(dirName, 0755) // 0755 sets permissions (read, write, execute for owner; read, execute for group/others)
	//if err != nil {
	//	fmt.Printf("Error creating directory: %v\n", err)
	//	return
	//}

	problemDict := make(map[string]any)
	problemDict["desc"] = problem.desc
	problemDict["sumType"] = problem.sumType
	problemDict["numbers"] = problem.numbers
	problemDict["correctAnswer"] = problem.correctAnswer
	problemDict["userAnswer"] = problem.userAnswer
	problemDict["status"] = problem.status
	problemDict["timeTaken"] = problem.timeTaken

	jsonData, err := json.Marshal(problemDict)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	writeData := []byte(string(jsonData))
	filename := "mentalist_progress/" + "problem" + strconv.Itoa(i) + date + ".json"
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	if _, err := f.Write([]byte(writeData)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		fmt.Println("Error writing problem.json:", err)
	}
}

type ProblemSet struct {
	problems   []Problem
	difficulty string
	date       string // todo
}

func (set *ProblemSet) generate() {
	// generate addition problem
	addProblem := Problem{
		desc:    "mEnTaLiSt ADDITION",
		sumType: "add",
	}
	set.problems = append(set.problems, addProblem)

	// generate subtraction problem
	subProblem := Problem{
		desc:    "mEnTaLiSt SUBTRACT",
		sumType: "sub",
	}
	set.problems = append(set.problems, subProblem)

	// generate multiplication problem
	mulProblem := Problem{
		desc:    "mEnTaLiSt MULTIPLY",
		sumType: "mul",
	}
	set.problems = append(set.problems, mulProblem)

	// generate division problem
	divProblem := Problem{
		desc:    "mEnTaLiSt DIVISION",
		sumType: "div",
	}
	set.problems = append(set.problems, divProblem)

	// generate currency conversions
	currencyProblem := Problem{
		desc:    "mEnTaLiSt INR to USD",
		sumType: "inr_to_usd",
	}
	set.problems = append(set.problems, currencyProblem)

	currencyProblemReverse := Problem{
		desc:    "mEnTaLiSt USD to INR",
		sumType: "usd_to_inr",
	}
	set.problems = append(set.problems, currencyProblemReverse)

	// generate metric conversions
	metricConversion := Problem{
		desc:    "mEnTaLiSt kms to miles",
		sumType: "kms_to_miles",
	}
	set.problems = append(set.problems, metricConversion)
	metricConversionReverse := Problem{
		desc:    "mEnTaLiSt miles to kms",
		sumType: "miles_to_kms",
	}
	set.problems = append(set.problems, metricConversionReverse)

	metricConversion2 := Problem{
		desc:    "mEnTaLiSt cms to inches",
		sumType: "cms_to_inches",
	}
	set.problems = append(set.problems, metricConversion2)
	metricConversion2Reverse := Problem{
		desc:    "mEnTaLiSt inches to cms",
		sumType: "inches_to_cms",
	}
	set.problems = append(set.problems, metricConversion2Reverse)
}

func (set *ProblemSet) print(d string) {
	set.difficulty = d

	fmt.Println("Difficulty: ", set.difficulty)
	fmt.Println("Date: ", set.date)

	for i, problem := range set.problems {
		problem.generateNumbers(d)
		problem.calculateAnswer()
		problem.print(i)
		fmt.Println("Enter your answer:")
		t1 := time.Now()
		var userInput float32
		_, err := fmt.Scanln(&userInput)
		if err != nil {
			fmt.Println("Error with user input:", err)
		}
		problem.checkAnswer(userInput)
		t2 := time.Now()
		elapsed := t2.Sub(t1)
		if problem.status {
			fmt.Println("Your answer is correct!")
			fmt.Println("It took you: ", elapsed)
		} else {
			fmt.Println("Your answer is incorrect!")
			fmt.Println("The correct answer is:", problem.correctAnswer)
		}
		fmt.Println("")
		problem.timeTaken = elapsed

		problem.store(i, set.date)
	}
}

func run() {
	fmt.Println("Hi Shreya! This is mEnTaLiSt!")
	var d string
	fmt.Println("Select difficulty type: LOW, MEDIUM, HIGH")
	_, err := fmt.Scanln(&d)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	pSet := ProblemSet{
		date: time.Now().Format("2006-01-02"),
	}

	pSet.generate()
	pSet.print(d)
}

// mentalistCmd represents the mentalist command
var mentalistCmd = &cobra.Command{
	Use:   "mentalist",
	Short: "Train your mental aptitude everyday!",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := mentalistCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.problemGenerator.go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	mentalistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
