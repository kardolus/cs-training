/*
From Leetcode (https://leetcode.com/explore/challenge/card/july-leetcoding-challenge/546/week-3-july-15th-july-21st/3394/)

There are a total of n courses you have to take, labeled from 0 to n-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, return the ordering of courses you should take to finish all courses.

There may be multiple correct orders, you just need to return one of them. If it is impossible to finish all courses, return an empty array.
/*

func findOrder(numCourses int, prerequisites [][]int) []int {  
    if len(prerequisites) == 0 { // Handling boundary conditions
        var result []int
        for i := numCourses; i > 0; i-- {
            result = append(result, i - 1)
        }
        return result
    }
    
    graph := make(map[int]*LinkedList)
    for _, classPair := range prerequisites {
        firstClass := classPair[1]
        lastClass := classPair[0]
        
        if _, ok := graph[lastClass]; !ok {
            graph[lastClass] = new(LinkedList)
        }
        
        if _, ok := graph[firstClass]; !ok {
            graph[firstClass] = NewLinkedList(lastClass)
            continue
        }
        
        graph[firstClass].Add(&Node{
            value: lastClass,
        })
    }         
    
    var (
        lengths []int
        nodes []int
    )
    
    for key, _ := range graph {
        // fmt.Printf("Root [%d] = [%s]\n", key, PrintList(value))
        // fmt.Printf("DFS[%d] = [%d]\n", key, DepthFirstSearch(graph, key))
        
        length := DepthFirstSearch(graph, key)
        
        lengths = append(lengths, length)
        nodes = append(nodes, key)
        
        for i := 0; i < len(lengths) - 1; i++ {
            if lengths[i] < lengths[len(lengths) - 1] {
                lengths[i], lengths[len(lengths) - 1] = lengths[len(lengths) - 1], lengths[i]
                nodes[i], nodes[len(nodes) - 1] = nodes[len(nodes) - 1], nodes[i]
            }
        }
    }
    
    return nodes[:numCourses]
}

func DepthFirstSearch(graph map[int]*LinkedList, root int) int {
    var result int
    
    if _, ok := graph[root]; !ok {
        return 0
    }
    
    focus := graph[root].head
    
    for {
        if focus == nil {
            break
        }
        result++
        
        DepthFirstSearch(graph, focus.value)
        focus = focus.next
    }
    
    return result
}

type Node struct {
    value int
    next *Node
}

type LinkedList struct { // Could have a "visited" attribute to account for circular graphs
    head *Node
}

func (l *LinkedList) Add(node *Node) {
    tmp := l.head
    l.head = node
    l.head.next = tmp
}

func NewLinkedList(value int) *LinkedList {
    var result LinkedList
    
    result.head = &Node {
        value: value,
    }
    
    return &result
}

func PrintList(list *LinkedList) string {
    var result string
    
    focus := list.head
    for focus != nil {
        result += fmt.Sprintf("[%d]-->", focus.value)
        focus = focus.next
    }
    
    return result
}
