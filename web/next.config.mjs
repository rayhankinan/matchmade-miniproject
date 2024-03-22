// @ts-check
import withPlaiceholder from "@plaiceholder/next";

/**
 * Run `build` or `dev` with `SKIP_ENV_VALIDATION` to skip env validation. This is especially useful
 * for Docker builds.
 */
await import("./src/env.js");

/** @type {import("next").NextConfig} */
const config = {
  output: "standalone",
  images: {
    remotePatterns: [
      {
        protocol: "https",
        hostname: "image.tmdb.org",
      },
    ],
  },
};

// TODO: Utilize "plaiceholder" to blur the images in server-side rendering.
export default withPlaiceholder(config);
