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
  const errors = fetcher.data?.errors;

  const usernameRef = useRef<HTMLInputElement | null>(null);
  const emailRef = useRef<HTMLInputElement | null>(null);

  useFocus(kind === "register" ? usernameRef : emailRef);

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
                placeholder="username"
                id={errors?.username ? styles["warn-username"] : undefined}
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
              type="email"
              ref={emailRef}
              name="email"
              required
              autoComplete="off"
              spellCheck="false"
              placeholder="email"
              id={errors?.email ? styles["warn-email"] : undefined}
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
              minLength={kind === "register" ? 8 : undefined}
              placeholder="password"
              id={errors?.password ? styles["warn-password"] : undefined}
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
