package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)
const maxTestCasesSets = 50
func TestRandomAlphaDigitString(t *testing.T){
	for i:=0 ; i < 100; i++ {
		test := RandomAlphaDigitString(10,20)
		for _, char := range test {
			x := int(char)
			ok := ( int('a') <= x && x <= int('z') ) ||
				( int('A') <= x && x <= int('Z') ) ||
				( int('0') <= x && x <= int('9') )
			assert.True(t,ok)
		}
	}
}

func TestRandomAlphaDigitCharacter(t *testing.T){
	for i:=0 ; i < 1000; i++ {
		char := RandomAlphaDigitByte()
		x := int(char)
		ok := ( int('a') <= x && x <= int('z') ) ||
			( int('A') <= x && x <= int('Z') ) ||
			( int('0') <= x && x <= int('9') )
		assert.True(t,ok)
	}
}

func TestRandomString(t *testing.T){

	for alphabetLen := 1 ; alphabetLen < 10; alphabetLen ++ {
		alphabet := RandomStringExactLength(alphabetLen, "01abcxyzABCXYZ")
		for test := 0; test < 100; test++ {
			minLen := RandomInt(1, 100)
			maxLen := minLen + RandomInt(0, 3)
			str := RandomString(minLen, maxLen, alphabet)
			for j := 0; j < len(str); j++ {
				isInAlphabet := false;
				for k := 0; k < len(alphabet) && !isInAlphabet; k++ {
					isInAlphabet = isInAlphabet || alphabet[k] == str[j]
				}
				assert.True(t, isInAlphabet, fmt.Sprintf(" Error: TestRandomString %v must "+
					"be in the alphabet %v", str[j], alphabet))
			}
		}
	}
}

func TestRandomLowerEnglishString(t *testing.T){
	for i:=0 ; i < 100; i++ {
		test := RandomEnglishLowerCaseString(10,20)
		for _, char := range test {
			assert.GreaterOrEqual(t,int(char), int('a'))
			assert.LessOrEqual(t,int(char), int('z'))
		}
	}
}

func TestRandomLowerEnglishCharacter(t *testing.T){
	for i:=0 ; i < 1000; i++ {
		char := RandomAlphaLowerByte()
		assert.GreaterOrEqual(t, int(char), int('a'))
		assert.LessOrEqual(t, int(char), int('z'))
	}
}

func TestRandomUpperEnglishString(t *testing.T){
	for i:=0 ; i < 100; i++ {
		test := RandomEnglishUpperCaseString(10,20)
		for _, char := range test {
			assert.GreaterOrEqual(t,int(char), int('A'))
			assert.LessOrEqual(t,int(char), int('Z'))
		}
	}
}

func TestRandomUpperEnglishCharacter(t *testing.T){
	for i:=0 ; i < 1000; i++ {
		char := RandomAlphaUpperByte()
		assert.GreaterOrEqual(t,int(char), int('A'))
		assert.LessOrEqual(t,int(char), int('Z'))
	}
}

func TestDistributionOneDigit(t *testing.T){
	auxTestDistribution(1,9,10000, t)
}

func TestDistributionTwoDigit(t *testing.T){
	auxTestDistribution(10,99, 1000 , t)
}

func TestDistributionThreeDigit(t *testing.T){
	auxTestDistribution(100,199, 1000, t)
}

func TestDistributionOneNumber(t *testing.T){
	auxTestDistribution(123,123, 1000 , t)
}

func TestDistributionTwoNumbers(t *testing.T){
	auxTestDistribution(12,13, 1000 , t)
}

func auxTestDistribution( min int, max int, expectedRepetitions int , t *testing.T){
	var options = max-min+1
	totalTest := options*expectedRepetitions

	var maxError = expectedRepetitions/5
	frequency := make(map[int]int)
	for k:=min ; k <=max; k++{
		frequency[k] = 0
	}
	for i:=0 ; i < totalTest; i++ {
		frequency[RandomInt(min,max)]++
	}

	minAccepted := expectedRepetitions - maxError
	maxAccepted := expectedRepetitions + maxError

	//fmt.Printf("\nFrequencies: %v\n",frequency)
	//fmt.Printf(  "expected     : %v\n",expectedRepetitions)
	//fmt.Printf(  "maxError     : %v\n",maxError)
	//fmt.Printf(  "minAccepted  : %v\n",minAccepted)
	//fmt.Printf(  "maxAccepted  : %v\n",maxAccepted)
	//fmt.Printf("option -> frequency-expected\n")

	for option:=min ; option <=max; option++{
		assert.GreaterOrEqual(t,frequency[option], minAccepted)
		assert.LessOrEqual(t,frequency[option], maxAccepted)
		//fmt.Printf("%d -> %d\n",option,frequency[option]-expectedRepetitions)
	}
}

func TestRandomIntTwoDigits(t *testing.T) {
	auxCasesTestRandomInt(10,99, 1000, t)
}

func TestRandomIntOneDigit(t *testing.T) {
	auxCasesTestRandomInt(0,9, 100, t)
	auxCasesTestRandomInt(1,9, 100, t)
}

func TestRandomIntOneNumber(t *testing.T) {
	auxCasesTestRandomInt(12,12, 100, t)
}

func auxCasesTestRandomInt( from int, to int, testCount int , t *testing.T )  {

	var minValueFound int
	var maxValueFound int

	for i:=0 ; i < testCount ; i++{

		tmpTest := RandomInt(from,to)
		if i == 0 || tmpTest < minValueFound {
			minValueFound = tmpTest
		}
		if i == 0 || tmpTest > maxValueFound {
			maxValueFound = tmpTest
		}
	}
	assert.GreaterOrEqual(t,minValueFound, from)
	assert.LessOrEqual(t,maxValueFound, to)
}

func TestRandomFloat64(t *testing.T){

	auxCasesTestRandomFloat64(0,0.001,100000,t )
	auxCasesTestRandomFloat64(0,10,100000,t )
	auxCasesTestRandomFloat64(10,99.999999,100000,t )
	auxCasesTestRandomFloat64(1000,9999.99999,100000,t )
	for i := 0; i < 1000; i++ {
		from := RandomFloat64(0,10e10)
		to := from + RandomFloat64(0.000001,10e10)
		auxCasesTestRandomFloat64(from,to,1000,t)
	}
}

func auxCasesTestRandomFloat64( from float64, to float64, testCount int , t *testing.T )  {

	var minValueFound float64
	var maxValueFound float64
	frequency := make(map[float64]int)
	for i:=0 ; i < testCount ; i++{

		tmpTest := RandomFloat64(from,to)
		if i == 0 || tmpTest < minValueFound {
			minValueFound = tmpTest
		}
		if i == 0 || tmpTest > maxValueFound {
			maxValueFound = tmpTest
		}
		frequency[tmpTest]++
	}
	assert.GreaterOrEqual(t,minValueFound, from, fmt.Sprintf("Error RandomFloat64(%v,%v) = %v",from,to,minValueFound))
	assert.Less(t,maxValueFound, to, fmt.Sprintf("Error RandomFloat64(%v,%v) = %v",from,to,maxValueFound))
	//fmt.Printf("RandomFloat(%v,%v) = Min : %v  Max : %v.   %d test. %d different numbers generated\n",
	//	from,to,minValueFound,maxValueFound,testCount,len(frequency))
}

func TestRandomValidIntSet( t *testing.T){
	for test := 0 ;test < maxTestCasesSets; test++ {
		size := RandomInt(1, 10)
		minValue := RandomInt(1, size)
		maxValue := minValue + 2*size
		err, m := RandomIntSet(size, minValue, maxValue)
		assert.Nil(t, err)
		assert.NotNil(t, m)
		assert.Equal(t, size, len(m))
	}
}

func TestRandomValidLargeIntSet( t *testing.T){
	for test := 0 ;test < maxTestCasesSets; test++ {
		size := RandomInt(1, 10)
		minValue := RandomInt(1, size)
		maxValue := minValue + 2*size
		err, m := RandomIntSet(size, minValue, maxValue)
		assert.Nil(t, err)
		assert.NotNil(t, m)
		assert.Equal(t, size, len(m))
	}
}

func TestRandomValidInt64Set( t *testing.T){
	for test := 0 ;test < maxTestCasesSets; test++ {
		size := RandomInt(1, 10)
		minValue := RandomInt(1,size)
		maxValue := minValue+2*size
		err, m := RandomInt64Set(size, int64(minValue), int64(maxValue))
		assert.Nil(t, err)
		assert.NotNil(t, m)
		assert.Equal(t, size, len(m))
	}
}

func TestRandomValidLargeInt64Set( t *testing.T){
	for test := 0 ;test < maxTestCasesSets; test++ {
		size := RandomInt(1, 10)
		minValue := RandomInt(1,size)
		maxValue := minValue+2*size
		err, m := RandomInt64Set(size, int64(minValue), int64(maxValue))
		assert.Nil(t, err)
		assert.NotNil(t, m)
		assert.Equal(t, size, len(m))
	}
}

func TestRandomEmail( t *testing.T ){
	for i := 0 ; i < 1000 ; i++ {
		email := RandomEmail()
		atFirstIndex := strings.Index(email,"@")
		atLastIndex := strings.LastIndex(email,"@")
		assert.Equal(t,atFirstIndex,atLastIndex)
		assert.Greater(t,atFirstIndex,0)
		assert.Less(t,atFirstIndex,len(email)-1)
		assert.GreaterOrEqual(t,len(email),3)
	}
}

func TestRandomPhone( t *testing.T ){
	for i := 0 ; i < 100 ; i++ {
		phone := RandomPhoneNumber()
		assert.GreaterOrEqual(t,len(phone),6)
		assert.LessOrEqual(t,len(phone),15)
	}
}

func TestRandomAddress( t *testing.T ){
	for i := 0 ; i < 10 ; i++ {
		address := RandomAddressCOL()
		assert.GreaterOrEqual(t,len(address),5)
	}
}

func TestRandomStringSet( t *testing.T){

	for test := 0 ;test < maxTestCasesSets; test++{
		size := RandomInt(10,100)
		minValue := len(alphaDigits)
		maxValue := minValue + 2 * size
		err, m := RandomStringSet(size,minValue, maxValue, alphaDigits)
		assert.Nil(t, err)
		assert.Equal(t, size, len(m))
	}
}

func TestRandomValidPhoneSet( t *testing.T){

	for test := 0 ;test < maxTestCasesSets; test++{
		size := RandomInt(1,1000)
		err, m := RandomPhoneSet(size)
		assert.Nil(t, err)
		assert.Equal(t, size, len(m))
	}
}

func TestRandomValidEmailSet( t *testing.T){

	for test := 0 ;test < maxTestCasesSets; test++{
		size := RandomInt(1,1000)
		err, m := RandomEmailSet(size)
		assert.Nil(t, err)
		assert.Equal(t, size, len(m))
	}
}

func TestChooseInt( t *testing.T ) {

	for test := 0; test < maxTestCasesSets; test++ {
		size := RandomInt(1,100)
		from := RandomInt(0,10)
		to := from + RandomInt(0,5)
		err, elements := RandomIntSlice(size,from,to)
		chosen,err := ChooseInt(elements)
		assert.NotNil(t, chosen)
		assert.Nil(t, err)
		assert.Contains(t,elements, chosen, fmt.Sprintf("Error, %v must contain %v", elements, chosen))
	}
}

func TestChooseInt64( t *testing.T ) {

	for test := 0; test < maxTestCasesSets; test++ {
		size := RandomInt(1,100)
		from := RandomInt64(0,10)
		to := from + RandomInt64(0,5)
		err, elements := RandomInt64Slice(size,from,to)
		chosen,err := ChooseInt64(elements)
		assert.NotNil(t, chosen)
		assert.Nil(t, err)
		assert.Contains(t,elements, chosen, fmt.Sprintf("Error, %v must contain %v", elements, chosen))
	}
}

func TestChooseFloat64( t *testing.T ) {

	for test := 0; test < maxTestCasesSets; test++ {
		size := RandomInt(1,100)
		from := RandomFloat64(0,10)
		to := from + RandomFloat64(0.01,1)
		err, elements := RandomFloat64Slice(size,from,to)
		chosen,err := ChooseFloat64(elements)
		assert.NotNil(t, chosen)
		assert.Nil(t, err)
		assert.Contains(t,elements, chosen, fmt.Sprintf("Error, %v must contain %v", elements, chosen))
	}
}

func TestChooseString( t *testing.T ) {

	for test := 0; test < maxTestCasesSets; test++ {
		size := RandomInt(1, 100)
		from := RandomInt(1, 10)
		to := from + RandomInt(1, 3)
		err, elements := RandomStringSlice(size, from, to, alphaDigits)
		chosen, err := ChooseString(elements)
		assert.NotNil(t, chosen)
		assert.Nil(t, err)
		assert.Contains(t, elements, chosen, fmt.Sprintf("Error, %v must contain %v", elements, chosen))
	}
}

func TestRandomIntSlice( t *testing.T ) {
	for test := 0; test < 1000; test++ {
		size := RandomInt(1,100)
		from := RandomInt(0,10)
		to := from + RandomInt(0,5)
		err, slice := RandomIntSlice(size,from,to)
		assert.Nil(t,err)
		assert.NotNil(t,slice)
		assert.Len(t,slice,size)
		for _,item := range slice {
			assert.Contains(t,slice,item, fmt.Sprintf("The slice %v must contain %v",slice,item))
		}
	}
}

func TestRandomInt64Slice( t *testing.T ) {
	for test := 0; test < 1000; test++ {
		size := RandomInt(1,100)
		from := RandomInt64(0,10)
		to := from + RandomInt64(0,5)
		err, slice := RandomInt64Slice(size,from,to)
		assert.Nil(t,err)
		assert.NotNil(t,slice)
		assert.Len(t,slice,size)
		for _,item := range slice {
			assert.Contains(t,slice,item, fmt.Sprintf("The slice %v must contain %v",slice,item))
		}
	}
}


func TestRandomFloat64Slice( t *testing.T ) {
	for test := 0; test < 1000; test++ {
		size := RandomInt(1,100)
		from := RandomFloat64(0,1)
		to := from + RandomFloat64(0,1)
		err, slice := RandomFloat64Slice(size,from,to)
		assert.Nil(t,err)
		assert.NotNil(t,slice)
		assert.Len(t,slice,size)
		for _,item := range slice {
			assert.Contains(t,slice,item, fmt.Sprintf("The slice %v must contain %v",slice,item))
		}
	}
}

func TestRandomStringSlice( t *testing.T ) {
	for test := 0; test < 1000; test++ {
		size := RandomInt(1,100)
		from := RandomInt(0,1)
		to := from + RandomInt(0,5)
		err, slice := RandomStringSlice(size, from, to, RandomAlphaDigitStringExactLength(test+2))
		assert.Nil(t,err)
		assert.NotNil(t,slice)
		assert.Len(t,slice,size)
		for _,item := range slice {
			assert.Contains(t,slice,item, fmt.Sprintf("The slice %v must contain %v",slice,item))
		}
	}
}

func TestDisjointSet( t *testing.T){

	for p := 0.1 ; p <= 0.9 ; p = p+0.1 {
		for test := 0; test < 300; test++ {
			size := RandomInt(10, 20)
			minValue := RandomInt(10, 20)
			maxValue := minValue + 2*size

			err, m := RandomIntSet(size, minValue, maxValue)

			a, b := GetTwoDisjointSets(m, p)

			//fmt.Printf("\n P = %v , setSize = %v \n",p,size)
			//fmt.Printf("%v\n",m)
			//fmt.Printf("%v\n",a)
			//fmt.Printf("%v\n\n",b)

			assert.Nil(t, err)
			assert.NotNil(t, m)
			assert.Equal(t, size, len(m))
			assert.NotNil(t, a)
			assert.NotNil(t, b)
			for _,item := range a {
				assert.NotContains(t,b,item,fmt.Sprintf("%v must not contain %v",a,item))
			}
			for _,item := range b {
				assert.NotContains(t,b,item,fmt.Sprintf("%v must not contain %v",b,item))
			}
		}
	}
}