package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)
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
		assert.GreaterOrEqual(t,int(char), int('a'))
		assert.LessOrEqual(t,int(char), int('z'))
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

	var maxError = expectedRepetitions/10
	frequency := make(map[int]int)
	for k:=min ; k <=max; k++{
		frequency[k] = 0
	}
	for i:=0 ; i < totalTest; i++ {
		frequency[RandomInt(min,max)]++
	}

	minAccepted := expectedRepetitions - maxError
	maxAccepted := expectedRepetitions + maxError

	fmt.Printf("\nFrequencies: %v\n",frequency)
	fmt.Printf(  "expected     : %v\n",expectedRepetitions)
	fmt.Printf(  "maxError     : %v\n",maxError)
	fmt.Printf(  "minAccepted  : %v\n",minAccepted)
	fmt.Printf(  "maxAccepted  : %v\n",maxAccepted)
	fmt.Printf("option -> frequency-expected\n")

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
func TestRandomValidIntSet( t *testing.T){
	minValue := 1
	maxValue := 10
	size := 5
	err, m := getRandomIntSet(size,minValue,maxValue)
	assert.Nil(t,err)
	assert.Equal(t,size,len(m))
}
func TestRandomInValidIntSet( t *testing.T){
	minValue := 1
	maxValue := 10
	size := 11
	err, _ := getRandomIntSet(size,minValue,maxValue)
	assert.NotNil(t,err)
	//assert.Equal(t,size,len(m))
}
func TestRandomValidLargeIntSet( t *testing.T){
	minValue := 1
	maxValue := 999999999
	size := 200000
	err, m := getRandomIntSet(size,minValue,maxValue)
	assert.Nil(t,err)
	assert.Equal(t,size,len(m))
}
func TestRandomValidInt64Set( t *testing.T){
	minValue := 1
	maxValue := 10
	size := 5
	err, m := getRandomInt64Set(size,int64(minValue),int64(maxValue))
	assert.Nil(t,err)
	assert.Equal(t,size,len(m))
}
func TestRandomInValidInt64Set( t *testing.T){
	minValue := 1
	maxValue := 10
	size := 11
	err, _ := getRandomInt64Set(size,int64(minValue),int64(maxValue))
	assert.NotNil(t,err)
	//assert.Equal(t,size,len(m))
}
func TestRandomValidLargeInt64Set( t *testing.T){
	var minValue int64
	minValue = 99999999
	var maxValue int64
	maxValue = 99999999900000000
	size := 200000
	err, m := getRandomInt64Set(size,int64(minValue),int64(maxValue))
	assert.Nil(t,err)
	assert.Equal(t,size,len(m))
}
