"use client";

import type JwtPayload from "~/types/jwt-payload";
import SessionContext from "~/context/auth";

export default function SessionProvider({
  children,
  value,
}: React.PropsWithChildren<{ value: JwtPayload | null }>) {
  return (
    <SessionContext.Provider value={value}>{children}</SessionContext.Provider>
  );
}
