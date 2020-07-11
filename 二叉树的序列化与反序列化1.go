/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package leadcode

/**
type Codecs struct {
	data []string
}

func Constructor() Codecs {
	return Codecs{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	re := this.rserialize(root, "")
	return re
}

func (this *Codec) rserialize(root *TreeNode, str string) string {
	if root == nil {
		str += "null,"
	} else {
		str += strconv.Itoa(root.Val) + ","
		str = this.rserialize(root.Left, str)
		str = this.rserialize(root.Right, str)
	}

	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	re := strings.Split(strings.Trim(data, ","), ",")
	this.data = re
	return this.rdeserialize()
}

func (this *Codec) rdeserialize() *TreeNode {
	if this.data[0] == "null" {
		this.data = this.data[1:]
		return nil
	}
	val, _ := strconv.Atoi(this.data[0])
	root := &TreeNode{Val: val}
	this.data = this.data[1:]
	root.Left = this.rdeserialize()
	root.Right = this.rdeserialize()
	return root
}
**/
/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
