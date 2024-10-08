import {
  isRouteErrorResponse,
  json,
  redirect,
  useRouteError,
  type ClientActionFunctionArgs,
} from "@remix-run/react";
import { login, type LoginParam } from "~/api/auth";
import AuthForm, { AUTH_INTENT } from "~/components/AuthForm";
import { dictMap, validEmail, validString } from "~/utils/funcs";
import { validate, type Validationfn } from "~/utils/validator";

export const meta = () => {
  return [{ title: "Chatty | Login" }];
};

export const clientAction = async ({ request }: ClientActionFunctionArgs) => {
  const { _action: action, ...data } = Object.fromEntries(
    await request.formData()
  );

  if (action === AUTH_INTENT.login) {
    const { email, password } = dictMap(data, (d) => d.toString());

    const param: LoginParam = {
      email,
      password,
    };

    const valfns: Validationfn<LoginParam>[] = [
      (t) => ({ cond: validString(t.email), msg: "email:cannot be blank" }),
      (t) => ({
        cond: validString(t.password),
        msg: "password:cannot be blank",
      }),
      (t) => ({
        cond: validEmail(t.email),
        msg: "email:provide a valid email",
      }),
    ];

    const { valid, errors } = validate(param, ...valfns);
    if (!valid) return json({ errors });

    try {
      const tokens = await login(param);
      return redirect("/");
    } catch (error) {
      console.log(error);
    }
  }

  return null;
};

export function ErrorBoundary() {
  const error = useRouteError();

  const detail = isRouteErrorResponse(error)
    ? error.data
    : error instanceof Error
      ? error.message
      : "unknown error";

  return (
    <div>
      <div>Lmao! you done cook beans</div>
      <div>
        ERR: <span>{detail}</span>
      </div>
    </div>
  );
}

export default function Login() {
  return (
    <div>
      <AuthForm kind="login" />
    </div>
  );
}
