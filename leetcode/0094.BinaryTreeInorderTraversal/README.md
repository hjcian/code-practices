# Question
https://leetcode.com/problems/binary-tree-inorder-traversal/

# Thoughts

**inorder**
- 需滿足「左中右」的走訪順序
- 需要一個 `stack` 來裝還不能走的 node
- 需要一個 `ret` 來裝結果

走到目前的 node 時，因為必須先看 left，故先把自己裝進 `stack` 裡，並且 `root = root.left` 來往下一個迭代走

若 `left == nil`，那麼就表示「左中右」的「左」已到底，此時需要將「中」的 node 裝進結果裡，故從 stack 中 pop 出 node 來裝結果

並且因為「左中右」的規則，接下來要處理「右」，故指定下一個迭代為它的 right node `root = node.right`

如此一來就正確地以 in-order 的順序走完 tree

**preorder**
- 需滿足「中左右」的走訪順序
- 需要一個 `stack` 來裝還不能走的 node
- 需要一個 `ret` 來裝結果

走到目前的 node 時，因為可先走訪「中」，故將值裝進結果中

而因為下一個要看的是 left，故 `root = root.left`，但同時，right node 就得先放到 stack 裡（檢查 `!= nil` 再放）

下一次迭代若發現此時 `root == nil`，就需要從 stack pop node 出來，重新 assign 給 root 就好，下一次迭代自然會處理好

**postorder**
- 需滿足「左右中」的走訪順序
- 需要一個 `stack` 來裝還不能走的 node
- 需要一個 `ret` 來裝結果


走訪到此 node 時
- if node != nil
  - 則將 right node 與自己裝到 stack 裡，然後指派下一次迭代為 left node
- else
  - 從 stack 裡取點出來
  - 檢查此 node 有沒有 right node
    - 若有，則指派下一次迭代為 right node
    - 若無，則將值存入結果中
