import React, { useState } from "react";
import Cookies from 'universal-cookie'
const cookies = new Cookies();

function getProductById(id){
    return fetch("http://127.0.0.1:8090/product/" + id).then(response => response.json())
  }

  
async function getCartProducts(){
    let current= cookies.get('cart')
    let items=[]
    let a=current.split(',')
    a.forEach(item =>{
        
        let prod= item.split('=')
        let prodID=prod[0]
       
        let cant= prod[1]
        if (prodID!=0 && cant!="undefined" ){
        let product = getProductById(prodID)
        product.quantity = cant;
        items.push(product)
      }
    })
let productos= items.map((product) =>
<ol key = { product.id } className="item">
<div className="title">{product.title}</div>
<div className="author">{product.author}</div>
<div className="price">{product.base_price}</div>
</ol>)
return (<div>{productos}</div>)
  }

  function Cart(){
 
getCartProducts()
  }
export default Cart;
