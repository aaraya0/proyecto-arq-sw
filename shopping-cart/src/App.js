//import logo from './logo.svg';
import './App.css';
import List from "./components/List";
import Create from "./components/Create";
import {BrowserRouter as Router} from "react-router-dom";
function App() {
  return (

    <Router>
  
    <div className="container">
      
        <List ></List>
      
         <Create ></Create>
      
 
    </div>
    </Router>
  );
}

export default App;
