"use client";

import { Avatar, AvatarFallback } from "@radix-ui/react-avatar";
import { Button } from "~/components/ui/button";
import useAuthContext from "~/hooks/auth";

export default function Profile() {
  const session = useAuthContext((state) => state.session);

  return session !== null ? (
    <Avatar>
      <AvatarFallback>{session.email}</AvatarFallback>
    </Avatar>
  ) : (
    <>
      <Button variant="outline">Sign in</Button>
      <Button>Sign Up</Button>
    </>
  );
}
