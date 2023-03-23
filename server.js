let http = require("http");
let fs = require("fs");

const server = http.createServer(async (req, res) => {

    fs.readFile('./assets/' + req.url, function (error, data) {
        console.log('./assets' + req.url);
        if (error) {
            res.statusCode = 404;
            res.end("Resourse not found!");
        }
        else {
            res.end(data);
        }
    });
});
server.listen(3000);
console.log("running");