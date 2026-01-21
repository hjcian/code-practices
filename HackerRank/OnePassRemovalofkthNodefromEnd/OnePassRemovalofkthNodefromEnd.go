package onepassremovalofkthnodefromend

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string) {
	for node != nil {
		fmt.Printf("%d", node.data)

		node = node.next

		if node != nil {
			fmt.Printf(sep)
		}
	}
}

/*
 * Complete the 'removeKthNodeFromEnd' function below.
 *
 * The function is expected to return an INTEGER_SINGLY_LINKED_LIST.
 * The function accepts following parameters:
 *  1. INTEGER_SINGLY_LINKED_LIST head
 *  2. INTEGER k
 */

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */

func removeKthNodeFromEnd(head *SinglyLinkedListNode, k int32) *SinglyLinkedListNode {
	// Write your code here
	mem := make(map[int32]*SinglyLinkedListNode, 0) // order - node
	start := head
	var order int32 = -1
	for start != nil {
		order++
		mem[order] = start
		start = start.next
	}

	if k > int32(order) {
		// invalid k
		return head
	}
	if order-k == 0 {
		// means that it removes the head element
		return head.next
	}
	// fmt.Println("order:", order, "k:", k)
	prev := mem[(order-k)-1] // prev node
	target := mem[order-k]
	prev.next = target.next // remove the target from the linked list
	return head
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     headCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     head := SinglyLinkedList{}
//     for i := 0; i < int(headCount); i++ {
//         headItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//         checkError(err)
//         headItem := int32(headItemTemp)
//         head.insertNodeIntoSinglyLinkedList(headItem)
//     }

//     kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)
//     k := int32(kTemp)

//     result := removeKthNodeFromEnd(head.head, k)

//     printSinglyLinkedList(result, "\n")
//     fmt.Printf("\n")
// }

// func readLine(reader *bufio.Reader) string {
//     str, _, err := reader.ReadLine()
//     if err == io.EOF {
//         return ""
//     }

//     return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
//     if err != nil {
//         panic(err)
//     }
// }
