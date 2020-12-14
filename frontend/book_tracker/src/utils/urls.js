export const bookListURL = 'http://localhost:3000/book'

export const bookURL = 'http://localhost:3000/book/'

export const newBookURL = 'http://localhost:3000/newBook'

export const noteURL = 'http://localhost:3000/note/'

export const newNoteURL = (bookID) => {
    if (bookID !== -1) {
        return `http://localhost:3000/newNote?bookid=${bookID}`
    } else return `http://localhost:3000/newNote`
}