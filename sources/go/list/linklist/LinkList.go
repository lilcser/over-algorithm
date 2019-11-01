/*
 *  单链表
 */
package linklist

import (
	"errors"
	"fmt"
)

type ElemeType int // 给 int起个别名

// 单链表 结点结构体
type Node struct {
	data ElemeType
	next *Node
}

// 单链结构体
type LinkList struct {
	length int // 该线性表最大长度
	first  *Node
}

// 构造表
func NewLinkList() *LinkList {
	return &LinkList{
		length: 0,
		first:  nil,
	}
}

// 获取长度
func (l *LinkList) Length() int {
	return l.length
}

// 显示表
func (l *LinkList) Display() {

	if l.length == 0 {
		fmt.Println("数据结构内无元素")
		return
	}

	currentNode := l.first
	for i := 1; i <= l.length; i++ {
		fmt.Printf("位置 %d 元素为：%d \n", i, currentNode.data)
		if i == l.length {
			break
		}
		currentNode = currentNode.next
	}

}

func (l *LinkList) Append(e ElemeType) {

	// 构造要插入的节点
	insertNode := &Node{
		data: e,
		next: nil,
	}

	// 第一次追加
	if l.length == 0 {
		l.first = insertNode
		l.length++
		return
	}

	currentNode := l.first
	for i := 1; i <= l.length; i++ {
		if i == l.length {
			currentNode.next = insertNode
		}
	}
	l.length++
}

// 插入元素
func (l *LinkList) Insert(index int, e ElemeType) error {

	if index < 1 || index > l.length {
		return errors.New("位序不正确")
	}

	// 构造要插入的节点
	insertNode := &Node{
		data: e,
		next: nil,
	}

	// 如果在第一个结点插入
	if index == 1 {
		insertNode.next = l.first
		l.first = insertNode
		l.length++
		return nil
	}

	// 插入：找到插入位置的前一个节点
	currentNode := l.first
	for i := 1; i <= index-1; i++ {
		currentNode = currentNode.next
	}
	insertNode.next = currentNode.next
	currentNode.next = insertNode
	l.length++
	return nil

}

// 删除 按照位序删除
func (l *LinkList) Delete(index int) error {

	if index < 1 || index > l.length+1 {
		return errors.New("位序不正确")
	}

	// 找到要删除元素的前一个元素
	currentNode := l.first
	for i := 1; i <= index-1; i++ {
		currentNode = currentNode.next
	}
	if index == l.length { // 如果要删除的是最后一个元素
		currentNode.next = nil
		l.length--
		return nil
	}

	currentNode.next = currentNode.next.next
	l.length--
	return nil

}

// 修改 按照位序修改
func (l *LinkList) Update(index int, e ElemeType) error {

	if index < 1 || index > l.length {
		return errors.New("位序不正确")
	}

	currentNode := l.first
	for i := 1; i <= index; i++ {
		if i == index {
			currentNode.data = e
			break
		}
		currentNode = currentNode.next
	}
	return nil
}

// 查询 按照位序查询值
func (l *LinkList) GetElem(index int) (e ElemeType, err error) {
	if index < 1 || index > l.length {
		err = errors.New("位序不合法")
		return
	}
	currentNode := l.first
	for i := 1; i <= index; i++ {
		if i == index {
			e = currentNode.data
			break
		}
		currentNode = currentNode.next
	}
	return
}

// 查询 按照值查询位序
func (l *LinkList) LocateElem(e ElemeType) (index int, err error) {
	currentNode := l.first
	for i := 1; i <= l.length; i++ {
		if currentNode.data == e {
			index = i
			return
		}
		currentNode = currentNode.next
	}
	err = errors.New("未找到元素")
	return
}

// 查询前驱
func (l *LinkList) PriorElem(e ElemeType) (pe ElemeType, err error) {

	if l.length <= 1 {
		err = errors.New("数据结构元素不足，无法查询")
		return
	}

	if l.first.data == e {
		err = errors.New("首元素无前驱")
		return
	}

	currentNode := l.first
	for i := 1; i <= l.length-1; i++ {
		if currentNode.next.data == e {
			pe = currentNode.data
			return
		}
		currentNode = currentNode.next
	}

	err = errors.New("未找到传入元素")
	return
}

// 查询后继
func (l *LinkList) NextElem(e ElemeType) (ne ElemeType, err error) {
	if l.length <= 1 {
		err = errors.New("数据元素不足，无法查询")
		return
	}
	currentNode := l.first
	for i := 1; i <= l.length; i++ {
		if i == l.length {
			err = errors.New("最后一个元素无后继")
			return
		}
		if currentNode.data == e {
			ne = currentNode.next.data
			return
		}
		currentNode = currentNode.next
	}
	err = errors.New("未找到传入元素")
	return
}

// 清空
func (l *LinkList) Clear() {
	l.first = nil
	l.length = 0
}
