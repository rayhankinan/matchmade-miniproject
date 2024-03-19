import axios from "axios";

import { env } from "~/env";

const movieApi = axios.create({
  baseURL: env.NEXT_PUBLIC_MOVIE_API_URL,
  headers: {
    Authorization: `Bearer ${env.NEXT_PUBLIC_MOVIE_API_TOKEN}`,
    "Content-Type": "application/json",
  },
});

export default movieApi;
