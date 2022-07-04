package main

import "fmt"

/*
Стратегия — это поведенческий паттерн проектирования, который определяет
семейство схожих алгоритмов и помещает каждый из них в собственный класс,
после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

Применять:
-Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
-Когда у вас есть множество похожих объектов, отличающихся только некоторым поведением.
-Когда вы не хотите обнажать детали реализации алгоритмов для других объектов.
-Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора.
-Каждая ветка такого оператора представляет собой вариацию алгоритма.

Плюсы:
-Горячая замена алгоритмов на лету.
-Изолирует код и данные алгоритмов от остальных классов.
-Уход от наследования к делегированию.
-Реализует принцип открытости/закрытости.

Минусы:
-Усложняет программу за счёт дополнительных сущностей.
-Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

// Интерфейс стратегии
type evictionAlgo interface {
	evict(c *cache)
}

type cache struct {
	storage      map[string]string
	evictionAlgo evictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e evictionAlgo) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *cache) setEvictionAlgo(e evictionAlgo) {
	c.evictionAlgo = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

// Конкретная стратегия
type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("Evicting by fifo strtegy")
}

// Конкретная стратегия
type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("Evicting by lfu strtegy")
}

// Конкретная стратегия
type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strtegy")
}

func main() {
	lfu := &lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")

}
