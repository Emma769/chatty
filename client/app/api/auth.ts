import { BASE_URL } from "~/constants";
import { TokenPayload } from "~/types/token";
import type { User } from "~/types/user";

export type RegisterParam = Pick<User, "username" | "email"> & {
  password: string;
};

const isUserPayload = (payload: unknown): payload is User => {
  return (
    payload !== null &&
    typeof payload === "object" &&
    "user_id" in payload &&
    "username" in payload &&
    "email" in payload &&
    "created_at" in payload
  );
};

export const register = async (param: RegisterParam) => {
  const resp = await fetch(`${BASE_URL}/api/users`, {
    method: "POST",
    headers: new Headers([["Content-Type", "application/json"]]),
    body: JSON.stringify(param),
  });

  if (!resp.ok) throw new Error("non-2** response", { cause: { resp } });
  const payload = await resp.json();
  if (isUserPayload(payload)) return payload;
  throw new Error("invalid user payload");
};

export type LoginParam = Pick<RegisterParam, "email" | "password">;

const isTokenPayload = (payload: unknown): payload is TokenPayload => {
  return (
    payload !== null &&
    typeof payload === "object" &&
    "access_token" in payload &&
    "user" in payload
  );
};

export const login = async (param: LoginParam) => {
  const resp = await fetch(`${BASE_URL}/api/tokens/login`, {
    method: "POST",
    headers: new Headers([["Content-Type", "application/json"]]),
    body: JSON.stringify(param),
  });

  if (!resp.ok) throw new Error("non-2** resp", { cause: { resp } });
  const payload = await resp.json();
  if (isTokenPayload(payload)) return payload;
  throw new Error("invalid token payload");
};
