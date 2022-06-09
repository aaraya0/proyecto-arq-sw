import React, { useState } from "react";
import Cookies from 'universal-cookie';
import "./Cart.css";
const cookies = new Cookies();

 async function getProductById(id){
    return fetch("http://127.0.0.1:8090/product/" + id).then(response => response.json())
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
        if (prodID!=0 && cant!="undefined" ){
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
  <div className="price">{ "$" + product.base_price}</div>
  <div className="quantity">
   <div className="amount"> Cantidad: {product.quantity} </div>
   
   </div>
   <div className="subtotal"> Subtotal: ${product.quantity * product.base_price} </div>
   </div>
  


    )
       }


       async function setCart(setter, setterTotal){
        let total = 0;
        await getCartProducts().then(response => {
          setter(response)
          response.forEach((item) => {
            total += item.base_price * item.quantity;
          });
          setterTotal(total)
        })
      }



  function Cart(){
 
  const [cartProducts, setCartProducts] = useState([]);
  const [total, setTotal] = useState(0);
  if (cartProducts.length <= 0){
    setCart(setCartProducts, setTotal)
  }
  const renderEmptyCart = (
    <a className="empty-cart">El carrito esta vacio</a>
  )

  return (
    <div className="cart">
  

<div className="prods">
      {cookies.get("cart") ? showProducts(cartProducts) : renderEmptyCart}
      </div>

      <div className="boton-checkout">
      <div className="tot"> Total: ${total} </div>
      <div><button id="check" >Checkout</button></div>  
      </div>
    </div>
  );


  }
export default Cart;
