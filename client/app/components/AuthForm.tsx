import { Link, useFetcher } from "@remix-run/react";
import { useRef } from "react";
import { useFocus } from "~/hooks/useFocus";
import styles from "./AuthForm.module.css";

export const AUTH_INTENT = {
  register: "register",
  login: "login",
};

type AuthFormProps = { kind: "register" } | { kind: "login" };

function AuthForm({ kind }: AuthFormProps) {
  const fetcher = useFetcher<{ errors: Record<string, string> }>();

  const data = fetcher.data;
  const errors = data?.errors;

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
      <p className={styles.title}>
        {kind === "register" ? "Register" : "Login"}
      </p>
      <fetcher.Form onSubmit={handleSubmit}>
        <input
          type="hidden"
          name="_action"
          value={kind === "register" ? AUTH_INTENT.register : AUTH_INTENT.login}
        />
        {kind === "register" && (
          <>
            <div className={styles.group}>
              <input
                type="text"
                ref={usernameRef}
                name="username"
                required
                autoComplete="off"
                spellCheck="false"
              />
              <label>Username</label>
            </div>
            <div className={styles.danger}>
              {errors?.username && errors.username}
            </div>
          </>
        )}
        <>
          <div className={styles.group}>
            <input
              type="text"
              ref={emailRef}
              name="email"
              required
              autoComplete="off"
              spellCheck="false"
            />
            <label>Email</label>
          </div>
          <div className={styles.danger}>{errors?.email && errors.email}</div>
        </>
        <>
          <div className={styles.group}>
            <input
              type="password"
              name="password"
              required
              autoComplete="off"
            />
            <label>Password</label>
          </div>
          <div className={styles.danger}>
            {errors?.password && errors.password}
          </div>
        </>
        <div className={styles.btn_wrapper}>
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
