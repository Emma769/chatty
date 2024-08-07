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
