package main

func main() {



}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	_, ok, ansList := dfs(root, 0, voyage)
	if ok {
		return ansList
	}
	return []int{-1}
}

func dfs(root *TreeNode, cur int, vayage []int) (int, bool, []int) {
	if root == nil{
		return cur, true, nil
	}
	var ansList []int
	if root.Val != vayage[cur] {
		return cur + 1, false, ansList
	}
	if cur+1 == len(vayage) {
		return cur + 1, true, ansList
	}
	if root.Right == nil {
		return dfs(root.Left, cur + 1, vayage)
	}
	if root.Left == nil {
		return dfs(root.Right, cur + 1, vayage)
	}
	if root.Left.Val != vayage[cur+1] {
		if root.Right.Val != vayage[cur+1] {
			return cur + 1, false, ansList
		}
		root.Left, root.Right = root.Right, root.Left
		ansList = append(ansList, root.Val)
	}
	tmp, ok, leftansList := dfs(root.Left, cur+1, vayage)
	if !ok {
		return tmp, false, ansList
	}
	tmp, ok, rightAnsList := dfs(root.Right, tmp, vayage)
	if !ok {
		return tmp, false, ansList
	}
	return tmp, true, append(append(ansList, leftansList...), rightAnsList...)
}
