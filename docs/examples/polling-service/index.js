const axios = require("axios");

const client = axios.create({
  baseURL: process.env.API_BASEURL || "http://localhost:3000"
});

const ms = isNaN(process.env.POLL_INTERVAL_MS)
  ? 5000
  : process.env.POLL_INTERVAL_MS;

setInterval(async () => {
  try {
    const response = await client.get("/get?hello=world");

    console.log(response.data.args);
  } catch (error) {
    console.error("error: " + error.code);

    console.log("please make sure that api is running (also check port)");
  }
}, ms);
