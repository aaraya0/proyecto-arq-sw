
import React, { useState } from "react";

import Cookies from 'universal-cookie';

import "./categorias.css";


const cookies = new Cookies();

async function getCategories(){
return await fetch('http://127.0.0.1:8090/category').then(response => response.json())
}

async function getCategoryById(id){
return await fetch('http://127.0.0.1:8090/category/' + id).then(response => response.json())
}

function gopath(path){
window.location = window.location.origin + path
}
async function handleCat(){
        cookies.set("category", 0)
        window.location.reload()
}
function showCategories(categories) {
const cats=categories.map((category) => <a onClick={() => 
        showCategoryProds(category.id)} 
        obj={category} key={category.id}>{category.name} </a>)
return (<div>
        <div>{cats}</div>
        <div><button id="sidenav2"  class="sidenav2" role="button" onClick={handleCat}>Todos los libros</button></div> 
</div>)
}

async function getProductsByCategoryId(id){

return await fetch('http://127.0.0.1:8090/products/' + id).then(response => response.json())

}
async function showCategoryProds(id){
        cookies.set("category", id)
        window.location.reload()
       /* let productos= await getProductsByCategoryId(id)
        return productos.map((item)=> (

        <div key = { item.id } className="item">

        <div id="titulo">{ item.title}</div>
        <div id="autor"> { item.author }</div>	
        <div id="precio">${ item.base_price }</div> 
        </div>
        ))*/
}

function Category() {

            const [categories, setCategories] = useState([])
            
            if(categories.length <= 0){
            getCategories().then(response => setCategories(response))
            }
                return (
                <div>
                <div id="mySidenav" className="sidenav">
                {categories.length > 0 ? showCategories(categories) : <a > Volver a cargar </a>}
                </div>
                </div>
                );
}

export default Category;