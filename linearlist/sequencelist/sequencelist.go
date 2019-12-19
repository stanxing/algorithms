package sequencelist

import "errors"

const maxLength = 20

// SequenceList sequence list
type SequenceList struct {
	maxLength int
	length    int
	data      []interface{}
}

// MakeEmptyList make an empty sequence list
func MakeEmptyList() *SequenceList {
	return &SequenceList{
		maxLength: maxLength,
		data:      make([]interface{}, 0, maxLength),
	}
}

// FindElement find the element according to the position
func (s *SequenceList) FindElement(k int) (interface{}, error) {
	if k <= 0 || k > s.length {
		return nil, errors.New("illegal position")
	}
	return s.data[k-1], nil

}

// FindPosition find the postion where the element first appearance
func (s *SequenceList) FindPosition(element interface{}) int {
	for index, item := range s.data {
		if element == item {
			return index
		}
	}
	return -1
}

// Append insert an element at the end
func (s *SequenceList) Append(element interface{}) error {
	if s.IsFull() {
		return errors.New("The list is full")
	}
	s.data = append(s.data, element)
	s.length++
	return nil
}

// Insert insert an element at specified position
func (s *SequenceList) Insert(element interface{}, k int) (bool, error) {
	if s.isIndexOver(k) {
		return false, errors.New("index over")
	}
	// prevent index out of range
	s.Append("")
	for i := s.length - 1; i > k-1; i-- {
		s.data[i] = s.data[i-1]
	}
	s.data[k-1] = element
	return true, nil
}

// Delete delete an element
func (s *SequenceList) Delete(k int) (interface{}, error) {
	if s.isIndexOver(k) {
		return nil, errors.New("index over")
	}
	if s.IsEmpty() {
		return nil, errors.New("The list is null")
	}
	data := s.data[k-1]
	for i := k; i < s.length; i++ {
		s.data[i-1] = s.data[i]
	}
	s.length--
	// after moving position, the last element need to be removed
	s.data = s.data[:s.length]
	return data, nil
}

// Length return list length
func (s *SequenceList) Length() int {
	return s.length
}

// IsEmpty whether list is null
func (s *SequenceList) IsEmpty() bool {
	if s.length == 0 {
		return true
	}
	return false
}

// IsFull whether list is full
func (s *SequenceList) IsFull() bool {
	if s.length == s.maxLength {
		return true
	}
	return false
}

func (s *SequenceList) isIndexOver(k int) bool {
	if k < 1 || k > s.length {
		return true
	}
	return false
}
