"use client";

import Link from "next/link";

import { Avatar, AvatarFallback } from "~/components/ui/avatar";
import { Button } from "~/components/ui/button";

import useSession from "~/components/hooks/auth";

export default function Profile() {
  const session = useSession();

  return session !== null ? (
    <Avatar>
      <AvatarFallback>{session.username}</AvatarFallback>
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
