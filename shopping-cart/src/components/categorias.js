
import React, { useState } from "react";

import Cookies from 'universal-cookie';

import "./categorias.css"

 
const cookies = new Cookies();



async function getCategories(){
    return await fetch('http://127.0.0.1:8090/category').then(response => response.json())
  }

  async function getProductsByCategoryId(id){
    return await fetch('http://127.0.0.1:8090/products/' + id).then(response => response.json())
  } 
  
  async function getCategoryById(id){
    return await fetch('http://127.0.0.1:8090/category/' + id).then(response => response.json())
  }


  function productsByCategoryId(id, setter, categorySetter) {
    getProductsByCategoryId(id).then(response => {
        setter(response); cookies.set("category", id); getCategoryById(id).then(category => categorySetter(category))})
  }
  
  function showCategories(categories, setter, categorySetter) {
    return categories.map((category) => <a onClick={() => productsByCategoryId(category.id, setter, categorySetter)} obj={category} key={category.id}>{category.name}</a>)
  }
  function gopath(path){
    window.location = window.location.origin + path
  }
  
  
 
  
  function retry() {
    gopath("/home")
  }
  function Category(){
    const [categories, setCategories] = useState([])
   const [ products, setProducts] = useState([])
    const [category, setCategory] = useState("")
    if(categories.length <= 0){
        getCategories().then(response => setCategories(response))
      }

      return(
<div>
<div id="mySidenav" className="sidenav">

{categories.length > 0 ? showCategories(categories, setProducts, setCategory) : <a onClick={retry}> Algo salió mal. Presione para reintentar. </a>}
</div>

<div id="main">
{cookies.get("category") > 0 ? <a className="categoryFilter"> {category.name}  </a> : <a/>}
</div>
</div>
      )
  }
  export default Category;