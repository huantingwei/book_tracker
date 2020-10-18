import React, { useEffect, useState } from 'react'
import PropTypes from 'prop-types'
import { Container, Item, } from 'semantic-ui-react'
import API from 'api'


const NoteList = (props) => {

    const { noteIDs } = props
    const [noteList, setNoteList] = useState([])

    useEffect(() => {
        async function _enumerate() {
            let notes = []
            for (var noteID of noteIDs) {
                const res = await API.getNote(noteID)
                notes.push(res)
                console.log(res)
            }
            setNoteList(notes)
        }
        _enumerate()
    }, [noteIDs])

    const handleNoteOnClick = (noteID) => {
        console.log(noteID)
        window.open('http://localhost:3000/note/' + noteID)
    }

    return (
        <>
            <Container>
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
            </Container >
        </>
    )
}

NoteList.defaultProps = {
    noteIDs: [],
}

NoteList.prototype = {
    noteIDs: PropTypes.array,
}

export default NoteList