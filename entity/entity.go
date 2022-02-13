package entity

type Entity interface {
	id() uint
	Renderable
}
