import {HashRouter, Route, Routes} from "react-router-dom";
import "./output.css"
import Home from "./pages/Home.jsx";

function App() {
    return (
        <div id="App">
            <HashRouter>
                <Routes>
                    <Route path="/" element={<Home />} />
                </Routes>
            </HashRouter>
        </div>
    )
}

export default App
