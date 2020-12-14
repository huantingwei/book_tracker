import React, { useEffect, useState } from 'react'
import {
    Button,
    Container,
    Form,
    Input,
    TextArea,
    Header,
    Icon
} from 'semantic-ui-react'
import API from 'api'

const emptyForm = {
    title: '',
    content: '',
}

const NoteInput = (props) => {
    const { mode, data, onClose } = props
    const [form, setForm] = useState(mode === 'new' ? emptyForm : data)
    const [bookid, setBookId] = useState(null)
    
    const handleChange = (e, { name, value }) => {
        setForm((prevForm) => ({
            ...prevForm,
            [name]: value,
        }))
    }

    const _new = async (form) => {
        try {
            const res = await API.addNote({
                ...form,
                createTime: new Date(),
                bookID: bookid
            })
            console.log('add note:', res)
        } catch (err) {
            console.error(err)
        }
    }

    const _edit = async (form) => {
        try {
            const res = await API.editNote(data.id, form)
            console.log('edit note:', res)
        } catch (err) {
            console.error(err)
        }
    }

    const handleSubmit = async () => {
        switch (mode) {
            case 'new':
                await _new(form)
                setForm(emptyForm)
                onClose()
                return
            case 'edit':
                await _edit(form)
                setForm(emptyForm)
                onClose()
                return
            default:
                setForm(emptyForm)
                onClose()
                return
        }

    }

    useEffect(() => {
        const search = window.location.search
        const params = new URLSearchParams(search)
        const id = params.get('bookid')
        setBookId(id)
    }, [])

    return (
        <Container>
            <Header as='h2'>
                <Icon name='space shuttle' />
                <Header.Content>
                    New Note
                    <Header.Subheader>
                        Inspire Yourself!
                    </Header.Subheader>
                </Header.Content>
            </Header>
            <Form onSubmit={handleSubmit} >
                <Form.Group>
                    <Form.Field
                        width={16}
                        control={Input}
                        label='Title'
                        name={'title'}
                        value={form.title}
                        placeholder='Title of Note'
                        onChange={handleChange}
                        required
                    />
                </Form.Group>
                <Form.Group>
                    <Form.Field
                        width={16}
                        control={TextArea}
                        rows={8}
                        label='Content'
                        name={'content'}
                        value={form.description}
                        placeholder='Input your note here...'
                        onChange={handleChange}
                        required
                    />
                </Form.Group>
                <Form.Group>
                    <Form.Field control={Button}>Save</Form.Field>
                </Form.Group>
            </Form>
        </Container>
    )
}

export default NoteInput