package main

import "fmt"

/*
Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

Плюсы:
-Позволяет создавать продукты пошагово.
-Позволяет использовать один и тот же код для создания различных продуктов.
-Изолирует сложный код сборки продукта от его основной бизнес-логики.

Минусы:
-Усложняет код программы из-за введения дополнительных классов.
-Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.
*/

type car struct {
	engine string
	tires  string
	seat   int
}

type iBuilder interface {
	setEngine()
	setTires()
	setSeat()
	getCar() car
}

func getBuilder(builderType string) iBuilder {
	if builderType == "sport" {
		return newSportBuilder()
	}

	if builderType == "bus" {
		return newBusBuilder()
	}
	return nil
}

type sportBuilder struct {
	engine string
	tires  string
	seat   int
}

func newSportBuilder() *sportBuilder {
	return &sportBuilder{}
}

func (b *sportBuilder) setEngine() {
	b.engine = "Sport engine"
}

func (b *sportBuilder) setTires() {
	b.tires = "Sport tires"
}

func (b *sportBuilder) setSeat() {
	b.seat = 2
}

func (b *sportBuilder) getCar() car {
	return car{
		engine: b.engine,
		tires:  b.tires,
		seat:   b.seat,
	}
}

type busBuilder struct {
	engine string
	tires  string
	seat   int
}

func newBusBuilder() *busBuilder {
	return &busBuilder{}
}

func (b *busBuilder) setEngine() {
	b.engine = "Bus engine"
}

func (b *busBuilder) setTires() {
	b.tires = "Bus tires"
}

func (b *busBuilder) setSeat() {
	b.seat = 45
}

func (b *busBuilder) getCar() car {
	return car{
		engine: b.engine,
		tires:  b.tires,
		seat:   b.seat,
	}
}

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildCar() car {
	d.builder.setEngine()
	d.builder.setTires()
	d.builder.setSeat()
	return d.builder.getCar()
}

func main() {
	sportBuilder := getBuilder("sport")
	busBuilder := getBuilder("bus")

	director := newDirector(sportBuilder)
	sportCar := director.buildCar()

	fmt.Printf("Sportcars engine Type: %s\n", sportCar.engine)
	fmt.Printf("Sportcars tires Type: %s\n", sportCar.tires)
	fmt.Printf("Sportcars seats Num: %d\n", sportCar.seat)

	director.setBuilder(busBuilder)
	bustCar := director.buildCar()

	fmt.Printf("\nBus engine Type: %s\n", bustCar.engine)
	fmt.Printf("Bus tires Type: %s\n", bustCar.tires)
	fmt.Printf("Bus seats Floor: %d\n", bustCar.seat)
}
