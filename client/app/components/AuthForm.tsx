import { useFetcher } from "@remix-run/react";
import styles from "./AuthForm.module.css";
import { useRef } from "react";
import { useFocus } from "~/hooks/useFocus";

type AuthFormProps = { kind: "register" } | { kind: "login" };

function AuthForm({ kind }: AuthFormProps) {
  const fetcher = useFetcher();

  const usernameRef = useRef<HTMLInputElement | null>(null);
  const emailRef = useRef<HTMLInputElement | null>(null);

  if (kind === "register") {
    useFocus(usernameRef);
  }

  if (kind === "login") {
    useFocus(emailRef);
  }

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
  };

  return (
    <div className={styles.wrapper}>
      <p>{kind === "register" ? "Register" : "Login"}</p>
      <fetcher.Form onSubmit={handleSubmit}>
        {kind === "register" && (
          <div>
            <label>Username:</label>
            <input type="text" ref={usernameRef} />
          </div>
        )}
        <div>
          <label>Email:</label>
          <input type="email" ref={emailRef} />
        </div>
        <div>
          <label>Password:</label>
          <input type="password" />
        </div>
        <div>
          <button>{kind === "register" ? "Register" : "Login"}</button>
        </div>
      </fetcher.Form>
    </div>
  );
}

export default AuthForm;
