import React from 'react';

const Panel = ({ focus, header, content }) =>
  <div className="panel">
    { /* TODO: Replace hardcoded data with data from the backend */ }
    <h3 className="panel-heading">
      <div className="focus"> { focus } </div>
      <div className="text-center"> { header } </div>
    </h3>
    {content ?
      <div className="panel-body">
        { content }
      </div>
      :
      ""
    }
  </div>;


export default Panel;
