const express = require('express');
const axios = require('axios');
const querystring = require('node:querystring'); 

const app = express();
const PORT = process.env.PORT || 31337;

//let codeValue = 'oHISzAR5RdTef7jVntWuA48Gf044Vr'

const fetchUser = async code => {
	let query = querystring.stringify({
		'client_id': '1242405621815316502',
		'client_secret': 'SdY9h2ilQb42AKV3dL8pscd9vcvUc0Bo',
		'grant_type': 'authorization_code',
		'code': code,
		'redirect_uri': 'http://localhost:5173/#/discord',
		'scope': 'identify'
    })
    let headers = {
        'Content-Type': 'application/x-www-form-urlencoded'
    }

	const token = await axios.post('https://discordapp.com/api/oauth2/token', query, headers)

	console.log("token", token.status)

    const user = await axios.get(`https://discordapp.com/api/users/@me`, {
        headers: {
            "Authorization": "Bearer " + token.data.access_token,
    	}
    })

	console.log("user", user.status)
	return await user
}

app.get('/', (req, res) => {
	console.log("code", req.query.code)
	fetchUser(req.query.code)
		.then(user => {
			console.log("response", user)
			console.log(user.status, user.statusText);
		 	res.send(user.data);
		});
});

app.listen(PORT, () => {
  console.log(`Server running on port ${PORT}`);
});





