import React from 'react'

import { Container } from 'semantic-ui-react'

const ContainerLayout = (props) => {
    const { children } = props

    return (
        <Container style={{ width: "70%", margin: "3rem 0" }}>{children}</Container>
    )
}

export default ContainerLayout