import React from "react";
import './Home.css';
import Category from "./categorias"
import Cookies from 'universal-cookie';

const cookies = new Cookies();


const cookiesCartKey= 'cart';
cookies.set(cookiesCartKey, `0=0`,{ path: '/' } )
async function AddToCart(e){

			let itemID = e.target.id
			let current= cookies.get(cookiesCartKey)
			let newCart=''
			let items=current.split(',')
			let found= false
				items.forEach(item =>{

				let comp= item.split('=')
				let compID=comp[0]
				let quantity= comp[1]
				if (itemID==compID){
				quantity++
				found=true
				}
				newCart=`${newCart},${compID}=${quantity}`
				})
			if (!found){
			newCart=`${newCart},${itemID}=1`
			}
			cookies.set(cookiesCartKey, newCart, { path: '/' } )
}

function search(){
		let input, filter, a, i, selection;
		input = document.getElementById("search");
		selection= document.getElementById("searchOp");
		filter = input.value.toUpperCase();
		a = document.getElementsByClassName("item");
			if (selection.selectedOptions[0].value=="t"){
				for (i = 0; i < a.length; i++) {
				let txtValue = a[i].children[0].textContent || a[i].children[0].innerText;
				if (txtValue.toUpperCase().indexOf(filter) > -1) {
				a[i].style.display = "inherit";
			} else {
				a[i].style.display = "none";
			}}
			if(input.value.toUpperCase().length <= 0){
				for(i = 0; i < a.length; i++){
				a[i].style.display = "inherit";
			}}}
			if (selection.selectedOptions[0].value=="a"){
			for (i = 0; i < a.length; i++) {
				let txtValue = a[i].children[1].textContent || a[i].children[1].innerText;
			if (txtValue.toUpperCase().indexOf(filter) > -1) {
			a[i].style.display = "inherit";
			} else {
			a[i].style.display = "none";
			}}
			if(input.value.toUpperCase().length <= 0){
				for(i = 0; i < a.length; i++){
				a[i].style.display = "inherit";
			}}}


}

        

class Home extends React.Component {


constructor(props) {
super(props);
this.state = {
items: [],
/*categories:[],*/
DataisLoaded: false

};
}

componentDidMount() {
	fetch(
	"http://localhost:8090/product")
	.then((res) => res.json())
	.then((json) => {
	this.setState({
	items: json,
	DataisLoaded: true
	})});
	/*fetch(
		"http://localhost:8090/category")
		.then((res) => res.json())
		.then((json) => {
		this.setState({
		categories: json
		});
})*/}

render() {
const { DataisLoaded, items/*, categories*/ } = this.state;
	if (!DataisLoaded) return <div>
	<h1> Please wait... </h1> </div> ;


const producto= items.map((item) => (

	<div id = { item.id } className="item">
		<div id="titulo">{ item.title}</div>
		<div id="autor"> { item.author }</div>	
		<img id="imagen" src={`img/${item.image}`}></img>
		<div id="precio">${ item.base_price }</div>
		
		
		<div><button id={ item.id } class="button-50" role="button" onClick={AddToCart}>Agregar 🛒</button></div>
	</div>

))
/*const categoria= categories.map((cats)=>(

	<div id="categs" value={cats.id} onClick={showProds}>{cats.name}</div>
	
)
)*/







		return (
			<div >
			<Category/>
			<div className="form">
			<div className="seleccion">
			<select name="busqueda" id="searchOp">
			<option value="t">Buscar por Titulo</option>
			<option value="a">Buscar por Autor</option>
			</select>
			</div>
			<div className="buscar">
			<input type="text" id="search" placeholder="Buscar libros" onChange={search}/>
			</div>
			</div>
			<div className="product">{producto}</div>
			</div>
			)}
}
export default Home;

/*<select name="cats" id="filter" onChange={filterCat}>
				{categoria}
			</select>*/