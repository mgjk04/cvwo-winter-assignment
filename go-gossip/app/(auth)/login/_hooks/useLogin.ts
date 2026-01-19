import { useMutation } from "@tanstack/react-query";
import { userCredentialsSchema } from "../../zod";
import z from "zod";

const loginURL = "http://localhost:8080/login";

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
  });
}
