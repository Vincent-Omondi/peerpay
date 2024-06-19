import React from 'react';
import { Route, Switch } from 'react-router-dom';
import Login from './components/Auth/Login';
import Signup from './components/Auth/Signup';
import Profile from './components/Profile/Profile';
import EditProfile from './components/Profile/EditProfile';
import BuyShares from './components/Shares/BuyShares';
import SellShares from './components/Shares/SellShares';
import AdminPanel from './components/Admin/AdminPanel';

function App() {
  return (
    <div className="App">
      <Switch>
        <Route path="/login" component={Login} />
        <Route path="/signup" component={Signup} />
        <Route path="/profile" component={Profile} />
        <Route path="/edit-profile" component={EditProfile} />
        <Route path="/buy-shares" component={BuyShares} />
        <Route path="/sell-shares" component={SellShares} />
        <Route path="/admin" component={AdminPanel} />
      </Switch>
    </div>
  );
}

export default App;
