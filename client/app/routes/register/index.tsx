import {
  isRouteErrorResponse,
  json,
  redirect,
  useNavigate,
  useRouteError,
  type ClientActionFunctionArgs,
} from "@remix-run/react";
import { register, RegisterParam } from "~/api/auth";
import AuthForm, { AUTH_INTENT } from "~/components/AuthForm";
import { dictMap, validEmail, validString } from "~/utils/funcs";
import { validate, type Validationfn } from "~/utils/validator";

export const meta = () => {
  return [{ title: "Chatty | Register" }];
};

export const clientAction = async ({ request }: ClientActionFunctionArgs) => {
  const { _action: action, ...data } = Object.fromEntries(
    await request.formData()
  );

  if (action === AUTH_INTENT.register) {
    const { username, email, password } = dictMap(data, (arg) =>
      arg.toString()
    );

    const param: RegisterParam = {
      username,
      email,
      password,
    };

    const valfns: Validationfn<RegisterParam>[] = [
      (t) => ({
        cond: validString(t.username),
        msg: "username:cannot be blank",
      }),
      (t) => ({ cond: validString(t.email), msg: "email:cannot be blank" }),
      (t) => ({
        cond: validString(t.password),
        msg: "password:cannot be blank",
      }),
      (t) => ({
        cond: validEmail(t.email),
        msg: "email:provide a valid email",
      }),
      (t) => ({
        cond: t.password.length >= 8,
        msg: "password:must be at least 8 characters long",
      }),
    ];

    const { valid, errors } = validate(param, ...valfns);

    if (!valid) {
      return json({ errors });
    }

    try {
      await register(param);
      return redirect("/login");
    } catch (error) {
      const message = error instanceof Error ? error.message : "unknown error";
      throw new Response(message, { status: 400 });
    }
  }

  return null;
};

export function ErrorBoundary() {
  const error = useRouteError();
  const detail = isRouteErrorResponse(error) ? error.data : "unknown error";

  const navigate = useNavigate();
  const goback = () => navigate(-2);

  return (
    <div>
      <p>Lmao! you done cook beans</p>
      <div>
        ERR: <span>{detail}</span>
      </div>
      <button type="button" onClick={goback}>
        Go Back
      </button>
    </div>
  );
}

export default function Register() {
  return (
    <div>
      <AuthForm kind="register" />
    </div>
  );
}
