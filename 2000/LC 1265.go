/*   Below is the interface for ImmutableListNode, which is already defined for you.
 *
 *   type ImmutableListNode struct {
 *       
 *   }
 *
 *   func (this *ImmutableListNode) getNext() ImmutableListNode {
 *		// return the next node.
 *   }
 *
 *   func (this *ImmutableListNode) printValue() {
 *		// print the value of this node.
 *   }
 */

func printLinkedListInReverse(head ImmutableListNode) {
    myPrint(head, nil)
}

func myPrint(l, r ImmutableListNode) {
    if l.getNext() == r {
        l.printValue()
        return
    }

    var mid ImmutableListNode
    for slow, fast := l, l; ; {
        mid = slow
        if fast == r || fast.getNext() == r {
            break
        }
        slow, fast = slow.getNext(), fast.getNext().getNext()
    }

    myPrint(mid, r)
    myPrint(l, mid)
}