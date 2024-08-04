import AuthForm from "~/components/AuthForm";

export const meta = () => {
  return [{ title: "Chatty | Login" }];
};

export default function Login() {
  return (
    <div>
      <AuthForm kind="login" />
    </div>
  );
}
