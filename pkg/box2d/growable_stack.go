package box2d

// Adapted from https://gist.github.com/bemasher/1777766

type GrowableStack struct {
	top  *StackElement
	size int
}

type StackElement struct {
	value interface{} // All types satisfy the empty interface, so we can store anything here.
	next  *StackElement
}

// Return the stack's length
func (s GrowableStack) GetCount() int {
	return s.size
}

// Push a new element onto the stack
func (s *GrowableStack) Push(value interface{}) {
	s.top = &StackElement{value, s.top}
	s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *GrowableStack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}
