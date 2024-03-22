import { readFile } from "node:fs/promises";

import { getPlaiceholder, type GetPlaiceholderOptions } from "plaiceholder";

async function getBlurData(path: string, options?: GetPlaiceholderOptions) {
  const buffer = await readFile(path);
  const data = await getPlaiceholder(buffer, options);

  return data;
}

export default getBlurData;
