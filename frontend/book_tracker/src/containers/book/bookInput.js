import React, { useState } from 'react'
import {
    Button,
    Container,
    Form,
    Input,
    Select,
    TextArea,
} from 'semantic-ui-react'
import SemanticDatepicker from 'react-semantic-ui-datepickers'
import API from 'api'
import { statusOptions } from 'utils/statusMapping'

const emptyForm = {
    title: '',
    author: '',
    status: '',
    description: '',
    startTime: '',
    endTime: '',
}

const BookInput = (props) => {
    const { mode, data, onClose } = props
    const [form, setForm] = useState(mode === 'new' ? emptyForm : data)

    const handleChange = (e, { name, value }) => {
        setForm((prevForm) => ({
            ...prevForm,
            [name]: value,
        }))
    }

    const _new = async (form) => {
        try {
            const res = await API.addBook(form)
            console.log('add book:', res)
        } catch (err) {
            console.error(err)
        }
    }
    const _edit = async (form) => {
        try {
            const res = await API.editBook(data.id, form)
            console.log('edit book:', res)
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

    return (
        <Container>
            <Form onSubmit={handleSubmit} >
                <Form.Group>
                    <Form.Field
                        width={16}
                        control={Input}
                        label='Title'
                        name={'title'}
                        value={form.title}
                        placeholder='Title of Book'
                        onChange={handleChange}
                        required
                    />
                    <Form.Field
                        width={8}
                        control={Input}
                        label='Author'
                        name={'author'}
                        value={form.author}
                        placeholder='Author of Book'
                        onChange={handleChange}
                        required
                    />
                    <Form.Field
                        width={8}
                        control={Select}
                        label='Status'
                        name={'status'}
                        value={form.status}
                        options={statusOptions}
                        placeholder='Status'
                        onChange={handleChange}
                        required
                    />
                </Form.Group>
                <Form.Group >
                    {/* <DatePicker
                        label={'startTime'}
                        selected={form.endTime}
                        onChange={(date) => handleDateChange('startTime', date)}
                    />
                    <DatePicker
                        selected={form.endTime}
                        onChange={(date) => handleDateChange('endTime', date)}
                    /> */}
                    <SemanticDatepicker name={'startTime'} label={'Start Time'} showToday={false} onChange={handleChange} />
                    <SemanticDatepicker name={'endTime'} label={'End Time'} showToday={false} onChange={handleChange} />
                </Form.Group>
                <Form.Group>
                    <Form.Field
                        width={16}
                        control={TextArea}
                        rows={8}
                        label='Description'
                        name={'description'}
                        value={form.description}
                        placeholder='Something about this book...'
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

export default BookInput