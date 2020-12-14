import React, { Fragment, useState, useEffect } from 'react'
import { Container, Header, Icon, Divider } from 'semantic-ui-react'
import { useParams } from 'react-router-dom'
import API from 'api'
import NoteList from 'containers/note/noteList'

const Book = (props) => {

    const { id } = useParams()
    const [book, setBook] = useState({})

    useEffect(() => {
        async function _get() {
            try {
                const res = await API.getBook(id)
                setBook(res)
            }
            catch (err) {
                console.error(err)
            }
        }
        _get()
    }, [id])

    return (
        <Fragment>
            <Container>
                <Header as='h2' icon textAlign='center' style={{ marginBottom: "2rem" }}>
                    <Icon name='book' circular />
                    <Header.Content>{book.title}</Header.Content>
                    <Header.Subheader>{book.author}</Header.Subheader>
                </Header>
                <Container text content={book.description} textAlign="center" />
                <Divider />
                <NoteList noteIDs={book.notes} bookData={book} />
            </Container>
        </Fragment>
    )
}

export default Book

