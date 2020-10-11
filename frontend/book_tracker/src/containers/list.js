import React from 'react'
import uuid from 'react-uuid'
import { Button, Icon, Item, Label, Grid } from 'semantic-ui-react'

const ListComponent = (props) => {
    const { items, onItemClick } = props

    const handleItemOnClick = (id) => {
        onItemClick(id)
    }
    return (
        <Item.Group link divided>
            {items.map((item) => {

                <Item key={uuid()} onClick={() => handleItemOnClick(id)}>
                    <Item.Image src='/images/wireframe/image.png' />
                    <Item.Content>
                        <Item.Header as='a'>{item.title}</Item.Header>
                        <Item.Meta>
                            <span className='cinema'>{item.subtitle}</span>
                        </Item.Meta>
                        {/* <Button icon size="small" floated='right'>
                            <Icon name='delete' />
                        </Button>
                        <Button icon size="small" floated='right'>
                            <Icon name='edit' />
                        </Button> */}
                        <Item.Description>
                            <Grid columns={item.content.length}>{item.content.map((c) => {
                                <Grid.Row >
                                    <Grid.Column style={{ paddingLeft: 0 }}>
                                        {c}
                                    </Grid.Column>
                                </Grid.Row>
                            })}
                            </Grid>

                        </Item.Description>
                        {/* <Item.Extra>
                            <Label>Tag</Label>
                            <Label icon='globe' content='Additional Languages' />
                        </Item.Extra> */}
                    </Item.Content>
                </Item>
            })}
        </Item.Group>)
}

export default ListComponent