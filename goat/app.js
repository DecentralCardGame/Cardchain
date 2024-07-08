const express = require('express');
const axios = require('axios');
const querystring = require('node:querystring');

const app = express();
const PORT = process.env.PORT || 31337;

const fetchUser = code => {
	let query = querystring.stringify({
		'client_id': '1242405621815316502',
		'client_secret': 'SdY9h2ilQb42AKV3dL8pscd9vcvUc0Bo',
		'grant_type': 'authorization_code',
		'code': code,
		'redirect_uri': 'https://crowdcontrol.network/#/discord',
		'scope': 'identify'
    })
    let headers = {
        'Content-Type': 'application/x-www-form-urlencoded'
    }

	return axios.post('https://discordapp.com/api/oauth2/token', query, headers)
	.then(token => {
		return axios.get(`https://discordapp.com/api/users/@me`, {
	        headers: {
	            "Authorization": "Bearer " + token.data.access_token,
	    	}
	    })
	})

}

app.get('/', (req, res) => {
	if (!req.query.code) {
		throw new Error('No code provided - you must provide a token code from Discord')
	}

	console.log("code", req.query.code)

	return fetchUser(req.query.code)
	.then(user => {
		console.log("response", user.data)
		console.log("status", user.status, user.statusText);
	 	res.send(user.data);
	})
	.catch(err => {
		console.error(err.response)
		res.status(500).send(err.message)
	})
		
});

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});





