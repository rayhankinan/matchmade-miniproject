import Cookies from "js-cookie";
import { jwtDecode } from "jwt-decode";
import { createStore } from "zustand";

import type JwtPayload from "~/types/jwt-payload";

export interface AuthProps {
  session: JwtPayload | null;
}

export interface AuthState extends AuthProps {
  updateSessionToMatchCookie: () => void;
  removeSession: () => void;
}

export type AuthStore = ReturnType<typeof createAuthStore>;

export const createAuthStore = (props: AuthProps) =>
  createStore<AuthState>()((set) => ({
    session: props.session,
    updateSessionToMatchCookie: () => {
      // Get the token from the cookie
      const token = Cookies.get("AUTH_TOKEN");
      const decodedToken = token ? jwtDecode<JwtPayload>(token) : null;

      // Set the session in the store
      set({ session: decodedToken });
    },
    removeSession: () => {
      // Remove the token from the cookie
      Cookies.remove("AUTH_TOKEN");

      // Remove the session from the store
      set({ session: null });
    },
  }));
