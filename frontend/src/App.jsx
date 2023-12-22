import {HashRouter, Route, Routes} from "react-router-dom";
import Home from "./pages/Home.jsx";
import {useEffect} from "react";
import {useColorMode} from "@chakra-ui/react";

function App() {
    const { colorMode, toggleColorMode } = useColorMode();

    // Forcing dark mode due to chakra being really weird about it
    useEffect(() => {
        if (colorMode === 'light') {
            toggleColorMode();
        }
    }, []);

    return (
        <div id="App">
            <HashRouter>
                <Routes>
                    <Route path="/" element={ <Home /> } />
                </Routes>
            </HashRouter>
        </div>
    )
}

export default App
