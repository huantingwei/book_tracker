import React, { Fragment, useEffect, useState, useReducer } from 'react'
import PropTypes from 'prop-types'
import moment from 'moment'
import ContainerLayout from 'components/containerLayout'
import { Container, Item, Grid, Button, Popup, Modal, Header } from 'semantic-ui-react'
import API from 'api'
import { bookURL, newBookURL } from 'utils/urls'
import { mapping } from 'utils/statusMapping'
import BookInput from './bookInput'

const fakeBooks = [
    {
        id: 1,
        title: 'book R',
        author: 't1',
        startTime: '2020-10-01',
        endTime: '2020-10-10',
        status: 'Reading',
    },
    {
        id: 2,
        title: 'book Z',
        author: 't1',
        startTime: '2020-02-13',
        endTime: '2020-03-10',
        status: 'Finished',
    },
    {
        id: 3,
        title: 'book Y',
        author: 't1',
        startTime: '2020-02-13',
        endTime: '2020-03-10',
        status: 'Finished',
    },
    {
        id: 4,
        title: 'book X',
        author: 't1',
        startTime: '2020-02-13',
        endTime: '2020-03-10',
        status: 'Finished',
    },
    {
        id: 5,
        title: 'book D',
        author: 't1',
        startTime: '2020-02-13',
        endTime: '2020-03-10',
        status: 'Finished',
    },
    {
        id: 6,
        title: 'Book 1',
        author: 'Author a',
        status: 0,
        startTime: '2020-01-01',
        endTime: '2020-03-01',
    },
    {
        id: 7,
        title: 'Book 2',
        author: 'Author b',
        status: 1,
        startTime: '2020-02-02',
        endTime: '2020-05-05',
    },
    {
        id: 8,
        title: 'Book 3',
        author: 'Author c',
        status: 0,
        startTime: '2020-06-11',
        endTime: '2020-07-19',
    },
    {
        id: 9,
        title: 'Book 4',
        author: 'Author d',
        status: 2,
        startTime: '2020-10-20',
        endTime: '2020-11-12',
    }
]

function reducer(state, action) {
    switch (action.type) {
        case 'EDIT':
            return {
                dialogOpen: true,
                dialogSize: 'small',
                bookData: action.bookData,
                dialogAction: 'edit',
                dialogHeader: 'Edit Book',
            }
        case 'DELETE':
            return {
                dialogOpen: true,
                dialogSize: 'mini',
                bookData: action.bookData,
                dialogAction: 'delete',
                dialogHeader: 'WARNING',
            }
        case 'CLOSE_DIALOG':
            return { dialogOpen: false }
        default:
            throw new Error()
    }
}


const BookList = (props) => {
    const [books, setBooks] = useState(fakeBooks)
    const [dialogState, dispatch] = useReducer(reducer, {
        dialogOpen: false,
        dialogSize: undefined,
        dialogAction: undefined,
        dialogHeader: '',
        bookData: undefined
    })
    const { dialogOpen, dialogSize, bookData, dialogAction, dialogHeader, refresh } = dialogState

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
    }, [refresh])

    const handleBookClick = (bookID) => {
        window.location = bookURL + bookID
    }

    const handleNewClick = () => {
        window.location = newBookURL
    }

    const handleDeleteClick = () => {
        //delete book
        console.log('delete')
    }

    return (
        <ContainerLayout>
            <Button size="medium" onClick={handleNewClick} style={{ float: 'right', margin: "1rem" }}>New Book</Button>
            <Item.Group divided>
                {books.map((book) => {
                    return (
                        <Item key={book.id} >
                            <Item.Content>
                                <Item.Header as='a' onClick={() => handleBookClick(book.id)}>
                                    {book.title}
                                </Item.Header>
                                <Item.Description>
                                    <Grid>
                                        <Grid.Row>
                                            <Grid.Column computer={4} tablet={16} mobile={16}>
                                                {book.author}
                                            </Grid.Column>

                                            <Grid.Column computer={4} tablet={8} mobile={16}>
                                                {mapping[book.status]}
                                            </Grid.Column>
                                            <Grid.Column textAlign="left" computer={6} tablet={8} mobile={16}>
                                                Start Time: &nbsp;&nbsp;
                                                    {book.startTime === '' ? 'N/A' : moment(book.startTime).format('yyyy-MM-DD')} <br></br>
                                                    End Time:  &nbsp;&nbsp;
                                                    {book.endTime === '' ? 'N/A' : moment(book.endTime).format('yyyy-MM-DD')}
                                            </Grid.Column >
                                            <Grid.Column computer={2} tablet={16} mobile={16}>
                                                <Button size="mini" icon='edit'
                                                    onClick={() => dispatch({ type: 'EDIT', bookData: book })} />
                                                <Button size="mini" icon='delete'
                                                    onClick={() => dispatch({ type: 'DELETE', bookData: book })} />
                                            </Grid.Column >
                                        </Grid.Row>
                                    </Grid>
                                </Item.Description>
                            </Item.Content>
                        </Item>
                    )
                })}

            </Item.Group>

            <Modal
                closeIcon
                dimmer={'blurring'}
                open={dialogOpen}
                size={dialogSize}
                onClose={() => dispatch({ type: 'CLOSE_DIALOG' })}
            >
                <Modal.Header>{dialogHeader}</Modal.Header>
                {dialogAction === 'edit' ?
                    <Modal.Content>
                        <BookInput mode='edit' data={bookData} onClose={() => dispatch({ type: 'CLOSE_DIALOG' })} />
                    </Modal.Content>
                    :
                    <Fragment>
                        <Modal.Content>
                            <p>Are you sure you want to delete this book?</p>
                        </Modal.Content>
                        <Modal.Actions>
                            <Button size="small" floated="right" onClick={() => dispatch({ type: 'CLOSE_DIALOG' })}>Cancel</Button>
                            <Button size="small" floated="right" negative onClick={handleDeleteClick}>Confirm</Button>
                        </Modal.Actions>
                    </Fragment>
                }
            </Modal>
        </ContainerLayout>
    )
}

BookList.defaultProps = {
    books: fakeBooks,
}

BookList.prototype = {
    books: PropTypes.array,
}
export default BookList