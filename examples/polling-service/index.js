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

    console.log(response.data);
  } catch (error) {
    console.error("error: " + error.code);

    if (error.code === "ECONNREFUSED") {
      console.log("please make sure that backend is running");
    }
  }
}, ms);
