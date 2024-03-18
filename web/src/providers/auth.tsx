"use client";

import { useRef } from "react";
import { createAuthStore, type AuthProps, type AuthStore } from "~/store/auth";
import AuthContext from "~/context/auth";

export default function AuthProvider({
  children,
  ...props
}: React.PropsWithChildren<AuthProps>) {
  const storeRef = useRef<AuthStore | null>(null);
  if (!storeRef.current) {
    storeRef.current = createAuthStore(props);
  }

  return (
    <AuthContext.Provider value={storeRef.current}>
      {children}
    </AuthContext.Provider>
  );
}
