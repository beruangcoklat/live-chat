import React, { useState, useEffect } from 'react';

export default function ChatRoomPage(props) {

  const [chatList, setChatList] = useState([{ message: "tes" }])

  useEffect(() => {
    const sse = new EventSource("http://localhost:8080/chat/1")

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
  }, [])

  return (
    <>
      {
        chatList.map(i => {
          return (
            <p>a {i.message}</p>
          )
        })
      }
    </>
  )
}
