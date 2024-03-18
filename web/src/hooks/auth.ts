import { useContext } from "react";
import { useStore } from "zustand";

import AuthContext from "~/context/auth";
import { type AuthState } from "~/store/auth";

export default function useAuthContext<T>(
  selector: (state: AuthState) => T,
): T {
  const store = useContext(AuthContext);
  if (!store) {
    throw new Error("useSession must be used within an AuthProvider");
  }

  return useStore(store, selector);
}
