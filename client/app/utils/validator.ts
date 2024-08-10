export type ValidationResult = { cond: boolean; msg: string };
export type Validationfn<T> = (t: T) => ValidationResult;

export const validate = <T>(t: T, ...fns: Validationfn<T>[]) => {
  const errors: Record<string, string> = {};

  fns.forEach((fn) => {
    const { cond, msg } = fn(t);

    if (!cond) {
      const parts = msg.split(":");
      if (parts.length !== 2) {
        throw new Error("invalid validation message format: <key:message>");
      }
      const [key, errmsg] = parts;
      errors[key] = errmsg;
    }
  });

  const valid = Object.keys(errors).length === 0;

  return { valid, errors };
};
