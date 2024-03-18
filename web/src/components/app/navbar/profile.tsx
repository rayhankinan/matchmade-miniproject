"use client";

import Link from "next/link";

import { Avatar, AvatarFallback } from "~/components/ui/avatar";
import { Button } from "~/components/ui/button";

import useSession from "~/hooks/auth";

export default function Profile() {
  const session = useSession();

  if (!session)
    return (
      <>
        <Button variant="outline" asChild>
          <Link href="/login">Sign in</Link>
        </Button>
        <Button>Sign Up</Button>
      </>
    );

  return (
    <Avatar>
      <AvatarFallback>
        {session.username[0] ? session.username[0].toUpperCase() : "?"}
      </AvatarFallback>
    </Avatar>
  );
}
