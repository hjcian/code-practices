package detectcycleinmoduledependencygraph

/*
 * Complete the 'hasCircularDependency' function below.
 *
 * The function is expected to return a BOOLEAN.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY dependencies
 */

type Node struct {
	Val int32
	In  map[int32]*Node
	Out map[int32]*Node
}

func NewNode(val int32) *Node {
	return &Node{
		Val: val,
		In:  make(map[int32]*Node, 0),
		Out: make(map[int32]*Node, 0),
	}
}

func (n *Node) From(node *Node) {
	n.In[node.Val] = node
}

func (n *Node) RemoveIn(node *Node) {
	delete(n.In, node.Val)
}

func (n *Node) To(node *Node) {
	n.Out[node.Val] = node
}

func (n *Node) RemoveOut(node *Node) {
	delete(n.Out, node.Val)
}

func hasCircularDependency(n int32, dependencies [][]int32) bool {
	// Write your code here
	nodes := make(map[int32]*Node, 0)
	for _, modules := range dependencies {
		// corner case: self-dependencies
		if modules[0] == modules[1] {
			return true
		}

		current, ok := nodes[modules[0]]
		if !ok {
			current = NewNode(modules[0])
			nodes[modules[0]] = current
		}
		next, ok := nodes[modules[1]]
		if !ok {
			next = NewNode(modules[1])
			nodes[modules[1]] = next
		}
		current.To(next)
		next.From(current)
	}
	queue := make([]*Node, 0)
	for _, node := range nodes {
		if len(node.In) == 0 {
			// fmt.Println(node)
			queue = append(queue, node)
		}
	}
	result := 0
	// fmt.Println(queue)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result++
		for _, out := range node.Out {
			out.RemoveIn(node)
			if len(out.In) == 0 {
				queue = append(queue, out)
			}
		}
	}

	return result != len(nodes)
}

// func main() {
//     reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

//     nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)
//     n := int32(nTemp)

//     dependenciesRows, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     dependenciesColumns, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
//     checkError(err)

//     var dependencies [][]int32
//     for i := 0; i < int(dependenciesRows); i++ {
//         dependenciesRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

//         var dependenciesRow []int32
//         for _, dependenciesRowItem := range dependenciesRowTemp {
//             dependenciesItemTemp, err := strconv.ParseInt(dependenciesRowItem, 10, 64)
//             checkError(err)
//             dependenciesItem := int32(dependenciesItemTemp)
//             dependenciesRow = append(dependenciesRow, dependenciesItem)
//         }

//         if len(dependenciesRow) != int(dependenciesColumns) {
//             panic("Bad input")
//         }

//         dependencies = append(dependencies, dependenciesRow)
//     }

//     result := hasCircularDependency(n, dependencies)

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
