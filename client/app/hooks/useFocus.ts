import { useEffect, RefObject } from "react";

export const useFocus = (
  ref: RefObject<HTMLInputElement> | RefObject<HTMLTextAreaElement>
) => {
  useEffect(() => {
    ref.current?.focus();
  }, []);
};
