const express = require('express')
const {spawn} = require('child_process');
const app = express()
const PORT = 8080
var bodyParser = require('body-parser')

app.use(bodyParser.urlencoded({extended: true}));
app.use(express.static(__dirname + '/views'));
app.set('view engine', 'ejs');

app.get('/', (req, res) => {
 res.render('sha256');
});

app.post('/getHash', (req, res) => {
  if (req.body.userInp === '') {
    res.send({'hash': ''});
  } else {
    const shaPyScript = spawn('python3', ['shaPyScript.py', req.body.userInp]);
    var hashedInp
    shaPyScript.stdout.on('data', function (data) {
      hashedInp = data.toString();
    });
    shaPyScript.on('close', (code) => {
      res.send({'hash': hashedInp});
    });
  }
});

app.listen(PORT, '0.0.0.0', () => console.log(`Listening on ${ PORT }`));
