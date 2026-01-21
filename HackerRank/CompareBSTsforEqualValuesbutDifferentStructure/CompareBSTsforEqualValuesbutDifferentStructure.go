package comparebstsforequalvaluesbutdifferentstructure

/*
 * Complete the 'verifySameMultisetDifferentStructure' function below.
 *
 * The function is expected to return a BOOLEAN.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY root1
 *  2. INTEGER_ARRAY root2
 */

/*
Given two binary search trees root1 and root2, return true if they contain the same multiset of values but have different structures, otherwise return false.

# Example

Input
```
root1 = [4, 2, 5, 1, 3, 100001, 100001]
root2 = [3, 1, 5, 100001, 2, 4, 100001]
```

Output
```
true
```
# [Explanation]
- First, collect the values of each tree (ignoring the sentinel 100001 for nulls).
- Tree1 has values [4, 2, 5, 1, 3], and Tree2 has [3, 1, 5, 2, 4].
- Sorting both gives [1,2,3,4,5] in each, so the multisets match.
- Next, compare structures: Tree1's root is 4 with children 2 (which itself has children 1 and 3) and 5; Tree2's root is 3 with children 1 (right child 2) and 5 (left child 4). - The shapes differ, so the function returns true.


# [Constraints]
0 <= root1.length <= 1000
0 <= root2.length <= 1000
BST property holds: for every node, all values in its left subtree <= node.value <= all values in its right subtree
Trees may contain duplicate values
*/

import "slices"

func verifySameMultisetDifferentStructure(root1 []int32, root2 []int32) bool {
	// Write your code here
	if len(root1) == 0 || len(root2) == 0 {
		return false
	}
	nums1 := make([]int32, 0)
	for i := range root1 {
		if root1[i] != 100001 {
			nums1 = append(nums1, root1[i])
		}
	}
	nums2 := make([]int32, 0)
	for i := range root2 {
		if root2[i] != 100001 {
			nums2 = append(nums2, root2[i])
		}
	}
	if len(nums1) != len(nums2) {
		return false
	}
	slices.Sort(nums1)
	slices.Sort(nums2)
	for i := 0; i < len(nums1); i++ {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	for i := 0; i < len(root1) && i < len(root2); i++ {
		if root1[i] != root2[i] {
			return true
		}
	}
	return false
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     root1Count, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var root1 []int32

//     for i := 0; i < int(root1Count); i++ {
//         root1ItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//         checkError(err)
//         root1Item := int32(root1ItemTemp)
//         root1 = append(root1, root1Item)
//     }

//     root2Count, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var root2 []int32

//     for i := 0; i < int(root2Count); i++ {
//         root2ItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//         checkError(err)
//         root2Item := int32(root2ItemTemp)
//         root2 = append(root2, root2Item)
//     }

//     result := verifySameMultisetDifferentStructure(root1, root2)

//     fmt.Printf("%s\n", (map[bool]string{true: "1", false: "0"})[result])
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
