package helpers

type Graph struct {
	nodeKeys map[int]*Node
}

type Node struct {
	key      int
	parents  []Vertex
	children []Vertex
	count    int
}

type Vertex struct {
	parent *Node
	child  *Node
}

func NewGraph() *Graph {
	return &Graph{
		make(map[int]*Node),
	}
}

func (g *Graph) addNode(key int) *Node {
	node, ok := g.nodeKeys[key]
	if !ok {
		node = &Node{
			key:   key,
			count: 0,
		}
		g.nodeKeys[key] = node
	}
	return node
}

func (g *Graph) AddChildToParent(childKey, parentKey int) {
	parentNode, ok := g.nodeKeys[parentKey]
	if !ok {
		parentNode = g.addNode(parentKey)
	}
	childNode, ok := g.nodeKeys[childKey]
	if !ok {
		childNode = g.addNode(childKey)
	}
	childNode.count += 1

	vertex := Vertex{
		parent: parentNode,
		child:  childNode,
	}

	parentNode.children = append(parentNode.children, vertex)
	childNode.parents = append(childNode.parents, vertex)
}

func (g *Graph) GetPaths(key int) int {
	sum := 0
	deviceNode, ok := g.nodeKeys[key]
	if !ok {
		return 0
	}
	sum += len(deviceNode.parents)
	for _, parVer := range deviceNode.parents {
		sum += g.GetPaths(parVer.parent.key)
	}
	return sum
}

func (g *Graph) GetPath(start, end int) []int {
	node, ok := g.nodeKeys[start]
	keys := []int{}
	if !ok {
		return keys
	}

	for _, chVer := range node.children {
		child := chVer.child
		if child.key == end {
			return []int{child.key}
		}
		path := g.GetPath(child.key, end)
		if len(path) > 0 {
			keys = append(keys, child.key)
			keys = append(keys, path...)
		}
	}
	return keys
}
