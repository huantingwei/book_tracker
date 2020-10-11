import React, { useState } from 'react'
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link
} from "react-router-dom"
import { Menu } from 'semantic-ui-react'
import BookList from 'containers/bookList'
import NoteList from 'containers/noteList'
import Note from 'containers/note'
import Book from 'containers/book'

const Navbar = (props) => {
    const [activeItem, setActiveItem] = useState(null)
    const handleItemClick = ({ name }) => setActiveItem(name)
    return (
        <Router>
            <Menu pointing secondary>
                <Menu.Item as={Link} to='/'>
                    <h3 className="ui header">BookTracker</h3>
                </Menu.Item>
                <Menu.Item as={Link} to='/book' name='book' active={activeItem === 'book'} onClick={handleItemClick}>
                    <h5 className="ui header">Book</h5>
                </Menu.Item>
                <Menu.Item as={Link} to='/note' name='note' active={activeItem === 'note'} onClick={handleItemClick}>
                    <h5 className="ui header">Note</h5>
                </Menu.Item>
            </Menu>
            <Switch>
                <Route exact path="/">
                    <div>Home Page</div>
                </Route>
                <Route exact path="/book">
                    <BookList />
                </Route>
                <Route exact path="/note">
                    <NoteList />
                </Route>
                <Route path="/note/:id" children={<Note />} />
                <Route path="/book/:id" children={<Book />} />
            </Switch>
        </Router>
    )
}

export default Navbar