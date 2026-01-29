package reverseevenindexednodesandappend

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
 * Complete the 'extractAndAppendSponsoredNodes' function below.
 *
 * The function is expected to return an INTEGER_SINGLY_LINKED_LIST.
 * The function accepts INTEGER_SINGLY_LINKED_LIST head as parameter.
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

func print(head *SinglyLinkedListNode) {
	nums := make([]int32, 0)
	for head != nil {
		nums = append(nums, head.data)
		head = head.next
	}
	fmt.Println(".   ", nums)
}

func extractAndAppendSponsoredNodes(head *SinglyLinkedListNode) *SinglyLinkedListNode {
	// Write your code here
	var currHead, newHead, tailHead *SinglyLinkedListNode = nil, nil, nil
	idx := 0
	for head != nil {
		next := head.next
		if idx%2 == 0 {
			tmp := tailHead
			head.next = tmp
			tailHead = head
			//fmt.Println("tail:")
			//print(tailHead)
		} else {
			if newHead == nil {
				newHead = head
				currHead = head
			} else {
				currHead.next = head
				currHead = head
			}
			//fmt.Println("curr:")
			//print(currHead)
		}
		idx++
		head = next
	}
	if newHead == nil {
		return tailHead
	}
	currHead.next = tailHead
	return newHead
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

//     result := extractAndAppendSponsoredNodes(head.head)

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
