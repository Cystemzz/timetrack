import {Button, Card, CardBody, CardHeader, Heading, Spinner, Stack, Text, useColorMode} from "@chakra-ui/react";
import {useEffect, useState} from "react";
import {GetWorkspaces} from "../../wailsjs/go/main/App.js";

const LOADING = 0
const NONE_LOADED = 1
const LOADED = 2

export default function Home() {
    const [stage, setStage] = useState(LOADING);
    const [workspaces, setWorkSpaces] = useState([]);

    useEffect(() => {
        GetWorkspaces().then(w => {
            if (w.length === 0) {
                setStage(NONE_LOADED);
                return
            }

            setWorkSpaces(w);
            setStage(LOADED);
            console.log(w)
        })
    }, []);

    if (stage === LOADING) {
        return (
            <div className="flex h-screen flex-col justify-center items-center">
                <Spinner size='xl' />
            </div>
        )
    }

    if (stage === NONE_LOADED) {
        return (
            <div className="flex h-screen flex-col justify-center items-center space-y-3">
                <p>You have no recent workspaces</p>
                <div className="space-x-3">
                    <Button colorScheme="blue">Create File</Button>
                    <Button colorScheme="green">Load File</Button>
                </div>
            </div>
        )
    }

    // TODO: Finish making a list of workspaces
    return (
        <div>
            {workspaces.map(w => (
                <Card className="m-4" key={w.Location} variant="outline">
                    <CardHeader>
                        <Heading size='md'>{w.Name}</Heading>
                    </CardHeader>
                    <CardBody>
                        <Text>{w.Location}</Text>
                    </CardBody>
                </Card>
            ))}
        </div>
    )
}
