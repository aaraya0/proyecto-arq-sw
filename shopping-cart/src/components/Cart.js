import React, { useState } from "react";

import Cookies from 'universal-cookie';
import "./Cart.css";
const cookies = new Cookies();

async function getProductById(id){
return fetch("http://127.0.0.1:8090/product/" + id).then(response => response.json())
}

function gopath(path){
window.location = window.location.origin + path
}

async function getCartProducts(){
      let current= cookies.get('cart')
      let items=[]
      let a=current.split(',')
      for (let i = 0; i < a.length; i++){
      let item = a[i];
      let prod= item.split('=')
      let prodID=prod[0]
      let cant= prod[1]
      if (prodID!=0 && cant!="undefined" && cant>0){
      let product = await getProductById(prodID)
      product.quantity = cant;
items.push(product)
}
}
return items;
}

function showProducts(products){
return products.map((product) =>


      <div className="productoCard">
      <div obj={product} key={product.id} className="product"> </div>

      <div className="title">{product.title}</div>
      <img className="image" src={`img/${product.image}`}></img>
      <div className="price">{ "$" + product.base_price}</div>
      <div className="quantity">
      <div className="amount"> Cantidad: {product.quantity} </div>

      </div>
      <div className="subtotal"> Subtotal: ${product.quantity * product.base_price} </div>
      <button id={product.id} className="Xbutton" role="button" onClick={DeleteCartProduct}>X</button>
      </div>

) }

async function setCart(setter, setterTotal){
      let total = 0;
      await getCartProducts().then(response => {
      setter(response)
      response.forEach((item) => {
      total += item.base_price * item.quantity;
      });
      setterTotal(total)
      cookies.set("total", total)
})
}

async function DeleteCartProduct(e){
      let itemID = e.target.id
      let current= cookies.get("cart")
      let newCart=''
      let items=current.split(',')


        items.forEach(item =>{

        let comp= item.split('=')
        let compID=comp[0]
        let quantity= comp[1]

        if (itemID==compID){
        quantity--

        }
      newCart=`${newCart},${compID}=${quantity}`
})

cookies.set("cart", newCart, { path: '/' })
window.location.reload(false)
}

async function cartCheckout(){
gopath("/checkout")
}




function Cart(){

      const [cartProducts, setCartProducts] = useState([]);
      const [total, setTotal] = useState(0);
      if (cartProducts.length <= 0){
      setCart(setCartProducts, setTotal)
}


return (
      <div className="cart">
      { showProducts(cartProducts)}
      <div className="prods">
      </div>
      <div id="main">
      <div className="boton-checkout">
      <div className="tot"> Total: ${total} </div>
      <div><button id="check"  class="CObutton" role="button" onClick={cartCheckout}>Checkout</button></div>  
      </div>
      </div>
      </div>
);
}

export default Cart;
