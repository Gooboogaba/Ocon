package mathfuncs

import "math"
import "strconv"




func Modulo(args []string) []string {
	a, _ := strconv.Atoi(args[0][1:])
	b, _ := strconv.Atoi(args[1][1:])
	return []string{"'" + strconv.Itoa(a % b)}
}

//one value funcs
func Sine(args []string) []string {
	a, _ := strconv.Atoi(args[0][1:]) //turn into number
	return []string{"'" + strconv.FormatFloat(math.Sin(float64(a)), 'f', -1, 64)}//add ' to the string of the sin of a
}

func CoSine(args []string) []string {
	a, _ := strconv.Atoi(args[0][1:]) 
	return []string{"'" + strconv.FormatFloat(math.Cos(float64(a)), 'f', -1, 64)}//this prob cuz math uses float64 and not int (my best gess)
}

func Tangent(args []string) []string {
	a, _ := strconv.Atoi(args[0][1:]) 
	return []string{"'" + strconv.FormatFloat(math.Tan(float64(a)), 'f', -1, 64)}
}

func Round(args []string) []string {
	a, _ := strconv.Atoi(args[0][1:])
	return []string{"'" + strconv.FormatFloat(math.Round(float64(a)), 'f', -1, 64)}
}

func AbsoluteValue(args []string) []string {
	a, _ := strconv.Atoi(args[0][1:])
	return []string{"'" + strconv.FormatFloat(math.Abs(float64(a)), 'f', -1, 64)}
}

func Logarithm(args []string) []string {
	a, _ := strconv.Atoi(args[0][1:])
	return []string{"'" + strconv.FormatFloat(math.Log(float64(a)), 'f', -1, 64)}
}