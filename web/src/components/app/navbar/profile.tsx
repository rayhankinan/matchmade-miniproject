"use client";

import Link from "next/link";

import { Avatar, AvatarFallback } from "~/components/ui/avatar";
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
      <Button variant="outline" asChild>
        <Link href="/login">Sign in</Link>
      </Button>
      <Button>Sign Up</Button>
    </>
  );
}
