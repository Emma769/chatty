import { Link, useFetcher } from "@remix-run/react";
import styles from "./AuthForm.module.css";
import { useRef } from "react";
import { useFocus } from "~/hooks/useFocus";

export const AUTH_INTENT = {
  register: "register",
  login: "login",
};

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
    fetcher.submit(e.currentTarget, {
      method: "post",
    });
  };

  return (
    <div className={styles.wrapper}>
      <p>{kind === "register" ? "Register" : "Login"}</p>
      <fetcher.Form onSubmit={handleSubmit}>
        <input
          type="hidden"
          name="_action"
          value={kind === "register" ? AUTH_INTENT.register : AUTH_INTENT.login}
        />
        {kind === "register" && (
          <div>
            <label>Username:</label>
            <input type="text" ref={usernameRef} name="username" />
          </div>
        )}
        <div>
          <label>Email:</label>
          <input type="email" ref={emailRef} name="email" />
        </div>
        <div>
          <label>Password:</label>
          <input type="password" name="password" />
        </div>
        <div>
          <button>{kind === "register" ? "Register" : "Login"}</button>
        </div>
      </fetcher.Form>
      <div className={styles.footer}>
        <small>
          {kind === "register" ? (
            <>
              Already registered? Login <Link to="/login">here</Link>
            </>
          ) : (
            <>
              Not registered? Signup <Link to="/register">here</Link>
            </>
          )}
        </small>
      </div>
    </div>
  );
}

export default AuthForm;
