import { type ClientActionFunctionArgs } from "@remix-run/react";
import AuthForm from "~/components/AuthForm";

export const meta = () => {
  return [{ title: "Chatty | Register" }];
};

export const clientAction = async ({ }: ClientActionFunctionArgs) => {
  return null;
};

export default function Register() {
  return (
    <div>
      <AuthForm kind="register" />
    </div>
  );
}
