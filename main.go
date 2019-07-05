package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)
var alphaDigits = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
var alphaUpper = "ABCDEFGHJKLMNPQRSTUVWXYZ"
var alphaLower = "abcdefghijkmnpqrstuvwxyz"

// Returns a random integer 32bit number in the interval [minValue,maxValue].
// Panic if one or more of the params is negative or maxValue < minValue
func RandomInt( minValue int, maxValue int ) int {
	rand.Seed(time.Now().UTC().UnixNano())
	if minValue < 0 || maxValue < 0 || maxValue < minValue {
		panic("Error, invalid arguments in randomInt function.")
	}
	return minValue + rand.Intn(maxValue-minValue+1)
}
// Returns a random integer 64bit signed number in the interval [minValue,maxValue].
// Panic if one or more of the params is negative or maxValue < minValue
func RandomInt64( minValue int64, maxValue int64 ) int64 {
	rand.Seed(time.Now().UTC().UnixNano())
	if minValue < 0 || maxValue < 0 || maxValue < minValue {
		panic("Error, invalid arguments in randomInt function.")
	}
	return minValue + rand.Int63n(maxValue-minValue+1)
}
//Returns a pseudo random digit or letter of the english alphabet
func RandomAlphaDigitByte() byte {
	return alphaDigits[RandomInt(0,len(alphaDigits)-1)]
}

//Returns a pseudo random english string in upper case or digits, it length will be at least 'minRandomLength' and
//less or equal to 'maxRandomLength'
func RandomAlphaDigitString( minRandomLength int, maxRandomLength int ) string {
	len := RandomInt(minRandomLength,maxRandomLength)
	return RandomAlphaDigitStringExactLength(len)
}

//Returns a pseudo random string in upper case letters or digits, it length will be exactly 'length'
func RandomAlphaDigitStringExactLength( length int ) string {

	return RandomStringExactLength(length,alphaDigits)
}

//Returns a pseudo random upper letter of the english alphabet
func RandomAlphaUpperByte() byte {
	return alphaUpper[RandomInt(0,len(alphaUpper)-1)]
}

//Returns a pseudo random english string in upper case, it length will be at least 'minRandomLength' and
//less or equal to 'maxRandomLength'
func RandomEnglishUpperCaseString( minRandomLength int, maxRandomLength int ) string {
	len := RandomInt(minRandomLength,maxRandomLength)
	return RandomEnglishUpperCaseStringExactLength(len)
}

//Returns a pseudo random english string in upper case, it length will be exactly 'length'
func RandomEnglishUpperCaseStringExactLength( length int ) string {

	return RandomStringExactLength(length,alphaUpper)
}

//Returns a pseudo random lower letter of the english alphabet
func RandomAlphaLowerByte() byte {
	return alphaLower[RandomInt(0,len(alphaLower)-1)]
}

//Returns a pseudo random english string in lowerCase, it length will be at least 'minRandomLength' and
//less or equal to 'maxRandomLength'
func RandomEnglishLowerCaseString( minRandomLength int, maxRandomLength int ) string {
	len := RandomInt(minRandomLength,maxRandomLength)
	return RandomEnglishLowerCaseStringExactLength(len)
}

//Returns a pseudo random english string in lowerCase, it length will be exactly 'length'
func RandomEnglishLowerCaseStringExactLength( length int ) string {

	return RandomStringExactLength(length,alphaLower)
}

//Returns a pseudo random string of character from the given alphabet, it length will be exactly 'length'
//Each character in the 'alphabet' string will have the same probability to appear in the resulting
//random string, if all characters in alphabet are the same, those character will have the same probability,
//the more occurrences a character has in alphabet, the more likely appear in the returning string
func RandomStringExactLength( length int , alphabet string) string {
	if len(alphabet) <= 0 {
		panic("Error, alphabet with no positive length")
	}
	var buffer bytes.Buffer
	for i := 0; i < length; i++ {
		buffer.WriteByte(alphabet[RandomInt(0,len(alphabet)-1)])
	}
	return buffer.String()
}

//Returns a map with 'size' different integers as its keys
func RandomIntSet( size int, minValue int , maxValue int ) (error, map[int]bool) {
	if minValue < 0 || maxValue < 0 || maxValue < minValue || maxValue-minValue+1 < size {
		return fmt.Errorf("error, invalid arguments in getRandomIntSet(size = %d, minValue = %d, maxValue = %d)",
			size, minValue, maxValue), nil
	}
	set := make(map[int]bool)
	for {
		set[RandomInt(minValue,maxValue)] = true
		if len(set) == size {
			break
		}
	}
	return nil, set
}
//Returns a map with 'size' different strings as its keys
func RandomInt64Set( size int, minValue int64 , maxValue int64 ) (error, map[int64]bool) {
	if minValue < 0 || maxValue < 0 || maxValue < minValue || maxValue-minValue+1 < int64(size) {
		return fmt.Errorf("error, invalid arguments in getRandomInt64Set(size = %d, minValue = %d, maxValue = %d)",
			size, minValue, maxValue), nil
	}
	set := make(map[int64]bool)
	for {
		set[RandomInt64(minValue,maxValue)] = true
		if len(set) == size {
			break
		}
	}
	return nil, set
}
//Returns a map with 'size' different integers as its keys
func RandomStringSet( size int, minLength int , maxLength int , alphabet string ) (error, map[string]bool) {
	maxDifferentWordsSet := int(math.Pow(float64(len(alphabet)),float64(minLength)))
	if size >= maxDifferentWordsSet || minLength < 0 || maxLength < 0 || maxLength < minLength || len(alphabet) < 2{
		return fmt.Errorf("error, invalid arguments in " +
			"getRandomStringSet(size = %d, minLength = %d, maxLength = %d, alphabet = %s)",
			size, minLength, maxLength,alphabet), nil
	}
	set := make(map[string]bool)
	for {
		set[RandomStringExactLength( RandomInt(minLength,maxLength), alphabet )] = true
		if len(set) == size {
			break
		}
	}
	return nil, set
}
//Return a valid pseudo random email address
func RandomEmail() string{
	alphabet := alphaLower+"0123456789_0123456789.0123456789"
	end := RandomStringExactLength( RandomInt(2,8),alphaLower)+
		"."+RandomStringExactLength(RandomInt(2,8),alphaLower)
	end += RandomStringExactLength(1,alphaLower)
	return RandomStringExactLength(1,alphaLower)+
		RandomStringExactLength(RandomInt(2,8),alphabet)+
		RandomStringExactLength(1,alphaLower)+"@"+
		RandomStringExactLength(1,alphaLower)+end
}

//Return a valid phone number, just a 10 symbols string formed only by digits 0 to 9
func RandomPhoneNumber() string{

	return RandomStringExactLength(RandomInt(10,10),"0123456789")
}
//Simple pseudoRandom Colombian Address Generator
func RandomAddress() string {
	var first = [5]string {"Calle ","Carrera ","Avenida ","Diagonal ","Transversal "}
	answer := first[RandomInt(0,len(first)-1)]
	answer += fmt.Sprintf("%d",RandomInt(0,250))
	if RandomInt(0,10) < 4 {
		answer += string(alphaUpper[RandomInt(0,len(alphaUpper)-1)])
	}
	answer += " # "
	answer += fmt.Sprintf("%d",RandomInt(0,250))
	if RandomInt(0,10) < 4 {
		answer += string(alphaUpper[RandomInt(0,10)])
	}
	answer += " - "
	answer += fmt.Sprintf("%d",RandomInt(0,100))

	return answer
}
func showEssentials(){
	// Create and seed the generator.
	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.
	r := rand.New(rand.NewSource(99))

	// The tabwriter here helps us generate aligned output.
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	show := func(name string, v1, v2, v3 interface{}) {
		fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", name, v1, v2, v3)
	}

	// Float32 and Float64 values are in [0, 1).
	show("Float32", r.Float32(), r.Float32(), r.Float32())
	show("Float64", r.Float64(), r.Float64(), r.Float64())

	// ExpFloat64 values have an average of 1 but decay exponentially.
	show("ExpFloat64", r.ExpFloat64(), r.ExpFloat64(), r.ExpFloat64())

	// NormFloat64 values have an average of 0 and a standard deviation of 1.
	show("NormFloat64", r.NormFloat64(), r.NormFloat64(), r.NormFloat64())

	// Int31, Int63, and Uint32 generate values of the given width.
	// The Int method (not shown) is like either Int31 or Int63
	// depending on the size of 'int'.
	show("Int31", r.Int31(), r.Int31(), r.Int31())
	show("Int63", r.Int63(), r.Int63(), r.Int63())
	show("Uint32", r.Uint32(), r.Uint32(), r.Uint32())

	// Intn, Int31n, and Int63n limit their output to be < n.
	// They do so more carefully than using r.Int()%n.
	show("Intn(10)", r.Intn(10), r.Intn(10), r.Intn(10))
	show("Int31n(10)", r.Int31n(10), r.Int31n(10), r.Int31n(10))
	show("Int63n(10)", r.Int63n(10), r.Int63n(10), r.Int63n(10))

	// Perm generates a random permutation of the numbers [0, n).
	show("Perm", r.Perm(5), r.Perm(5), r.Perm(5))
}
func main() {
	//showEssentials()
	_, m := RandomStringSet(10, 3,10,"abcd")
	fmt.Printf("\n__%v__\n",m)
}
