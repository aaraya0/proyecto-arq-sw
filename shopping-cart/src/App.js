

import React from "react";
import './App.css';

class App extends React.Component {

	// Constructor
	constructor(props) {
		super(props);

		this.state = {
			items: [],
			DataisLoaded: false
		};
	}

	// ComponentDidMount is used to
	// execute the code
	componentDidMount() {
		fetch(
"http://localhost:8090/product")
			.then((res) => res.json())
			.then((json) => {
				this.setState({
					items: json,
					DataisLoaded: true
				});
			})
	}
	render() {
		const { DataisLoaded, items } = this.state;
		if (!DataisLoaded) return <div>
			<h1> Please wait... </h1> </div> ;

		return (
		<div className = "App">
			<h1> Productos </h1> {
				items.map((item) => (
				<ol key = { item.id } >
					Title: { item.title},
					Author: { item.author },
				  Price: { item.base_price }
					</ol>
				))
			}
		</div>
	);
}
}

export default App;