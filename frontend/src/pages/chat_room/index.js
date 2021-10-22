import React, { useState, useEffect } from 'react'
import { useParams } from "react-router-dom"
import { Badge, Container, Row, Col, Form, Button } from 'react-bootstrap'

export default function ChatRoomPage(props) {

  const params = useParams()
  const [chatList, setChatList] = useState([])
  const [msg, setMsg] = useState('')

  const unixToDate = (unix) => {
    var date = new Date(unix)
    var hours = date.getHours();
    var minutes = "0" + date.getMinutes();
    var seconds = "0" + date.getSeconds();
    return hours + ':' + minutes.substr(-2) + ':' + seconds.substr(-2);
  }

  useEffect(() => {
    fetch(`/api/chat?channel_id=${params.channel_id}&channel_id&limit=10`)
      .then(r => r.json())
      .then(r => setChatList(r.reverse()))

    const sse = new EventSource(`/api/chat/${params.channel_id}`)

    sse.onmessage = e => {
      const event = JSON.parse(e.data)
      setChatList(oldList => [...oldList, event])
    }

    sse.onerror = () => {
      sse.close()
    }

    return () => {
      sse.close()
    }
  }, [params.channel_id])

  const onSubmit = function (e) {
    e.preventDefault()

    const payload = {
      channel_id: params.channel_id,
      sender: localStorage.getItem('username'),
      message: msg,
    }

    fetch('/api/chat', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
    })
      .then(() => {
        setMsg('')
      })
  }

  return (
    <>
      <Container>
        {
          chatList.map((i, key) => {
            return (
              <Row key={key} style={{ paddingTop: '10px' }}>
                <Col xs={2}>
                  <label style={{ fontSize: '13px', marginRight: '10px', color: '#777' }}>
                    {unixToDate(i.created_at)}
                  </label>
                  <Badge bg="primary">{i.sender}</Badge>
                </Col>
                <Col>{i.message}</Col>
              </Row>
            )
          })
        }
      </Container>

      <br />
      <Form required onSubmit={onSubmit}>
        <Form.Group className="mb-3">
          <Form.Control
            type="text"
            placeholder="Enter message"
            required
            value={msg}
            onChange={e => setMsg(e.target.value)}
          />
        </Form.Group>

        <Button variant="primary" type="submit">
          Send
        </Button>
      </Form>
    </>
  )
}
