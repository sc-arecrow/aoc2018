package util

/* CharCount */

// CharCount : struct representing the counts of each lowercase letter in a string
type CharCount struct {
    str string
    count [26]int
}

// NewCharCount : creates a new CharCount with the given string
func NewCharCount(str string) CharCount {
    var charCount CharCount
    charCount.str = str
    charCount.count = charCount.getCountFromString()
    return charCount
}

// GetString : returns the string represented by the charCount
func (cc CharCount) GetString() string {
    return cc.str
}

// GetCount : returns the count of the string represented by the charCount
func (cc CharCount) GetCount() [26]int {
    return cc.count
}

func (cc CharCount) getCountFromString() [26]int {
    var count [26]int

    for _, char := range cc.str {
        count[AlphabetToInt(char)]++
    }

    return count
}

/* IntSet */

var exists = struct{}{}

// IntSet : struct representing a set of distinct integer values
type IntSet struct {
	m map[int]struct{}
}

// NewIntSet : creates a new IntSet and returns a pointer to it
func NewIntSet() *IntSet {
	s := &IntSet{}
	s.m = make(map[int]struct{})
	return s
}

// Add : adds given integer value into the IntSet
func (s *IntSet) Add(value int) {
	s.m[value] = exists
}

// Remove : removes given integer value if present from the IntSet
func (s *IntSet) Remove(value int) {
	delete(s.m, value)
}

// Contains : returns true if given integer value is inside the IntSet
func (s *IntSet) Contains(value int) bool {
	_, ok := s.m[value]
	return ok
}

// Size : returns the size of the IntSet
func (s *IntSet) Size() int {
	return len(s.m)
}

// Values : returns the values in the IntSet
func (s *IntSet) Values() []int {
    values := []int{}

    if len(s.m) > 0 {
        for value := range s.m {
            values = append(values, value)
        }
    }

    return values
}