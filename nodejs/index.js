
(async () => {
	const axios = require('axios')

	console.log("Fetching application data....")

	const data = await axios('https://engineering-task.elancoapps.com/api/applications').then(resp => resp.data);


	data.forEach(application => console.log(application))
})()