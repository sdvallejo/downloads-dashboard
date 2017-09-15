import React, {Component} from 'react';
import _ from 'lodash';
import { DateRangePicker, SingleDatePicker, DayPickerRangeController, isInclusivelyBeforeDay } from 'react-dates';
import axios from 'axios';
import moment from 'moment';
import 'react-dates/lib/css/_datepicker.css';

/**
* Countries component: info about downloads by country
*/
class ByCountry extends Component {
  constructor() {
    super();
    this.state = {
      groups : [],
      totalDownloads: 0
    };
  }
  getDownloads() {
    let url = this.props.backendurl + '/downloads/by-country';
    this.props.endDate.set({hour:23,minute:59,second:59,millisecond:59});
    let params = {
      startDate: this.props.startDate.format(),
      endDate: this.props.endDate.format()
    };
    axios.get(url, {params})
      .then((response) => {
        let data = response.data;
        this.setState({
          groups: data,
          totalDownloads: _.sumBy(data, 'total')
        });
      })
  }
  render() {
    return (
      <div>
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

        { this.state.groups.length > 0 &&
          <table className="table table-bordered">
              <thead>
                <tr>
                  <th>Country</th>
                  <th>Downloads</th>
                  <th>%</th>
                </tr>
              </thead>
              <tbody>
                {this.state.groups.map(country => {
                  return (
                    <tr key={country.country}>
                      <td>{country.country}</td>
                      <td className="text-right">{country.total}</td>
                      <td className="text-right">{parseFloat(country.total / this.state.totalDownloads * 100).toFixed(2)}</td>
                    </tr>
                  )
                })}
                <tr>
                  <td><strong>Total</strong></td>
                  <td colSpan="2" className="text-right"><strong>{this.state.totalDownloads}</strong></td>
                </tr>
              </tbody>
            </table>
        }
      </div>
    )
  }
}

export default ByCountry
