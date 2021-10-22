import './App.css';
import { useEffect } from 'react'
import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";
import ChatRoomPage from './pages/chat_room/index'
import ChannelListPage from './pages/channel_list/index'

function App() {

  useEffect(() => {
    if (localStorage.getItem('username') === null) {
      localStorage.setItem('username', prompt('Input Username'))
    }
  }, [])

  return (
    <>
      <Router>
        <div style={{
          width: '80%',
          margin: 'auto',
          marginTop: '20px',
        }}>

          <Switch>
            <Route path="/chat/:channel_id">
              <ChatRoomPage />
            </Route>
            <Route path="/">
              <ChannelListPage />
            </Route>
          </Switch>
        </div>
      </Router>
    </>
  );
}

export default App;