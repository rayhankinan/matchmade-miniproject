import { redirect } from "next/navigation";

import LoginForm from "~/components/app/main/login";

import getSession from "~/server/auth";

export default function LoginPage() {
  const session = getSession();

  if (session) redirect("/");

  return (
    <main className="flex min-h-screen flex-col items-center justify-center">
      <LoginForm />
    </main>
  );
}
