import { BASE_URL } from "~/constants";
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

  if (!resp.ok) {
    const payload = await resp.json();
    throw new Error(payload?.detail ?? resp.statusText);
  }

  const payload = await resp.json();

  if (!isUserPayload(payload)) throw new Error("invalid user payload");

  return payload;
};

export type LoginParam = Pick<RegisterParam, "email" | "password">;

export const login = async (param: LoginParam) => { };
