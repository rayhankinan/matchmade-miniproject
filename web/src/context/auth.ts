import { createContext } from "react";

import type JwtPayload from "~/types/jwt-payload";

const SessionContext = createContext<JwtPayload | null>(null);

export default SessionContext;
