import './App.css';
import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";
import ChatRoomPage from './pages/chat_room/index'
import ChannelListPage from './pages/channel_list/index'

function App() {
  return (
    <>
      <Router>
        <div style={{
          width: '50%',
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