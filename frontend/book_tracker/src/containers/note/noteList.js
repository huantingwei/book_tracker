import React, { Fragment, useEffect, useState } from 'react'
import PropTypes from 'prop-types'
import { Container, Item, Button } from 'semantic-ui-react'
import ContainerLayout from 'components/containerLayout'
import API from 'api'
import { noteURL, newNoteURL } from 'utils/urls'


const fakeNotes = [
    {
        id: 1,
        title: 'first note',
        content: 'Hello World!',
        createTime: '2020-10-23',
    },
    {
        id: 2,
        title: 'note 2',
        content: 'What\'s the Future?',
        createTime: '2020-12-25',
    }
]
const NoteList = (props) => {

    const { noteIDs, bookData } = props
    const [noteList, setNoteList] = useState(fakeNotes)

    useEffect(() => {
        async function _enumerate(ids) {
            try {
                let notes = []
                for (var noteID of ids) {
                    const res = await API.getNote(noteID)
                    notes.push(res)
                }
                setNoteList(notes)
            } catch (err) {
                console.error(err)
            }
        }
        _enumerate(noteIDs)
    }, [noteIDs])

    const handleNoteOnClick = (noteID) => {
        window.open(noteURL + noteID)
    }

    const handleNewClick = () => {
        window.location = newNoteURL(bookData['id'])
    }

    return (
        <ContainerLayout>
            <Button size="medium" onClick={handleNewClick} style={{ float: 'right', margin: "1rem" }}>New Note</Button>
            {noteList.length === 0 ?
                <Container text textAlign="center" content='No Notes' /> :
                <Item.Group link divided>
                    {noteList.map((note) => {
                        return (
                            <Item key={note.id}>
                                <Item.Content>
                                    <Item.Header as='a' onClick={() => handleNoteOnClick(note.id)}>
                                        {note.title}
                                    </Item.Header>
                                    <Item.Meta>
                                        <span>{note.createTime}</span>
                                    </Item.Meta>
                                    <Item.Description>
                                        {note.content}
                                    </Item.Description>
                                </Item.Content>
                            </Item>
                        )
                    })}

                </Item.Group >
            }
        </ContainerLayout >
    )
}

NoteList.defaultProps = {
    noteIDs: [],
}

NoteList.prototype = {
    noteIDs: PropTypes.array,
}

export default NoteList