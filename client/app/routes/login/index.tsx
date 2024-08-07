import { type ClientActionFunctionArgs } from "@remix-run/react";
import AuthForm, { AUTH_INTENT } from "~/components/AuthForm";

export const meta = () => {
  return [{ title: "Chatty | Login" }];
};

export const clientAction = async ({ request }: ClientActionFunctionArgs) => {
  const { _action: action, ...data } = Object.fromEntries(
    await request.formData()
  );

  if (action === AUTH_INTENT.login) {
  }

  return null;
};

export default function Login() {
  return (
    <div>
      <AuthForm kind="login" />
    </div>
  );
}
