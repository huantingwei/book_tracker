import React, { useState, useEffect } from 'react'
import { Container, Header } from 'semantic-ui-react'
import { useParams } from 'react-router-dom'
import API from 'api'

const fakeNote = {
    title: 'Hello World!',
    content: 'Fake Note\nThis course more about the theoretical level, they have Labs to practice, but not a lot. I will say it helps you to understand the basic of each services. And maybe give you some direction to prepare.'
}
const Note = (props) => {
    const { id } = useParams()
    const [note, setNote] = useState(fakeNote)


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