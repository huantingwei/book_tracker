import React, { useEffect, useState } from 'react'
import PropTypes from 'prop-types'
import { Container, Item, Grid } from 'semantic-ui-react'
import API from 'api'

const books = [
    {
        id: 1,
        title: 'bookA',
        startTime: '2020-10-01',
        endTime: '2020-10-10',
        status: 'Reading',
    },
    {
        id: 2,
        title: 'bookB',
        startTime: '2020-02-13',
        endTime: '2020-03-10',
        status: 'Finished',
    },
    {
        id: 3,
        title: 'bookB',
        startTime: '2020-02-13',
        endTime: '2020-03-10',
        status: 'Finished',
    },
    {
        id: 4,
        title: 'bookB',
        startTime: '2020-02-13',
        endTime: '2020-03-10',
        status: 'Finished',
    },
    {
        id: 5,
        title: 'bookB',
        startTime: '2020-02-13',
        endTime: '2020-03-10',
        status: 'Finished',
    }
]


const BookList = () => {
    const [books, setBooks] = useState([])

    useEffect(() => {
        async function _enumerate() {
            try {
                const res = await API.getBookList()
                setBooks(res)
            } catch (err) {
                console.error(err)
            }
        }
        _enumerate()
    }, [])

    const handleBookOnClick = (bookID) => {
        window.open('http://localhost:3000/book/' + bookID)
    }

    return (
        <Container style={{ width: "70%" }}>
            <Item.Group divided>
                {books.map((book) => {
                    return (
                        <Item key={book.id} >
                            {/* <Item.Image src={imageURL} /> */}
                            <Item.Content>
                                <Item.Header as='a' onClick={() => handleBookOnClick(book.id)}>
                                    {book.title}
                                </Item.Header>
                                <Item.Meta>
                                    <span className='cinema'>{book.author}</span>
                                </Item.Meta>
                                {/* <Button icon size="small" floated='right'>
                                    <Icon name='delete' />
                                </Button>
                                <Button icon size="small" floated='right'>
                                    <Icon name='edit' />
                                </Button> */}
                                <Item.Description>
                                    <Grid columns={3} divided>
                                        <Grid.Row >
                                            <Grid.Column>
                                                Status: {book.status}
                                            </Grid.Column>
                                            <Grid.Column>
                                                End Time: {book.endTime}
                                            </Grid.Column>
                                            <Grid.Column>
                                                Start Time: {book.startTime}
                                            </Grid.Column >
                                        </Grid.Row>
                                    </Grid>
                                </Item.Description>
                                {/* <Item.Extra>
                                    <Label>Tag</Label>
                                    <Label icon='globe' content='Additional Languages' />
                                </Item.Extra> */}
                            </Item.Content>
                        </Item>
                    )
                })}

            </Item.Group >
        </Container >
    )
}

BookList.defaultProps = {
    books: books,
}

BookList.prototype = {
    books: PropTypes.array,
}
export default BookList