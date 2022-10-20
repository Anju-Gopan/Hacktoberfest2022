// Import Node.js core module i.e http
var http = require('http');

// Create web server
var server = http.createServer(function (req, res) {
	
	// Check the URL of the current request
	if (req.url == '/') {
		
		// Set response header
		res.writeHead(200, { 'Content-Type': 'text/html' });
		
		// Set response content	
		res.write(
		`<html><body style="text-align:center;">
			<h1 style="color:green;">Home Page</h1>
			<p>Hello World!</p>
			</body></html>`);
		res.end();//end the response
	
	}
	else if (req.url == "/products") {
		
		res.writeHead(200, { 'Content-Type': 'text/html' });
		res.write(`
		<html><body style="text-align:center;">
			<h1 style="color:green;">Welcome to Products Page</h1>
		</body></html>`);
		res.end();//end the response
	
	}
	else if (req.url == "/profile") {
		
		res.writeHead(200, { 'Content-Type': 'text/html' });
		res.write(`<html><body style="text-align:center;">
		<h1 style="color:green;">You are at Profiles Page</h1>
		</body></html>`);
		res.end(); //end the response
	
	}

	else
		res.end('Invalid Request!'); //end the response

// Server object listens on port 8081
}).listen(3000, ()=>console.log('Server running on port 3000'));
