import { Component } from 'react';
import * as React from 'react';
import * as Props from '../props';
import ChatArea from './ChatArea';


class Chat extends Component<Props.IChatProps, {}> {
  public render() {
    return (
      <div className="container">
        <ChatArea client={this.props.client}/>
      </div>
    );
  }
}

export default Chat;