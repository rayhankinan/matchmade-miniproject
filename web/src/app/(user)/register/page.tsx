import RegisterForm from "~/components/app/main/register";
import LoginAlert from "~/components/warning/login-alert";

import getSession from "~/server/auth";

export default function RegisterPage() {
  const session = getSession();

  return (
    <main className="flex min-h-screen flex-col items-center justify-center gap-12">
      {session !== null ? <LoginAlert /> : <RegisterForm />}
    </main>
  );
}
