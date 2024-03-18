import { createContext } from "react";

import { type AuthStore } from "~/store/auth";

const AuthContext = createContext<AuthStore | null>(null);

export default AuthContext;
