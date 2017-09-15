import React, {Component} from 'react';
import {Map, InfoWindow, Marker, GoogleApiWrapper} from 'google-maps-react';
import { DateRangePicker, SingleDatePicker, DayPickerRangeController, isInclusivelyBeforeDay } from 'react-dates';
import axios from 'axios';
import moment from 'moment';
import 'react-dates/lib/css/_datepicker.css';

/**
* MapContainer component: includes a DateRangePicker and a Google Maps component with markers
*/
export class MapContainer extends Component {
  constructor(){
    super();
    this.state = {
      locations : []
    };
  }
  getDownloads() {
    let url = this.props.backendurl + '/downloads';
    this.props.endDate.set({hour:23,minute:59,second:59,millisecond:59});
    let params = {
      startDate: this.props.startDate.format(),
      endDate: this.props.endDate.format()
    };
    axios.get(url, {params})
      .then((response) => {
        let data = response.data;
        this.setState({locations: data});
      })
  }
  onMarkerClick(props, marker, e) {
    this.setState({
      selectedPlace: props,
      activeMarker: marker,
      showingInfoWindow: true
    });
  }

  render() {
      return (
        <div style={{position:'relative'}}>
          <DateRangePicker
              startDate={this.props.startDate}
              endDate={this.props.endDate}
              isOutsideRange={day => !isInclusivelyBeforeDay(day, moment())}
              minimumNights={0}
              onDatesChange={({ startDate, endDate }) => { this.props.setDates(startDate, endDate) }}
              focusedInput={this.state.focusedInput}
              onFocusChange={focusedInput => this.setState({ focusedInput })}
            />

          <button type="button" className="btn btn-primary" style={{margin: '20px'}}
            onClick={this.getDownloads.bind(this)}
             disabled={!this.props.startDate || !this.props.endDate}>
            Get downloads
          </button>

          <Map
            google={this.props.google}
            zoom={3}
            className={'map'}
            style={{height: '500px'}}
            initialCenter={{
              lat: 45.464161,
              lng: 9.190336
            }}
            clickableIcons={false}
          >
            {
              this.state.locations.map(location=> {
                return  <Marker
                        onClick={this.onMarkerClick.bind(this)}
                        key={location.id}
                        name={moment(location.downloaded_at).format('M/D/YYYY H:m')}
                        position={{lat: location.latitude, lng: location.longitude}} />
              })
            }
            { this.state.activeMarker &&
              <InfoWindow
                marker={this.state.activeMarker}
                visible={this.state.showingInfoWindow}>
                  <div>
                    {this.state.selectedPlace.name}
                  </div>
              </InfoWindow>
            }

          </Map>
        </div>
      );
    }
}


export default GoogleApiWrapper({
  apiKey: ('AIzaSyC0qX44R_iP1eY59bhdnsyw_BHO3-OPB5o')
})(MapContainer)
