import React from 'react';

import Panel from './partials/Panel'

import 'whatwg-fetch';

class Hubs extends React.Component {
  constructor(props) {
    super(props);
    this.state = { hubs: [] };
  }

  componentDidMount() {
    fetch('http://localhost:4000/api/hub').then(response => response.json()).then((body) => {
      this.setState({ hubs: body.hubs });
    });
  }

  render() {
    return (
      <div>
        <div>
          <h2 className="text-primary text-center">Hubs</h2>
          {this.state.hubs.map(hub => (
            <Panel key={hub.id} focus={hub.slug} header={hub.name} content={hub.about} />
          ))}
        </div>
      </div>
    );
  }
}

export default Hubs;
