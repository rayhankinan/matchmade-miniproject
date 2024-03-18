"use client";

import { useRef } from "react";
import { createAuthStore, type AuthStore } from "~/store/auth";
import AuthContext from "~/context/auth";

export default function AuthProvider({ children }: React.PropsWithChildren) {
  const storeRef = useRef<AuthStore | null>(null);
  if (!storeRef.current) {
    storeRef.current = createAuthStore();
  }

  return (
    <AuthContext.Provider value={storeRef.current}>
      {children}
    </AuthContext.Provider>
  );
}
