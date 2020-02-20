package structure

import "container/heap"

// Item 優先度付きキューのデータ
type Item struct {
	value string
	priority int
	index int
}

// Items 優先度付きキューのデータ構造
type Items []*Item

// Len 長さ
func (it Items) Len() int { return len(it) }

// Less 優先度を判定する
func (it Items) Less(i, j int) bool {
	return it[i].priority > it[j].priority
}

// Swap インデクスを入れ替える
func (it Items) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
	it[i].index = i
	it[j].index = j
}

// Push 優先度付きキューに追加する
func (it *Items) Push(x interface{}) {
	n := len(*it)
	item := x.(*Item)
	item.index = n
	*it = append(*it, item)
}

// Pop 優先度付きキューからアイテムを取得する
func (it *Items) Pop() interface{} {
	old := *it
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*it = old[0:n-1]
	return item
}

func (it *Items) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(it, item.index)
}

// PriorityQueue 優先度付きキュー
type PriorityQueue struct {
	items Items
}

// NewPriorityQueue コンストラクタ
func NewPriorityQueue() *PriorityQueue {
	var items Items
	heap.Init(&items)
	return &PriorityQueue{
		items: items,
	}
}

// Push push!!
func (pq *PriorityQueue) Push(item *Item) {
	heap.Push(&pq.items, item)
}

// Pop pop!!
func (pq *PriorityQueue) Pop() *Item {
	return heap.Pop(&pq.items).(*Item)
}

// Empty 空の判定
func (pq *PriorityQueue) Empty() bool {
	return len(pq.items) == 0
}