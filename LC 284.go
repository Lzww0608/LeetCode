/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
    iterator *Iterator
    hasPeeked bool
    peekedElement *int
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{
        iterator: iter,
    }
}

func (this *PeekingIterator) hasNext() bool {
    return this.hasPeeked || this.iterator.hasNext()
}

func (this *PeekingIterator) next() int {
    if this.hasPeeked {
        this.hasPeeked = false
        result := this.peekedElement
        this.peekedElement = nil
        return *result
    }
    return this.iterator.next()
}

func (this *PeekingIterator) peek() int {
    if !this.hasPeeked {
        nextVal := this.iterator.next()
        this.peekedElement = &nextVal
        this.hasPeeked = true
    }
    return *this.peekedElement
}