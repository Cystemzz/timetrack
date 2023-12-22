import {extendTheme, withDefaultColorScheme} from "@chakra-ui/react";

const theme = extendTheme({
        initialColorMode: 'dark',
        useSystemColorMode: false
    },
    withDefaultColorScheme({
        colorScheme: 'green',
    })
)

export default theme;