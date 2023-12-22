import React from 'react'
import {createRoot} from 'react-dom/client'
import App from './App'
import {ChakraProvider, extendBaseTheme, theme as chakraTheme} from "@chakra-ui/react";
import theme from "./theme.js";

const container = document.getElementById('root')

const root = createRoot(container)

root.render(
    <React.StrictMode>
        <ChakraProvider theme={theme}>
            <App/>
        </ChakraProvider>
    </React.StrictMode>
)
