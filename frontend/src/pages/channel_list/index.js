import React, { useState, useEffect } from 'react'
import { useHistory } from "react-router-dom"
import { ListGroup } from 'react-bootstrap'

export default function ChannelListPage(props) {

  const history = useHistory()
  const [channelList, setChannelList] = useState([])

  useEffect(() => {
    fetch('/api/channel')
      .then(r => r.json())
      .then(r => setChannelList(r))
  }, [])

  const onClick = (channel) => {
    history.push('/chat/' + channel.id)
  }

  return (
    <>
      <p>Channel List</p>
      <ListGroup>
        {
          channelList.map((i, key) => {
            return (
              <ListGroup.Item
                style={{ cursor: 'pointer' }}
                key={key}
                onClick={() => onClick(i)}
              >{i.name}</ListGroup.Item>
            )
          })
        }
      </ListGroup>
    </>
  )
}
