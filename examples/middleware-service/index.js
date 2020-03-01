const express = require("express");
const axios = require("axios");
const querystring = require("querystring");

const app = express();

const client = axios.create({
  baseURL: process.env.BACKEND_BASEURL || "http://localhost:3000"
});

app.get("/get", async (req, res) => {
  const { query } = req;

  console.log(query);

  query.middlewareInvolved = true;

  try {
    const response = await client.get("/get?" + querystring.stringify(query));

    return res.status(200).json(response.data);
  } catch (error) {
    console.error("error: " + error.code);

    if (error.code === "ECONNREFUSED") {
      console.log("please make sure that backend is running");
    }

    return res.status(400).end();
  }
});

const port = isNaN(process.env.PORT) ? 3500 : process.env.PORT;

app.listen(port, () => {
  console.log(`middleware listening on port ${port}.`);
});
