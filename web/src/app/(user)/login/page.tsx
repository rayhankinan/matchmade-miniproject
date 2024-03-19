import LoginForm from "~/components/app/main/login";
import LoginAlert from "~/components/warning/login-alert";

import getSession from "~/server/auth";

export default function LoginPage() {
  const session = getSession();

  return (
    <main>
      <div className="flex min-h-screen flex-col items-center justify-center">
        {session !== null ? <LoginAlert /> : <LoginForm />}
      </div>
    </main>
  );
}
