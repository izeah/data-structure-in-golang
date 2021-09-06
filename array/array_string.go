package array

import (
	"errors"
	"strings"
)

// Slice of string as a new data type.
type arrOfStr []string

// func of mapCallback as a new data type.
type mapCallbackFn func(value string) string

// func of forEachCallback as a new data type.
type forEachCallbackFn func(value string)

// func of filterCallback as a new data type.
type filterCallbackFn func(value string) bool

// Creates a new array of string.
func NewArrayOfString() arrOfStr {
	return arrOfStr{}
}

// Get all elements of the array.
func (a arrOfStr) GetAll() arrOfStr {
	return a
}

// Get all reversed elements of the array.
func (a arrOfStr) GetAllReversed() arrOfStr {
	var newArr arrOfStr
	for i := len(a) - 1; i >= 0; i-- {
		newArr = append(newArr, a[i])
	}

	return newArr
}

// Get one item of the array.
func (a arrOfStr) GetOne(index int) string {
	if index < 0 || index >= len(a) {
		return ""
	}

	return a[index]
}

// Update one item of the array.
func (a *arrOfStr) UpdateOne(index int, newStr string) error {
	if index < 0 || index >= len(*a) {
		return errors.New("data tidak ditemukan")
	}

	(*a)[index] = newStr
	return nil
}

// Delete one item of the array.
func (a *arrOfStr) DeleteOne(index int) error {
	if index < 0 || index >= len(*a) {
		return errors.New("data tidak ditemukan")
	}

	*a = append((*a)[:index], (*a)[index+1:]...)
	return nil
}

// Appends new elements to the end of an array, and returns the new length of the array.
func (a *arrOfStr) Append(elems ...string) int {
	*a = append((*a), elems...)
	return len(*a)
}

// Inserts new elements at the start of an array, and "unshifts" older elements.
// Returns the new length of the array.
func (a *arrOfStr) Prepend(elems ...string) int {
	var revArr arrOfStr
	for i := len(elems) - 1; i >= 0; i-- {
		revArr = append(revArr, elems[i])
	}
	*a = append(revArr, *a...)
	return len(*a)
}

// Removes the last element from an array and returns it. If the array is empty, array is not modified.
func (a *arrOfStr) Pop() string {
	index := len(*a) - 1
	popped := (*a)[index]
	if index < 0 {
		return ""
	}

	*a = append(arrOfStr{}, (*a)[:index]...)
	return popped
}

// Shifting is equivalent to popping, working on the first element instead of the last.
// Removes the first array element and "shifts" all other elements to a lower index.
// Returns the value that was "shifted out".
func (a *arrOfStr) Shift() string {
	if len(*a) > 0 {
		shifted := (*a)[0]

		*a = append((*a)[:0], (*a)[1:]...)

		return shifted
	}

	return ""
}

// Removes elements from an array and, if necessary, inserts new elements in their place, returning the deleted elements.
func (a *arrOfStr) Splice(start int, deleteCount int, items ...string) arrOfStr {
	splicedArr := append(arrOfStr{}, (*a)...)
	if len(items) > 0 {
		*a = append((*a)[:start], append(items, (*a)[start+deleteCount:]...)...)
	} else {
		*a = append((*a)[:start], (*a)[start+deleteCount:]...)
	}

	return splicedArr[start : start+deleteCount]
}

// Creates a new array by performing a function on each array element.
// does not execute the function for array elements without values.
// does not change the original array.
func (a *arrOfStr) Map(callbackFn mapCallbackFn) arrOfStr {
	var newArr arrOfStr
	for _, v := range *a {
		newArr = append(newArr, callbackFn(v))
	}

	return newArr
}

// Performs the specified action for each element in an array.
func (a *arrOfStr) ForEach(callbackFn forEachCallbackFn) {
	for _, v := range *a {
		callbackFn(v)
	}
}

// Creates a new array with array elements that passes a test.
func (a *arrOfStr) Filter(callbackFn filterCallbackFn) arrOfStr {
	var filteredArr arrOfStr
	for _, v := range *a {
		if callbackFn(v) {
			filteredArr = append(filteredArr, v)
		}
	}

	return filteredArr
}

// Returns the value of the first array element that passes a test function.
func (a *arrOfStr) Find(callbackFn filterCallbackFn) string {
	for _, v := range *a {
		if callbackFn(v) {
			return v
		}
	}

	return ""
}

// Determines whether an array includes a certain element, returning true or false as appropriate.
func (a *arrOfStr) Includes(findStr string) bool {
	for _, v := range *a {
		if v == findStr {
			return true
		}
	}

	return false
}

// Converts an array to a joined string of array values with comma separated.
func (a *arrOfStr) ToString() string {
	return strings.Join(*a, ",")
}
