import React, {Component} from 'react';
import MapContainer from './map/MapContainer.jsx';
import Home from './home/Home.jsx';
import ByCountry from './by-country/ByCountry.jsx';
import ByAppId from './by-appid/ByAppId.jsx';
import ByTime from './by-time/ByTime.jsx';

import {
  BrowserRouter as Router,
  Route,
  NavLink,
  Link
} from 'react-router-dom'

import 'react-dates/lib/css/_datepicker.css';

/**
* App main component
*/
class App extends Component{
  constructor(props){
   super(props);
   this.state = {
	 backendUrl: '//' + location.hostname + ':8001',
     activeTab: location.pathname,
     startDate: null,
     endDate: null
   };
 }
 setDates(startDate, endDate) {
   this.setState({startDate, endDate});
 }
  render(){
    return (
      <Router>
        <div>
          <nav className="navbar navbar-inverse navbar-fixed-top">
            <div className="container">
              <div className="navbar-header">
                <button type="button" className="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
                  <span className="sr-only">Toggle navigation</span>
                  <span className="icon-bar"></span>
                  <span className="icon-bar"></span>
                  <span className="icon-bar"></span>
                </button>
                <Link className="navbar-brand" to="/">Downloads</Link>
              </div>
              <div id="navbar" className="collapse navbar-collapse">
                <ul className="nav navbar-nav">
                  <li><NavLink activeClassName='active' to="/map">Map</NavLink></li>
                  <li><NavLink activeClassName='active' to="/countries">By country</NavLink></li>
                  <li><NavLink activeClassName='active' to="/appid">By app_id</NavLink></li>
                  <li><NavLink activeClassName='active' to="/time">By Time of the Day</NavLink></li>
                </ul>
              </div>
            </div>
        </nav>

        <div className="container">
          <Route exact path="/" component={Home} />
          <Route  path="/map" render={()=><MapContainer backendurl={this.state.backendUrl} setDates={this.setDates.bind(this)} {...this.state} />} />
          <Route  path="/countries" render={()=><ByCountry backendurl={this.state.backendUrl} setDates={this.setDates.bind(this)} {...this.state}/>}/>
          <Route  path="/appid" render={()=><ByAppId backendurl={this.state.backendUrl} setDates={this.setDates.bind(this)} {...this.state}/>} />
          <Route  path="/time" render={()=><ByTime backendurl={this.state.backendUrl} setDates={this.setDates.bind(this)} {...this.state} />} />
        </div>
      </div>
    </Router>

    )
  }

}

export default App
