import React from 'react';

const About = () =>
	<div>
		<div className="main">
			<h2 className="text-primary text-center">Upcoming Events</h2>
			<div className="panel">
				{ /* TODO: Replace hardcoded data with data from the backend */ }
				<h3 className="panel-heading">
					<span className="focus">HUB</span> Title of Event
				</h3>
				<div className="panel-body">
					<p>A small description of the event.</p>
					<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. In eu aliquet felis. Fusce consectetur suscipit ipsum, eu convallis diam cursus vel. Praesent molestie dui quis ultricies pharetra. Pellentesque hendrerit sed massa quis accumsan. Mauris odio mi, aliquam at metus in, bibendum consectetur augue. Duis finibus, massa at pellentesque gravida, ante velit hendrerit enim, quis facilisis sapien velit ut justo. Fusce ut velit vel sem porttitor rutrum in a mauris. Etiam tristique leo quis odio feugiat condimentum. Cras in nunc consectetur, malesuada eros at, accumsan lacus.</p>
				</div>
			</div>
		</div>
		<div className="sidebar">
			<h2 className="text-primary text-center">Trending</h2>
		</div>
	</div>;


export default About;
