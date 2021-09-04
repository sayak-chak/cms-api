import http from "k6/http";
import { sleep } from "k6";

// Fill up mock data in database before running tests
export let options = {
  stages: [
    { duration: "10s", target: 100 },
    { duration: "25s", target: 100 },
    { duration: "10s", target: 1400 },
    { duration: "25s", target: 1400 },
    { duration: "10s", target: 100 },
    { duration: "25s", target: 100 },
    { duration: "10s", target: 0 },
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
