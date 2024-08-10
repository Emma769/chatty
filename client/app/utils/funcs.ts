export const dictMap = <T, U>(
  param: Record<string, T>,
  fn: (arg: T) => U
): Record<string, U> => {
  const entries = Object.entries(param);
  return Object.fromEntries(
    entries.map(([k, v]) => {
      return [k, fn(v)];
    })
  );
};

export const validEmail = (s: string) => {
  const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
  return emailRegex.test(s);
};

export const validString = (s: string) => {
  return s.trim().length !== 0;
};
