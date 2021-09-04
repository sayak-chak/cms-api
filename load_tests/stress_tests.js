import http from "k6/http";
import { sleep } from "k6";

// Fill up mock data in database before running tests
export let options = {
  stages: [
    { duration: "10s", target: 100 },
    { duration: "25s", target: 100 },
    { duration: "10s", target: 200 },
    { duration: "25s", target: 200 },
    { duration: "10s", target: 300 },
    { duration: "25s", target: 300 },
    { duration: "10s", target: 400 },
    { duration: "25s", target: 400 },
  ],
};

export default function () {
  const BASE_URL = ""; // use the appropriate base url while testing

  let responses = http.batch([
    ["GET", `${BASE_URL}/top-contents`],
    ["GET", `${BASE_URL}/top-contents/action`],
    ["GET", `${BASE_URL}/top-contents/adventure`],
  ]);

  sleep(1);
}
