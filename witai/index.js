"use strict";

const express = require('express');
const { Wit } = require('node-wit');

const app = express();
const port = process.env.PORT || '80';
const host = process.env.HOST || '';
const witaiToken = process.env.WITAITOKEN;

if (!witaiToken) {
	console.error("please set WITAITOKEN to your wit.ai app token");
	process.exit(1);
}

const witaiClient = new Wit({ accessToken: witaiToken });

const prereqs = {
	"CSE 143": ["CSE 142"],
	"CSE 373": ["CSE 143"],
	"INFO 330": ["INFO 360"],
	"INFO 341": ["CSE 142", "CSE 143"],
	"INFO 343": ["CSE 142", "INFO 201"],
	"INFO 344": ["INFO 343", "INFO 340"],
	"INFO 445": ["INFO 340"]
};

function handlePreReqs(req, res, witaiData) {
	//TODO: use the entities in the `witaiData`
	//to determine which course the user is 
	//asking about, and respond with a natural
	//language message indicating the pre-reqs
}

app.get("/chatbot", (req, res, next) => {
	//TODO: use witaiClient.message() to
	//extract meaning from the value in the
	//`q` query string param and respond
	//accordingly

});

app.listen(port, host, () => {
	console.log(`server is listening at http://${host}:${port}`);
});
