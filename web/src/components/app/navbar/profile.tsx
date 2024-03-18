"use client";

import Link from "next/link";
import { useRouter } from "next/navigation";
import { useMutation } from "@tanstack/react-query";

import { Avatar, AvatarFallback } from "~/components/ui/avatar";
import { Button } from "~/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "~/components/ui/dropdown-menu";
import { useToast } from "~/components/ui/use-toast";

import useSession from "~/hooks/auth";

import api from "~/client/api";

type LogoutButtonResponse = {
  data: null;
};

export default function Profile() {
  const session = useSession();
  const router = useRouter();

  const { toast } = useToast();

  const logoutMutation = useMutation({
    mutationFn: async () => await api.post<LogoutButtonResponse>("/logout"),
    onSuccess: () => router.refresh(),
    onError: (error) =>
      toast({
        variant: "destructive",
        title: "An error occurred",
        description: error.message,
      }),
  });

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
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Avatar className="cursor-pointer">
          <AvatarFallback>
            {session.username[0] ? session.username[0].toUpperCase() : "?"}
          </AvatarFallback>
        </Avatar>
      </DropdownMenuTrigger>
      <DropdownMenuContent>
        <DropdownMenuLabel>My Account</DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuItem onClick={() => logoutMutation.mutate()}>
          Log Out
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
