import React, { useState, useEffect } from 'react'
import { Container, Header } from 'semantic-ui-react'
import { useParams } from 'react-router-dom'
import API from 'api'

const Note = (props) => {
    const { id } = useParams()
    const [note, setNote] = useState({})


    useEffect(() => {
        async function _get() {
            try {
                const res = await API.getNote(id)
                setNote(res)
            }
            catch (err) {
                console.error(err)
            }
        }
        _get()
    }, [id])

    return (
        <Container text>
            <Header size='large'>{note.title}</Header>
            <p>
                {note.content}
            </p>
        </Container>
    )

}

export default Note