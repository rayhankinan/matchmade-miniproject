import { cookies } from "next/headers";
import { jwtDecode } from "jwt-decode";

import type JwtPayload from "~/types/jwt-payload";

export default function getAuth() {
  const token = cookies().get("AUTH_TOKEN");
  const decodedToken = token ? jwtDecode<JwtPayload>(token.value) : null;

  return decodedToken;
}
