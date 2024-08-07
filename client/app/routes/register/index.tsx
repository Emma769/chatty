import {
  isRouteErrorResponse,
  redirect,
  useNavigate,
  useRouteError,
  type ClientActionFunctionArgs,
} from "@remix-run/react";
import { register } from "~/api/auth";
import AuthForm, { AUTH_INTENT } from "~/components/AuthForm";
import { dictMap } from "~/utils/funcs";

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

    try {
      await register({ username, email, password });
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
