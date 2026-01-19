import { useMutation } from "@tanstack/react-query";
import z from "zod";
import { userCredentialsSchema } from "../../zod";

const signupURL = 'http://localhost:8080/signup'

const signup = async (values: z.infer<typeof userCredentialsSchema>) => {
    const res = await fetch(signupURL, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(values),
    credentials: 'include'
  });
  if (!res.ok) throw new Error(String(res.status));
  return res.json();
}

export default function useSignup(){
    return useMutation({
        mutationFn: signup, 
    })
}