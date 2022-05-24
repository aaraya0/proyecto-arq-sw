package model

type OrderDet struct {
	Prods Products
	Cost  int    `gorm:"type:int;not null"`
}

type OrderDets []OrderDet


//ordenes de compra individuales cada vez q se entra a la pagina 
//y se agrega algo al carrito, luego orden final cuando se presiona el boton para comprar??
//el producto tiene el precio individual de cada articulo, la orden detalle tiene el costo total por orden parcial y 
//la orden final tiene el costo suma de todas las ordenes parciales