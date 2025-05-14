package ALSS

import "fmt"

type list struct {
	len  int
	root *node
}

type node struct {
	next  *node
	value *agent
}

// newList создаёт пустой список list
func newList() *list {
	return &list{0, nil}
}

// add добавляет агента в конец списка
func (l *list) add(a *agent) {
	//увеличиваем счётчик длинны
	l.len++
	//создаём новую ноду
	newNode := &node{
		nil,
		a,
	}
	//если корень пустой, то новая нода - новый корень
	if l.root == nil {
		l.root = newNode
		return
	}
	//если корень полон, проходимся по всем остальным узлам в поиске
	//nil - конца списка, для добавления новый ноды туда.
	current := l.root
	for current.next != nil {
		current = current.next
	}
	//нашли конец - добавляем новую ноду сюда!
	current.next = newNode
}

// del ищет ноду с указанным агентом и удаляет её.
// Если агента нет, выдаёт ошибку
func (l *list) del(a *agent) error {
	//проверяем первый, корневой эллемент
	var last *node = l.root
	if last.value == a {
		l.len--
		l.root = l.root.next
		return nil
	}

	//выбираем текущий элемент для проверки
	var current *node = l.root.next

	//и проверяем все последующие элементы на вхождение
	for current != nil && current.value == a {
		last = current
		current = current.next
	}

	//если дошли до конца - выдаём ошибку. Значит агента никогда не было в списке
	//что странно...
	if current == nil {
		//в идеальном мире это никогда не отрабатывает!
		return fmt.Errorf("Cant remove value. Value not inside list!")
	}

	//если значение текущего элемента
	if current.value == a {
		l.len--
		last.next = current.next
	}
	return nil
}

// addAfter добавляет агента a сразу после указанного агента b
func (l *list) addAfter(base, new *agent) error {
	newNode := &node{
		next:  nil,
		value: new,
	}
	_ = newNode

	var current *node = l.root

	for current != nil && current.value == base {
		current = current.next
	}

	//если дошли до конца - выдаём ошибку. Значит агента никогда не было в списке
	//что странно...
	if current == nil {
		//в идеальном мире это никогда не отрабатывает!
		return fmt.Errorf("Cant remove value. Value not inside list!")
	}
	//если значение текущего элемента искомая база, то добавляем новую ноду
	if current.value == base {
		l.len++
		newNode.next = current.next
		current.next = newNode
	}
	return nil
}
