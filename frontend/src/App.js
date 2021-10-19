import './App.css';
import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";
import ChatRoomPage from './pages/chat_room/index'

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
            <Route path="/">
              <ChatRoomPage />
            </Route>
          </Switch>
        </div>
      </Router>
    </>
  );
}

export default App;