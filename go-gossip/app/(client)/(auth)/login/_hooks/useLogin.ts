import { useMutation } from "@tanstack/react-query";
import { userCredentialsSchema } from "../../zod";
import z from "zod";
import { setCookie } from "cookies-next";
import dayjs from "dayjs";

const loginURL = `${process.env.NEXT_PUBLIC_API_URL}/login`;

const login = async (values: z.infer<typeof userCredentialsSchema>) => {
  const res = await fetch(loginURL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(values),
    credentials: "include",
  });
  if (!res.ok) throw new Error(String(res.status));
  return res.json();
};

export default function useLogin() {
  return useMutation({
    mutationFn: login,
    onSuccess: (data) => {
      setCookie("user_id", data.user_id, {
        expires: dayjs().add(15, "m").toDate(),
      });
      setCookie("username", data.username, {
        expires: dayjs().add(15, "m").toDate(),
      });
    },
  });
}
