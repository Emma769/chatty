import { User } from "./user";

export type AccessToken = {
  value: string;
  valid_till: string;
};

export type RefreshToken = {
  value: string;
  valid_till: string;
};

export type TokenPayload = {
  user: User;
  access_token: AccessToken;
  refresh_token?: RefreshToken;
};
