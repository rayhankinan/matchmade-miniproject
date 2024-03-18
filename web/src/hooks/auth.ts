import { useContext } from "react";

import SessionContext from "~/context/auth";

export default function useSession() {
  return useContext(SessionContext);
}
